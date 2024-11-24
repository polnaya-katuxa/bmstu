#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define CHILD_NUM 2
#define SLEEP_TIME 2

int main(void)
{   
    pid_t child_pid[CHILD_NUM];

    for (size_t i = 0; i < CHILD_NUM; i++) 
    {
        if ((child_pid[i] = fork()) == -1)
        {
            perror("Can't fork\n");

            exit(1);
        }
        if (child_pid[i] == 0)
        {   
            printf("\nBefore sleep: Child: pid = %d, ppid = %d, pgrp = %d\n", getpid(), getppid(), getpgrp());

            sleep(SLEEP_TIME);

            printf("\nAfter sleep: Child: pid = %d, ppid = %d, pgrp = %d\n", getpid(), getppid(), getpgrp());

            return 0;
        }
        printf("\n\nParent: pid = %d, pgrp = %d, child = %d\n\n", getpid(), getpgrp(), child_pid[i]);
    }

    return 0;
}
