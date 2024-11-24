#include <sys/types.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <unistd.h>

#define BUF_SIZE 50
#define CHILD_NUM 3

int main(int argc, char ** argv)
{
    int fd[2];
    char buf[BUF_SIZE];
    pid_t child_pid[CHILD_NUM];

    if (socketpair(AF_UNIX, SOCK_DGRAM, 0, fd) < 0) 
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
            //close(fd[1]);
            sprintf(buf, "%d", getpid());
            write(fd[0], buf, sizeof(buf));
            printf("Child wrote %s\n", buf);
            //sleep(1);
            read(fd[0], buf, sizeof(buf));
            printf("Child %d read %s\n", getpid(), buf);
           // close(fd[0]);

            return EXIT_SUCCESS;
        }
        else 
        {
            //close(fd[0]);
            read(fd[1], buf, sizeof(buf));
            printf("Parent read %s\n", buf);
            sprintf(buf, "%s %d", buf, getpid());
            write(fd[1], buf, sizeof(buf));
            printf("Parent wrote %s\n", buf);
        }
    }

    //close(fd[0]);
    //close(fd[1]);

    return EXIT_SUCCESS;
}