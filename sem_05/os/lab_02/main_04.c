#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include <sys/wait.h>
#include <sys/types.h>

#include <string.h>

#define CHILD_NUM 2
#define TEXT_LEN 50

void check_status(int status)
{
    if (WIFEXITED(status))
    {
        printf("Сhild process has finished correctly\n");
        printf("Child process has finished with code: \t%d\n\n", WEXITSTATUS(status));
    }
    else if (WIFSIGNALED(status))
    {
        printf("Child process has finished with non-interceptable signal\n");
        printf("Signal: \t%d\n\n", WTERMSIG(status));
    }
    else if (WIFSTOPPED(status))
    {
        printf("Child process has stopped\n");
        printf("Signal: \t%d\n\n", WSTOPSIG(status));
    }
}

int main(void)
{
    char msg[CHILD_NUM][TEXT_LEN] = {"skdjf;WEJ;forJEFKNWEEFFIEJOekfjwkeG", "wjdfne"};
    pid_t child_pid[CHILD_NUM];
    int fd[2];
    int status;

    if (pipe(fd) == -1)
    {
        perror("Can't pipe\n");

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
            close(fd[0]);
            write(fd[1], msg[i], strlen(msg[i]) + 1);
            printf("Child %d sent message: %s\n", getpid(), msg[i]);

            return 0;
        }
    }

    for (size_t i = 0; i < CHILD_NUM; i++) 
    {
        close(fd[1]);
        read(fd[0], msg[i], strlen(msg[i]) + 1);
        if (waitpid(child_pid[i], &status, WUNTRACED) == -1)
        {
            perror("Wait error\n");

            exit(1);
        }
        check_status(status);
        printf("Parent %d recieved message: %s from child %d\n", getpid(), msg[i], child_pid[i]);
    }

    return 0;
}
