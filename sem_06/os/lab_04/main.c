#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <dirent.h>
#include <errno.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>
#define _POSIX1_SOURCE 2

#define BUF_SIZE 102400

// pfn = 0-54 first bits.
// soft-dirty = 55 bit.
// file/shared = 61 bit.
// swapped = 62 bit.
// data = 63 bit.
void print_page(uint64_t address, uint64_t data, FILE *out) {
    fprintf(out, "0x%015lx: %-10lx %-13ld %-14ld %-11ld %-10ld\n", address, data & (((uint64_t)1 << 55) - 1), (data >> 55) & 1, (data >> 61) & 1, (data >> 62) & 1, (data >> 63) & 1);
}

void get_pagemap_info(const char *pid, FILE *out) {
	fprintf(out, "PAGEMAP\n");
	
    fprintf(out, "addr               pfn        soft-dirty    file/shared    swapped     present\n");

    char path[PATH_MAX];
    snprintf(path, PATH_MAX, "/proc/%s/maps", pid);
    
    FILE *maps = fopen(path, "r");

    snprintf(path, PATH_MAX, "/proc/%s/pagemap", pid);
    
    int pm_fd = open(path, O_RDONLY);
    
    char buf[BUF_SIZE + 1] = "\0";
    int len;

    // Read maps.
    while ((len = fread(buf, 1, BUF_SIZE, maps)) > 0) {
        for (int i = 0; i < len; i++)
            if (buf[i] == '\0')
                buf[i] = '\n';

        buf[len] = '\0';

        // Read lines from maps.
        char *save_row;
        char *row = strtok_r(buf, "\n", &save_row);

        while (row) {
            // Get first column (address) from maps line.
            char *addresses = strtok(row, " ");
            char *start_str, *end_str;

            // Get start and end of address space. Address column similar to 123-456.
            if ((start_str = strtok(addresses, "-")) && (end_str = strtok(NULL, "-"))) {
                uint64_t start = strtoul(start_str, NULL, 16);
                uint64_t end = strtoul(end_str, NULL, 16);

				// From start address to end address with PAGE SIZE step. (Pages loop)
                for (uint64_t addr = start; addr < end; addr += sysconf(_SC_PAGE_SIZE)) {
                    uint64_t page_data;
                    
                    
                    // addr / PAGE SIZE is page number. index = offset of page info.
                    uint64_t index = addr / sysconf(_SC_PAGE_SIZE) * sizeof(page_data);

					// Read offset from pagemap file.
                    pread(pm_fd, &page_data, sizeof(page_data), index);
                    
                    // Print page info.
                    print_page(addr, page_data, out);
                }
            }

            row = strtok_r(NULL, "\n", &save_row);
        }
    }

    fclose(maps);
    close(pm_fd);
    
    fprintf(out, "------------------------------------------------------------------------\n\n\n");
}


void get_maps_info(const char *pid, FILE *out) {
	fprintf(out, "MAPS\n");
	
    fprintf(out, "address                   perms  offset dev   inode                      pathname                                                                            pagenum\n");
    
    char path[PATH_MAX];
    snprintf(path, PATH_MAX, "/proc/%s/maps", pid);
    
    FILE *maps = fopen(path, "r");

    char buf[BUF_SIZE + 1] = "\0";
    int len;
    
    uint64_t pagesum = 0;

    while ((len = fread(buf, 1, BUF_SIZE, maps)) > 0) 
    {
                
        for (int i = 0; i < len; i++) 
            if (buf[i] == '\0') 
                buf[i] = '\n';

        buf[len] = '\0';

        char *save_row;
        char *row = strtok_r(buf, "\n", &save_row);

        while (row) {
            fprintf(out, "%-155s", row);
            
            // Get first column (address) from maps line.
            char *addresses = strtok(row, " ");
            char *start_str, *end_str;
            uint64_t pagenum;

            // Get start and end of address space. Address column similar to 123-456.
            if ((start_str = strtok(addresses, "-")) && (end_str = strtok(NULL, "-"))) 
            {
                uint64_t start = strtoul(start_str, NULL, 16);
                uint64_t end = strtoul(end_str, NULL, 16);
                
                pagenum = (end - start) / sysconf(_SC_PAGE_SIZE);
                
                pagesum += pagenum;
            }
            
            fprintf(out, "  %ld\n", pagenum);
            
            row = strtok_r(NULL, "\n", &save_row);
        }
    }

    fclose(maps);
    
    fprintf(out, "\nPAGE NUM: %ld\nBYTES: %ld\n", pagesum, pagesum * sysconf(_SC_PAGE_SIZE));
    
    fprintf(out, "------------------------------------------------------------------------\n\n\n");
}

