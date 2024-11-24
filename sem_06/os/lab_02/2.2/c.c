#include <sys/types.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <unistd.h>
#include <string.h>
#include <sys/un.h>

#define BUF_SIZE 50

int main(int argc, char **argv)
{
    char buf[BUF_SIZE];

    if (argc != 2)
    {
        perror("args not enough\n");
        return EXIT_FAILURE;
    }
    sprintf(buf, "%d %s", getpid(), argv[1]);

    int sockfd = socket(AF_UNIX, SOCK_DGRAM, 0);
    if (sockfd == -1)
    {
        perror("socket() failed\n");
        return EXIT_FAILURE;
    }

    struct sockaddr sa;
    sa.sa_family = AF_UNIX;
    strcpy(sa.sa_data, "socket.soc");
    socklen_t len = sizeof(sa);

    char name[20];
    sprintf(name, "%d.soc", getpid());

    struct sockaddr ca;
    ca.sa_family = AF_UNIX;
    strcpy(ca.sa_data, name);
    if (bind(sockfd, &ca, sizeof(ca)) == -1) 
    {
        perror("bind() failed\n");
        close(sockfd);
        return EXIT_FAILURE;
    }

    if (sendto(sockfd, buf, sizeof(buf), 0, &sa, len) == -1)
    {
        perror("sendto() failed\n");
        unlink(name);
        close(sockfd);
        return EXIT_FAILURE;
    }
    if (recvfrom(sockfd, buf, sizeof(buf), 0, NULL, NULL) == -1) //&sa, &len
    {
        perror("recvfrom() failed\n");
        unlink(name);
        close(sockfd);
        return EXIT_FAILURE;
    }
    printf("\nreceived: %s\n", buf);
    unlink(name);
    close(sockfd);
    return EXIT_SUCCESS;
}