#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include <sys/wait.h>
#include <sys/types.h>

#define CHILD_NUM 2

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
    pid_t child_pid[CHILD_NUM];
    char* prog_rel_path[CHILD_NUM] = {"./progs/levenshtein", "./progs/reg_expression"};
    int status;

    for (size_t i = 0; i < CHILD_NUM; i++) 
    {
        if ((child_pid[i] = fork()) == -1)
        {
            perror("Can't fork\n");

            exit(1);
        }   
        if (child_pid[i] == 0)
        {
            if (execlp(prog_rel_path[i], NULL) == -1)
            {
                perror("\nError: Child can't execute exec()\n");

                exit(1);
            }
        }
        printf("\nParent: pid = %d, pgrp = %d, child = %d\n", getpid(), getpgrp(), child_pid[i]);
    }

    for (size_t i = 0; i < CHILD_NUM; i++) 
    {
        if (waitpid(child_pid[i], &status, WUNTRACED) == -1)
        {
            perror("Wait error\n");

            exit(1);
        }
        printf("\n\nProcess status: %d, child pid = %d\n", status, child_pid[i]);
        check_status(status);
    }

    return 0;
}
