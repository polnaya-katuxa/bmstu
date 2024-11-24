#include <fcntl.h>
#include <stdio.h> 
#include <sys/types.h> 
#include <sys/stat.h>
#include <unistd.h>

#define FILENAME "tmp3.txt"

int main()
{
	struct stat buf;

	int f1 = open(FILENAME, O_WRONLY | O_CREAT, S_IRUSR | S_IWUSR);

	int rc = fstat(f1, &buf);
	if (!rc)
	{
		fprintf(stdout, "after open f1: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);
	}

	int f2 = open(FILENAME, O_WRONLY | O_CREAT, S_IRUSR | S_IWUSR);

	rc = fstat(f2, &buf);
	if (!rc)
	{
		fprintf(stdout, "after open f2: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);
	}

	for (char c = 'a'; c <= 'z'; c++)
	{
		if (c % 2)
		{
			write(f2, &c, 1);
		}
		else
		{
			write(f1, &c, 1); 
		}
	}

	rc = fstat(f1, &buf);
	if (!rc)
	{
		fprintf(stdout, "before close f1: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);
	}

	rc = fstat(f2, &buf);
	if (!rc)
	{
		fprintf(stdout, "before close f2: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);
	}

	close(f1);

	rc = stat(FILENAME, &buf);
	fprintf(stdout, "after close f1: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);

	close(f2);

	rc = stat(FILENAME, &buf);
	fprintf(stdout, "after fclose f2: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);

	return 0;
}