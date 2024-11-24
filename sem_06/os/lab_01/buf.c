#include <sys/types.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <unistd.h>

//#define BUF_SIZE 150
#define CHILD_NUM 3
#define TEXT_LEN 50

int main(int argc, char ** argv)
{
    int fd[2];
    char buf[TEXT_LEN];
    pid_t child_pid[CHILD_NUM];

    if (socketpair(AF_UNIX, SOCK_STREAM, 0, fd) < 0) 
    {
        perror("socketpair() failed");
        exit(1);
    }

    for (size_t i = 0; i < CHILD_NUM; i++) 
    {
        if ((child_pid[i] = fork()) == -1)
        {
            perror("Can't fork\n");
            exit(1);
        }
        if (child_pid[i] == 0)
        {
            close(fd[1]);
            sprintf(buf, "%d", getpid());
            write(fd[0], buf, sizeof(buf));
            printf("Child wrote msg %s\n", buf);
            close(fd[0]);

            return EXIT_SUCCESS;
        }
    }

    sleep(1);

    for (size_t i = 0; i < CHILD_NUM; i++)
    {
        read(fd[1], buf, sizeof(buf));
        printf("Parent read msg %s\n", buf);
    }


    close(fd[0]);
    close(fd[1]);

    return EXIT_SUCCESS;
}