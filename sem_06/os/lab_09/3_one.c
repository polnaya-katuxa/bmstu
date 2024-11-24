#include <fcntl.h>
#include <stdio.h> 
#include <sys/types.h> 
#include <sys/stat.h>
#include <unistd.h>

#define FILENAME "tmp1.txt"

int main()
{
	struct stat buf;

	FILE* f1 = fopen(FILENAME, "w");

	int rc = fstat(f1->_file, &buf);
	if (!rc)
	{
		fprintf(stdout, "after fopen f1: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);
	}

	FILE* f2 = fopen(FILENAME, "w");

	rc = fstat(f2->_file, &buf);
	if (!rc)
	{
		fprintf(stdout, "after fopen f2: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);
	}

	for (char c = 'a'; c <= 'z'; c++)
	{
		if (c % 2)
		{
			fprintf(f1, "%c", c);
		}
		else
		{
			fprintf(f2, "%c", c); 
		}
	}

	rc = fstat(f1->_file, &buf);
	if (!rc)
	{
		fprintf(stdout, "before fclose f1: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);
	}

	rc = fstat(f2->_file, &buf);
	if (!rc)
	{
		fprintf(stdout, "before fclose f2: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);
	}

	fclose(f1);

	rc = stat(FILENAME, &buf);
	fprintf(stdout, "after fclose f1: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);

	fclose(f2);

	rc = stat(FILENAME, &buf);
	fprintf(stdout, "after fclose f2: inode - %ju, total size - %lld\n", (uintmax_t)buf.st_ino, buf.st_size);

	return 0;
}