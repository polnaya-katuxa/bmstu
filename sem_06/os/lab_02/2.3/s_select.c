#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <unistd.h>
#include <string.h>
#include <fcntl.h>

#define BUF_SIZE 50
#define CLIENT_NUM 5

int sockfd;

void signal_handler(int signal)
{
    printf("\nCaught signal = %d\n", signal);
    close(sockfd);
    printf("\nServer exiting.\n");
    exit(0);
}

int main(int argc, char** argv) 
{
	if (argc != 2)
    {
        perror("args not enough\n");
        return EXIT_FAILURE;
    }
    short port = (short) atoi(argv[1]);

	if ((signal(SIGINT, signal_handler) == SIG_ERR)) {
        perror("Can't attach handler\n");
        return EXIT_FAILURE;
    }

	sockfd = socket(AF_INET, SOCK_STREAM, 0);

	if (sockfd == -1) 
	{
		perror("socket() failed");
        return EXIT_FAILURE;
	}

	fcntl(sockfd, F_SETFL, O_NONBLOCK);

	struct sockaddr_in sa;
	sa.sin_family = AF_INET;
	sa.sin_addr.s_addr = INADDR_ANY;
	sa.sin_port = htons(port);

	if (bind(sockfd, (struct sockaddr*) &sa, sizeof(sa)) < 0) 
	{
		perror("bind() failed");
		return EXIT_FAILURE;
	}

	listen(sockfd, CLIENT_NUM);

	char buf[BUF_SIZE];
	int k = 0;
	int newsockfd[CLIENT_NUM];
	struct timeval tv;
	tv.tv_sec = 5;
	tv.tv_usec = 0;

	while (1)
	{
		fd_set set;
		FD_ZERO(&set);
		FD_SET(sockfd, &set);
		int maxfd = sockfd;

		for (int i = 0; i < k; i++)
		{
			FD_SET(newsockfd[i], &set);
			if (newsockfd[i] > maxfd)
			{
				maxfd = newsockfd[i];
			}
		}

		select(maxfd + 1, &set, NULL, NULL, &tv);
		if (FD_ISSET(sockfd, &set)) 
		{
			struct sockaddr_in sa;
			int addrSize = sizeof(sa);

			newsockfd[k] = accept(sockfd, (struct sockaddr*) &sa, (socklen_t*) &addrSize);
			if (newsockfd[k] < 0)
			{
			    perror("accept() failed");
			    return EXIT_FAILURE;
			}

			k++;
			    
			printf("\nClient connected №%d %s:%d\n", 
                        k - 1, inet_ntoa(sa.sin_addr), ntohs(sa.sin_port));
		} 
		else
		{
			for (int i = 0; i < k; i++)
			{
				if (FD_ISSET(newsockfd[i], &set))
				{
					if (read(newsockfd[i], buf, sizeof(buf)) == 0)
                    {
                        printf("\nclient %d disconnected\n", i);
                        newsockfd[i] = 0;
                        continue;
                    }
					printf("\nreceived message: %s\n", buf);
					write(newsockfd[i], buf, sizeof(buf));
					printf("\nsent message: %s\n", buf);
				}
			}
		}
	}

	return EXIT_SUCCESS;
}