void get_stat_info(const char *pid, FILE *out) 
{
	fprintf(out, "STAT\n");
		
	char* headers[] = {"pid", "comm", "state", "ppid", "pgrp", "session", "tty_nr", "tpgid", "flags", "minflt", "cminflt", "majflt", "cmajflt", "utime", "stime", "cutime", "cstime", "priority", "nice", "num_threads", "itrealvalue", "starttime", "vsize", "rss", "rsslim", "startcode", "endcode", "startstack", "kstkesp", "kstkeip", "signal", "blocked", "sigignore", "sigcatch", "wchan", "nswap", "cnswap", "exit_signal", "processor", "rt_priority", "policy", "delayacct_blkio_ticks", "guest_time", "cguest_time", "start_data", "end_data", "start_brk", "arg_start", "arg_end", "env_start", "env_end", "exit_code"};
		
	char path[PATH_MAX];
	snprintf(path, PATH_MAX, "/proc/%s/stat", pid);
	
	FILE *stat = fopen(path, "r");

    char buf[BUF_SIZE + 1] = "\0";
    int len;

    while ((len = fread(buf, 1, BUF_SIZE, stat)) > 0) 
    {
        for (int i = 0; i < len; i++) 
            if (buf[i] == '\0') 
                buf[i] = ' ';

        buf[len] = '\0';

        char *save_row;
        char *row = strtok_r(buf, " ", &save_row);
		int i = 0;
        while (row) 
        {
            fprintf(out, "%-24s  %s\n", headers[i], row);
            row = strtok_r(NULL, " ", &save_row);
            i++;
        }
    }

    fclose(stat);
    
    fprintf(out, "------------------------------------------------------------------------\n\n\n");
}

void get_status_info(const char *pid, FILE *out) 
{
	fprintf(out, "STATUS\n");
    
    char path[PATH_MAX];
	snprintf(path, PATH_MAX, "/proc/%s/status", pid);
	
	FILE *stat = fopen(path, "r");

    char buf[BUF_SIZE + 1] = "\0";
    int len;

    while ((len = fread(buf, 1, BUF_SIZE, stat)) > 0) 
    {
    	fprintf(out, "%s", buf);
    }

    fclose(stat);
    
    fprintf(out, "\n------------------------------------------------------------------------\n\n\n");
}

void get_task_info(const char *pid, FILE *out) 
{
	fprintf(out, "TASK\n");

	char path[PATH_MAX];
	snprintf(path, PATH_MAX, "/proc/%s/task", pid);

	struct dirent *d;
	DIR *dh = opendir(path);
	if (!dh) 
	{
		perror("open task dir\n");
		exit(1);
	}

	while ((d = readdir(dh)) != NULL)
	{
		if (d->d_name[0] == '.')
			continue;
			
		fprintf(out, "%s\n", d->d_name);
	}

	fprintf(out, "------------------------------------------------------------------------\n\n\n");
}

void get_environ_info(const char *pid, FILE *out) 
{
	fprintf(out, "ENVIRON\n");

	char buf[BUF_SIZE+1] = "\0";
	int len, i;
	FILE* f;

	char path[PATH_MAX];
	snprintf(path, PATH_MAX, "/proc/%s/environ", pid);

	f = fopen(path, "r");
	while ((len = fread(buf, 1, BUF_SIZE, f)) > 0) 
	{
		for (i = 0; i < len; i++)
			if (buf[i] == '\0')
			buf[i] = '\n';
			
		buf[len] = '\0';
		fprintf(out, "%s", buf);
	}

	fclose(f);

	fprintf(out, "------------------------------------------------------------------------\n\n\n");
}

