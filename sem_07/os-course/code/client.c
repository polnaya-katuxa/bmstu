//
// Created by Екатерина Карпова on 22.01.2024.
//
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <dirent.h>
#include <errno.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>
#include <signal.h>
#include "client.h"
#include <sys/stat.h>

#define DEV_FILE "/dev/proclcd"
#define _POSIX1_SOURCE 2
#define BUF_SIZE 102400
#define ANSI_COLOR_RED     "\x1b[31m"
#define ANSI_COLOR_GREEN   "\x1b[32m"
#define ANSI_COLOR_YELLOW  "\x1b[33m"
#define ANSI_COLOR_BLUE    "\x1b[34m"
#define ANSI_COLOR_MAGENTA "\x1b[35m"
#define ANSI_COLOR_CYAN    "\x1b[36m"
#define ANSI_COLOR_RESET   "\x1b[0m"

int pid = 1;
int opt = 0;
int stop = 0;

int existDir(const char * name)
{
    struct stat s;
    if (stat(name,&s)) return 0;
    return S_ISDIR(s.st_mode);
};

int print_menu() {
    int rc = EXIT_SUCCESS;

    printf("\n\n===============================\n");
    printf("\n" ANSI_COLOR_MAGENTA "PROCESS INFO ASSISTANT" ANSI_COLOR_RESET "\n");

    printf(ANSI_COLOR_BLUE "Choose info to view:" ANSI_COLOR_RESET "\n");
    printf(ANSI_COLOR_BLUE "0) " ANSI_COLOR_RESET "exit\n");
    printf(ANSI_COLOR_BLUE "1) " ANSI_COLOR_RESET "cmdline\n");
    printf(ANSI_COLOR_BLUE "2) " ANSI_COLOR_RESET "opened fd number\n");
    printf(ANSI_COLOR_BLUE "3) " ANSI_COLOR_RESET "thread number\n");
    printf(ANSI_COLOR_BLUE "4) " ANSI_COLOR_RESET "virtual memory size\n");
    printf(ANSI_COLOR_BLUE "5) " ANSI_COLOR_RESET "state\n");
    printf(ANSI_COLOR_BLUE "6) " ANSI_COLOR_RESET "executable filename\n");

    printf(ANSI_COLOR_BLUE "Input menu option: " ANSI_COLOR_RESET);
    if ((rc = scanf("%d", &opt)) != 1 || opt < 0 || opt > 6) {
        printf(ANSI_COLOR_RED "invalid menu option" ANSI_COLOR_RESET "\n");
        printf("\n===============================\n");
        return EXIT_FAILURE;
    }

    if (opt != 0) {
        printf(ANSI_COLOR_BLUE "Input process PID: " ANSI_COLOR_RESET);
        if ((rc = scanf("%d", &pid)) != 1) {
            printf(ANSI_COLOR_RED "invalid process pid" ANSI_COLOR_RESET "\n");
            printf("\n===============================\n");
            return EXIT_FAILURE;
        }

        char str[1024];
        sprintf(str, "/proc/%d", pid);

        if (existDir(str) == 0) {
            printf(ANSI_COLOR_RED "process doesn't exist" ANSI_COLOR_RESET "\n");
            printf("\n===============================\n");
            return EXIT_FAILURE;
        }
    }

    printf("\n===============================\n");
    return EXIT_SUCCESS;
}

void exit_handler() {
    opt = -1;
    printf("\n" ANSI_COLOR_YELLOW "program exiting" ANSI_COLOR_RESET "\n");
}

void cmdline_handler() {
    char buf[BUF_SIZE+1] = "\0";
    int len, i;
    FILE* f;

    int dev = open(DEV_FILE, O_RDWR);

    char path[PATH_MAX];
    snprintf(path, PATH_MAX, "/proc/%d/cmdline", pid);

    f = fopen(path, "r");

    while (stop == 0) {
        while ((len = fread(buf, 1, BUF_SIZE, f)) > 0)
        {
            buf[len] = '\0';
            write(dev, buf, len-1);
        }

        fseek(f, 0, SEEK_SET);
        sleep(3);
    }

    stop = 0;

    fclose(f);
    close(dev);
}

void fds_handler() {
    char path[PATH_MAX];
    snprintf(path, PATH_MAX, "/proc/%d/fd", pid);

    int dev = open(DEV_FILE, O_RDWR);

    struct dirent *d;
    DIR *dh = opendir(path);
    if (!dh)
    {
        perror("open task dir\n");
        exit(1);
    }

    long pos = telldir(dh);

    while (stop == 0) {
        int fd_num = 0;
        while ((d = readdir(dh)) != NULL)
        {
            if (d->d_name[0] == '.')
                continue;

            fd_num++;
        }

        char str[1024];
        sprintf(str, "%d", fd_num);
        write(dev, str, strlen(str));

        seekdir(dh,pos);

        sleep(3);
    }

    stop = 0;

    closedir(dh);
    close(dev);
}

