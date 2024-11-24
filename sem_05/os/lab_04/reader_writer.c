#include <stdlib.h>
#include <stdio.h>
#include <sys/sem.h>
#include <sys/shm.h>
#include <sys/stat.h>
#include <time.h>
#include <wait.h>
#include <unistd.h>

#define READER_AMOUNT 5
#define WRITER_AMOUNT 4
#define MAX_READER_DELAY 2000
#define MAX_WRITER_DELAY 4000
#define WRITE_ITERATIONS 4
#define READ_ITERATIONS 8

#define SEM_READERS_WAITING 0
#define SEM_WRITERS_WAITING 1
#define SEM_ACTIVE_READERS 2
#define SEM_ACTIVE_WRITERS 3
#define SEM_BIN_WRITE 4

struct sembuf start_read[4] = {
	{SEM_WRITERS_WAITING, 0, 0},
	{SEM_ACTIVE_WRITERS, 0, 0},
	{SEM_BIN_WRITE, 0, 0},
	{SEM_ACTIVE_READERS, 1, 0}
};

struct sembuf stop_read[1] = {
	{SEM_ACTIVE_READERS, -1, 0}
};

struct sembuf start_write[6] = {
	{SEM_WRITERS_WAITING, 1 , 0},
	{SEM_ACTIVE_WRITERS, 0, 0},
	{SEM_ACTIVE_WRITERS, 1, 0},
	{SEM_ACTIVE_READERS, 0, 0},
	{SEM_BIN_WRITE, 1, 0},
	{SEM_WRITERS_WAITING, -1, 0}
};

struct sembuf stop_write[2] = {
	{SEM_BIN_WRITE, -1, 0},
	{SEM_ACTIVE_WRITERS, -1, 0}
};

void sleepms(int mss)
{
	struct timespec request = {mss/1000, (mss * 1000)%1000};
	nanosleep(&request, NULL);
}

void reader_loop(const int sem_desc, const int reader_id, int *adr)
{
	if (semop(sem_desc, start_read, 4) == -1)
	{
		perror("Читатель не может изменить значение семафора\n");
		exit(1);
	}

	printf("	Читатель %d \"%d\"\n", reader_id, *adr);

	if (semop(sem_desc, stop_read, 1) == -1)
	{
		perror("Читатель не может изменить значение семафора\n");
		exit(1);
	}
	sleepms(rand() % MAX_READER_DELAY);
}

void create_reader(const int sem_desc, const int reader_id, int *adr)
{
	pid_t childpid;
	if ((childpid = fork()) == -1)
	{
		perror("Ошибка при создании читателя\n");
		exit(1);
	}
	if (childpid == 0)
	{
		srand(time(NULL) + reader_id);
		for (int i = 0; i < READ_ITERATIONS; i++)
			reader_loop(sem_desc, reader_id, adr);
		exit(0);
	}
}

void writer_loop(const int sem_desc, const int writer_id, int *adr)
{
	sleepms(rand() % MAX_WRITER_DELAY);
	if (semop(sem_desc, start_write, 6) == -1)
	{
		perror("Писатель не может изменить значение семафора\n");
		exit(1);
	}
	(*adr)++;
	printf("Писатель %d \"%d\"\n", writer_id, *adr);


	if (semop(sem_desc, stop_write, 2) == -1)
	{
		perror("Писатель не может изменить значение семафора\n");
		exit(1);
	}
}

void create_writer(const int sem_desc, const int writer_id, int *adr)
{
	pid_t childpid;
	if ((childpid = fork()) == -1)
	{
		perror("Ошибка при создании писателя\n");
		exit(1);
	}
	if (childpid == 0)
	{
		srand(time(NULL) + writer_id);
		for (int i = 0; i < WRITE_ITERATIONS; i++)
			writer_loop(sem_desc, writer_id, adr);
		exit(0);
	}
}

int main(void)
{
	int sem_desc, status, shm_id;
	int perms = S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH;
	key_t shmkey, semkey;
	int *adr;

	if ((shmkey = ftok("/", 253)) == -1)
	{
		perror("Ошибка ftok\n");
		exit(1);
	}
	if ((shm_id = shmget(shmkey, 64, IPC_CREAT | perms)) == -1)
	{
		perror("Не удалось создать разделяемый сегмент\n");
		exit(1);
	}

	if ((adr = shmat(shm_id, NULL, 0)) == (void *) -1)
	{
		perror("Не удалось получить указатель на разделяемый сегмент\n");
		exit(1);
	}

	*adr = 0;

	if ((semkey = ftok("/tmp", 123)) == -1)
	{
		perror("Ошибка ftok\n");
		exit(1);
	}
	if ((sem_desc = semget(semkey, 5, IPC_CREAT | perms)) == -1)
	{
		perror("Не удалось создать семафоры\n");
		exit(1);
	}

	for (int i = 0; i < READER_AMOUNT; i++)
		create_reader(sem_desc, i + 1, adr);
	
	for (int i = 0; i < WRITER_AMOUNT; i++)
		create_writer(sem_desc, i + 1, adr);

	for (int i = 0; i < READER_AMOUNT + WRITER_AMOUNT; i++)
		wait(&status);

	if (shmdt(adr) == -1)
	{
		perror("Ошибка при попытке отключить разделяемый сегмент от адресного пространства процесса\n");
		exit(1);
	}
	if (shmctl(shm_id, IPC_RMID, NULL))
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