void get_cwd_info(const char *pid, FILE *out) 
{
	fprintf(out, "CWD\n");

	char buf[BUF_SIZE+1] = "\0";

	char path[PATH_MAX];
	snprintf(path, PATH_MAX, "/proc/%s/cwd", pid);

	int l = readlink(path, buf, BUF_SIZE);
	buf[l] = '\0';
	fprintf(out, "%s\n", buf);

	fprintf(out, "------------------------------------------------------------------------\n\n\n");
}

void get_exe_info(const char *pid, FILE *out) 
{
	fprintf(out, "EXE\n");

	char buf[BUF_SIZE+1] = "\0";

	char path[PATH_MAX];
	snprintf(path, PATH_MAX, "/proc/%s/exe", pid);

	int l = readlink(path, buf, BUF_SIZE);
	buf[l] = '\0';
	fprintf(out, "%s\n", buf);

	fprintf(out, "------------------------------------------------------------------------\n\n\n");
}

void get_cmdline_info(const char *pid, FILE *out) 
{
	fprintf(out, "CMDLINE\n");
    
    char path[PATH_MAX];
	snprintf(path, PATH_MAX, "/proc/%s/cmdline", pid);
	
	FILE *stat = fopen(path, "r");

    char buf[BUF_SIZE + 1] = "\0";
    int len;

    while ((len = fread(buf, 1, BUF_SIZE, stat)) > 0) 
    {
    	fprintf(out, "%s", buf);
    }

    fclose(stat);
    
    fprintf(out, "\n------------------------------------------------------------------------\n\n\n");
}

void get_root_info(const char *pid, FILE *out) 
{
	fprintf(out, "ROOT\n");

	char buf[BUF_SIZE+1] = "\0";

	char path[PATH_MAX];
	snprintf(path, PATH_MAX, "/proc/%s/root", pid);

	int l = readlink(path, buf, BUF_SIZE);
	buf[l] = '\0';
	fprintf(out, "%s\n", buf);

	fprintf(out, "------------------------------------------------------------------------\n\n\n");
}

void get_fd_info(const char *pid, FILE *out) 
{
	fprintf(out, "FD\n");

	char path[PATH_MAX];
	snprintf(path, PATH_MAX, "/proc/%s/fd", pid);

	struct dirent *d;
	DIR *dh = opendir(path);
	if (!dh) 
	{
		perror("open task dir\n");
		exit(1);
	}

	while ((d = readdir(dh)) != NULL)
	{
		if (d->d_name[0] == '.')
			continue;
			
		fprintf(out, "%s -> ", d->d_name);
		
		char link_path[PATH_MAX];
		strcpy(link_path, path);
		strcat(link_path, "/");
		strcat(link_path, d->d_name);
		
		char buf[BUF_SIZE+1] = "\0";
		
		int l = readlink(link_path, buf, BUF_SIZE);
	    buf[l] = '\0';
	    fprintf(out, "%s\n", buf);
	}

	fprintf(out, "------------------------------------------------------------------------\n\n\n");
}

int main(int argc, char **argv) 
{
    char *pid = argv[1];
    FILE *file = fopen(argv[2], "w");
    if (!file) 
    {
        perror("fopen() error");
        exit(1);
    }
    
    printf("PID: %s\n", pid);
    
    if (argc >= 4)
    	get_pagemap_info(pid, file);
   
    get_maps_info(pid, file);
    get_stat_info(pid, file);
    get_status_info(pid, file);
    get_task_info(pid, file);
    get_environ_info(pid, file);
    get_cwd_info(pid, file);
    get_exe_info(pid, file);
    get_cmdline_info(pid, file);
    get_root_info(pid, file);
    get_fd_info(pid, file);
    
    return 0;
}
