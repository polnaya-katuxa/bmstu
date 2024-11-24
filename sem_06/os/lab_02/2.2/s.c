#include <sys/types.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <unistd.h>
#include <string.h>

#define BUF_SIZE 50

int sockfd;

void signal_handler(int signal)
{
    printf("\nCaught signal = %d\n", signal);
    unlink("socket.soc");
    close(sockfd);
    printf("\nServer exiting.\n");
    exit(0);
}

int main(int argc, char **argv)
{
    if ((signal(SIGINT, signal_handler) == SIG_ERR)) 
    {
        perror("Can't attach handler\n");
        return EXIT_FAILURE;
    }

    char buf[BUF_SIZE];
    sockfd = socket(AF_UNIX, SOCK_DGRAM, 0);

    if (sockfd == -1)
    {
        perror("socket() failed");
        return EXIT_FAILURE;
    }

    struct sockaddr sa;
    sa.sa_family = AF_UNIX;
    strcpy(sa.sa_data, "socket.soc");

    if (bind(sockfd, &sa, sizeof(sa)) == -1) 
    {
        perror("bind() failed");
        unlink("socket.soc");
        close(sockfd);
        return EXIT_FAILURE;
    }

    int bytes;
    while (1) 
    {
        struct sockaddr ca;
        socklen_t len = sizeof(ca);

        bytes = recvfrom(sockfd, buf, sizeof(buf), 0, &ca, &len);
        if (bytes == -1) 
        {
            perror("recvfrom() failed");
            unlink("socket.soc");
            close(sockfd);
            return EXIT_FAILURE;
        }
        printf("\nreceived: %s\n", buf);
        sprintf(buf, "%s %d", buf, getpid());
        
        if (sendto(sockfd, buf, sizeof(buf), 0, &ca, len) == -1)
        {
            perror("sendto() failed");
            unlink("socket.soc");
            close(sockfd);
            return EXIT_FAILURE;
        }
        printf("sent: %s\n", buf);
    }
    return EXIT_SUCCESS;
}