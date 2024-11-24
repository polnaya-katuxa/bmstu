#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <sys/stat.h>
#include <sys/shm.h>
#include <sys/sem.h>
#include <unistd.h>
#include <wait.h>

#define PRODUCERS_AMOUNT 3
#define CONSUMERS_AMOUNT 3
#define ITEMS_PER_PRODUCER 5
#define N PRODUCERS_AMOUNT * ITEMS_PER_PRODUCER
#define MAX_PRODUCER_DELAY 1500
#define MAX_CONSUMER_DELAY 2000

#define SEM_FULL 0
#define SEM_EMPT 1
#define SEM_BIN 2

struct sembuf init_values[2] = {
	{SEM_EMPT, N, 0},
	{SEM_BIN, 1, 0}
};

struct sembuf producer_begin[2] = {
	{SEM_BIN, -1, 0},
	{SEM_EMPT, -1 , 0}
};

struct sembuf producer_end[2] = {
	{SEM_BIN, 1, 0},
	{SEM_FULL, 1 , 0}
};

struct sembuf consumer_begin[2] = {
	{SEM_FULL, -1, 0},
	{SEM_BIN, -1, 0}
};

struct sembuf consumer_end[2] = {
	{SEM_BIN, 1, 0},
	{SEM_EMPT, 1, 0}
};


void sleepms(int mss)
{
	struct timespec request = {mss/1000, (mss * 1000) %1000};
	nanosleep(&request, NULL);
}

void produce(const int producer_id, const int sem_id, char **adr, char *symbol_adr)
{
	sleepms(rand() % MAX_PRODUCER_DELAY);
	if (semop(sem_id, producer_begin, 2) == -1)
	{
		perror("Производитель не может изменить значение семафора\n");
		exit(1);
	}
	
	printf("Производитель %d, ", producer_id);
	**adr = *symbol_adr;
	printf("\"%c\"\n", *symbol_adr);
	(*symbol_adr)++;
	(*adr)++;

	if (semop(sem_id, producer_end, 2))
	{
		perror("Производитель не может изменить значение семафора\n");
		exit(1);
	}
}

void create_producer(const int producer_id, const int sem_id, void *adr, char *symbol_adr)
{
	pid_t childpid;
	if ((childpid = fork()) == -1)
	{
		perror("Ошибка при создании процесса производителя");
		exit(1);
	}
	if (childpid == 0)
	{
		srand(time(NULL) + producer_id);
		for (int i = 0; i < ITEMS_PER_PRODUCER; i++)
			produce(producer_id, sem_id, adr, symbol_adr);
		exit(0);
	}
}

void consume(const int consumer_id, const int sem_id, char **adr)
{
	if (semop(sem_id, consumer_begin, 2) == -1)
	{
		perror("Потребитель не может изменить значение семафора\n");
		exit(1);
	}
	printf("	Потребитель %d ", consumer_id);
	printf("\"%c\"\n", **adr);
	(*adr)++;

	if (semop(sem_id, consumer_end, 2))
	{
		perror("Потребитель не может изменить значение семафора\n");
		exit(1);
	}
	sleepms(rand() % MAX_CONSUMER_DELAY);
}

void create_consumer(const int consumer_id, const int sem_id, void * adr)
{
	pid_t childpid;
	if ((childpid = fork()) == -1)
	{
		perror("Ошибка при создании процесса производителя");
		exit(1);
	}
	if (childpid == 0)
	{
		srand(time(NULL) + consumer_id);
		for (int i = 0; i < ITEMS_PER_PRODUCER; i++)
			consume(consumer_id, sem_id, adr);
		exit(0);
	}
}

int main(void)
{
	int shmid, sem_desc, status;
	int perms = S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH;
	key_t shmkey, semkey;
	void *adr = NULL;
	char **producer_pos, **consumer_pos;
	char *symbol;


	if ((shmkey = ftok("/tmp", 253)) == -1)
	{
		perror("Ошибка ftok\n");
		exit(1);
	}
	if ((shmid = shmget(shmkey, 128, IPC_CREAT | perms)) == -1)
	{
		perror("Не удалось создать разделяемый сегмент\n");
		exit(1);
	}

	if ((adr = shmat(shmid, NULL, 0)) == (void *) -1)
	{
		perror("Не удалось получить указатель на сегмент\n");
		exit(1);
	}

	producer_pos = adr;
	consumer_pos = adr + sizeof(char *);
	symbol = adr + 2 * sizeof(char *);
	*symbol = 'a';
	*producer_pos = (void *)symbol + sizeof(char);
	*consumer_pos = *producer_pos;

	if ((semkey = ftok("/tmp", 254)) == -1)
	{
		perror("Ошибка ftok\n");
		exit(1);
	}
	if ((sem_desc = semget(semkey, 3, IPC_CREAT | perms)) == -1)
	{
		perror("Не удалось создать набор семафоров\n");
		exit(1);
	}
	if (semop(sem_desc, init_values, 2) == -1)
	{
		perror("Не удалось изменить семафор\n");
		exit(1);
	}

	for (int i = 0; i < PRODUCERS_AMOUNT; i++)
		create_producer(i + 1, sem_desc, producer_pos, symbol);

	for (int i = 0; i < CONSUMERS_AMOUNT; i++)
		create_consumer(i + 1, sem_desc, consumer_pos);

	for (int i = 0; i < PRODUCERS_AMOUNT + CONSUMERS_AMOUNT; i++)
		wait(&status);

	if (shmdt(adr) == -1)
	{
		perror("Ошибка при попытке отключить разделяемый сегмент от адресного пространства процесса\n");
		exit(1);
	}
	if (shmctl(shmid, IPC_RMID, NULL))
	{
		perror("Ошибка удаления разделяемого сегмента\n");
		exit(1);
	}
	if (semctl(sem_desc, 0, IPC_RMID) == -1)
	{
		perror("Ошибка при попытке удаления семафора\n");
		exit(1);
	}

	return 0;
}