void threads_handler() {
    char path[PATH_MAX];
    snprintf(path, PATH_MAX, "/proc/%d/task", pid);

    int dev = open(DEV_FILE, O_RDWR);

    struct dirent *d;
    DIR *dh = opendir(path);
    if (!dh)
    {
        perror("open task dir\n");
        exit(1);
    }

    long pos = telldir(dh);

    while (stop == 0) {
        int thread_num = 0;
        while ((d = readdir(dh)) != NULL)
        {
            if (d->d_name[0] == '.')
                continue;

            thread_num++;
        }

        char str[1024];
        sprintf(str, "%d", thread_num);
        write(dev, str, strlen(str));

        seekdir(dh,pos);

        sleep(3);
    }

    stop = 0;

    closedir(dh);
    close(dev);
}

void vm_handler() {
    char path[PATH_MAX];
    snprintf(path, PATH_MAX, "/proc/%d/statm", pid);

    int dev = open(DEV_FILE, O_RDWR);

    FILE *statm = fopen(path, "r");

    char buf[BUF_SIZE + 1] = "\0";
    int len, n;

    while (stop == 0) {
        fscanf(statm, "%d", &n);

        char str[1024];
        sprintf(str, "%d MB", n * sysconf(_SC_PAGE_SIZE) / 1024 / 1024);
        write(dev, str, strlen(str));

        fseek(statm, 0, SEEK_SET);
        sleep(3);
    }

    stop = 0;

    fclose(statm);
    close(dev);
}

void state_handler() {
    char path[PATH_MAX];
    snprintf(path, PATH_MAX, "/proc/%d/stat", pid);

    int dev = open(DEV_FILE, O_RDWR);

    FILE *stat = fopen(path, "r");

    char buf[BUF_SIZE + 1] = "\0";
    int len, n;
    char name[1024];
    char state;

    while (stop == 0) {
        fscanf(stat, "%d %s %c", &n, name, &state);

        char str[1024];
        if (state == 'R') {
            sprintf(str, "%c - runnable", state);
        } else if (state == 'D') {
            sprintf(str, "%c - uninterruptable", state);
        } else if (state == 'T') {
            sprintf(str, "%c - stopped", state);
        } else if (state == 'S') {
            sprintf(str, "%c - sleeping", state);
        } else if (state == 'Z') {
            sprintf(str, "%c - zombie", state);
        } else if (state == '<') {
            sprintf(str, "%c - negative nice", state);
        } else if (state == 'N') {
            sprintf(str, "%c - positive nice", state);
        }

        write(dev, str, strlen(str));

        fseek(stat, 0, SEEK_SET);
        sleep(3);
    }

    stop = 0;

    fclose(stat);
    close(dev);
}

void comm_handler() {
    char buf[BUF_SIZE+1] = "\0";
    int len, i;
    FILE* f;

    int dev = open(DEV_FILE, O_RDWR);

    char path[PATH_MAX];
    snprintf(path, PATH_MAX, "/proc/%d/comm", pid);

    f = fopen(path, "r");

    while (stop == 0) {
        while ((len = fread(buf, 1, BUF_SIZE, f)) > 0)
        {
            buf[len] = '\0';
            write(dev, buf, len-1);
        }

        fseek(f, 0, SEEK_SET);
        sleep(3);
    }

    stop = 0;

    fclose(f);
    close(dev);
}

void signal_handler(int signal)
{
    printf("\n" ANSI_COLOR_YELLOW "caught signal = %d" ANSI_COLOR_RESET "\n", signal);
    stop = 1;
}

int main(int argc, char **argv) {
    if ((signal(SIGINT, signal_handler) == SIG_ERR)) {
        perror("Can't attach handler\n");
        return EXIT_FAILURE;
    }

    if (argc != 1) {
        printf(ANSI_COLOR_RED "no arguments needed" ANSI_COLOR_RESET "\n");
        return EXIT_FAILURE;
    }

    void (*handlers[7])(void);
    handlers[0] = &exit_handler;
    handlers[1] = &cmdline_handler;
    handlers[2] = &fds_handler;
    handlers[3] = &threads_handler;
    handlers[4] = &vm_handler;
    handlers[5] = &state_handler;
    handlers[6] = &comm_handler;

    int rc;
    while (opt >= 0) {
        rc = print_menu();
        if (rc != EXIT_SUCCESS) {
            printf(ANSI_COLOR_RED "menu error" ANSI_COLOR_RESET "\n");
            opt = 0;
            pid = 1;
            stop = 0;
        } else {
            handlers[opt]();
        }
    }

    return EXIT_SUCCESS;
}
