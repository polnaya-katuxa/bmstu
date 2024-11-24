#include <sys/wait.h>
#include <stdlib.h>
#include <sys/shm.h>
#include <sys/sem.h>
#include <stdio.h>
#include <sys/stat.h>
#include <unistd.h>

#define WRITER_CNT 3
#define READER_CNT 4

#define ACTIVE_READERS 0
#define WRITE_QUEUE 1
#define READ_QUEUE 2
#define ACTIVE_WRITER 3
#define BIN_WRITER 4

int init_monitor(int semid)
{
  if (semctl(semid, ACTIVE_READERS, SETVAL, 0) == -1)
    return EXIT_FAILURE;
  if (semctl(semid, WRITE_QUEUE, SETVAL, 0) == -1)
    return EXIT_FAILURE;
  if (semctl(semid, READ_QUEUE, SETVAL, 0) == -1)
    return EXIT_FAILURE;
  if (semctl(semid, ACTIVE_WRITER, SETVAL, 0) == -1)
    return EXIT_FAILURE;
  if (semctl(semid, BIN_WRITER, SETVAL, 1) == -1)
    return EXIT_FAILURE;
  return EXIT_SUCCESS;
}

struct sembuf start_read_sbuf[] = {
  {READ_QUEUE, 1, 0},
  {ACTIVE_WRITER, 0, 0},
  {WRITE_QUEUE, 0, 0},
  {ACTIVE_READERS, 1, 0},
  {READ_QUEUE, -1, 0},
};
struct sembuf stop_read_sbuf[] = {
  {ACTIVE_READERS, -1, 0},
};
struct sembuf start_write_sbuf[] = {
  {WRITE_QUEUE, 1, 0},
  {ACTIVE_READERS, 0, 0},
  {ACTIVE_WRITER, 0, 0},
};
struct sembuf continue_write_sbuf[] = {
  {ACTIVE_WRITER, 1, 0},
  {WRITE_QUEUE, -1, 0},
  {BIN_WRITER, -1, 0},  
};
struct sembuf stop_write_sbuf[] = {
  {ACTIVE_WRITER, -1, 0},
  {BIN_WRITER, 1, 0},
};

int start_read(int semid)
{
  return semop(semid, start_read_sbuf, 5);
}
int stop_read(int semid)
{
  return semop(semid, stop_read_sbuf, 1);
}
int start_write(int semid)
{
  if (semop(semid, start_write_sbuf, 3) == -1)
    return -1;
  return semop(semid, continue_write_sbuf, 3);
}
int stop_write(int semid)
{
  return semop(semid, stop_write_sbuf, 2);
}

int writer(int semid, int *number)
{
  srand(getpid());
  for (int i = 0; i < 4; i++)
  {
    usleep(rand() % 80000);

    if (start_write(semid) == -1)
    {
      perror("cant start_write");
      shmdt(number);
      return EXIT_FAILURE;
    }
    (*number)++;
    printf("writer incremented, current value = \'%d\'\n", *number);
    if (stop_write(semid) == -1)
    {
      perror("cant stop_write");
      shmdt(number);
      return EXIT_FAILURE;
    }
  }
  if (shmdt(number) == -1)
        return EXIT_FAILURE;
  return EXIT_SUCCESS;
}

int reader(int semid, int *number)
{
  srand(getpid());
  for (int i = 0; i < 4; i++)
  {
    usleep(rand() % 100000);

    if (start_read(semid) == -1)
    {
      perror("cant start_read");
      shmdt(number);
      return EXIT_FAILURE;
    }
    printf("reader got value = \'%d\'\n", *number);
    if (stop_read(semid) == -1)
    {
      perror("cant stop_read");
      shmdt(number);
      return EXIT_FAILURE;
    }
  }
  if (shmdt(number) == -1)
        return EXIT_FAILURE;
  return EXIT_SUCCESS;
}

void log_status_info(pid_t pid, int status) {
  if (WIFEXITED(status)) 
    printf("child with pid %d has finished, code: %d\n", pid, WEXITSTATUS(status));
  else if (WIFSIGNALED(status))
    printf("child with pid %d has finished by unhandlable signal, signum: %d\n", pid, WTERMSIG(status));
  else if (WIFSTOPPED(status))
    printf("child with pid %d has finished by signal, signum: %d\n", pid, WSTOPSIG(status));
}

void log_exit(const char *msg)
{
  perror(msg);
  exit(EXIT_FAILURE);
}

int main(void) 
{
  const int perms = S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH;
  int *number;
  int shmid;
  int semid;

  key_t semkey = ftok("/", 1);
  if (semkey == -1)
    log_exit("cant get key");
  if ((semid = semget(semkey, 5, IPC_CREAT | perms)) == -1)
    log_exit("cant semget");
  key_t shmkey = ftok("/", 2);
  if (shmkey == -1)
    log_exit("cant get key");
  shmid = shmget(shmkey, sizeof(int), IPC_CREAT | perms);
  if (shmid == -1)
    log_exit("cant shmget");
  number = shmat(shmid, 0, perms);
  if (number == (int *)-1)
    log_exit("cant allocate_number");
  *number = 0;

  if (init_monitor(semid))
    log_exit("cant init_monitor");

  pid_t chpid[WRITER_CNT + READER_CNT];
  for (int i = 0; i < WRITER_CNT; i++) 
  {
    if ((chpid[i] = fork()) == -1)
      log_exit("cant fork");
    else if (chpid[i] == 0)
      {
        int *number = shmat(shmid, 0, perms);
        if (number == (int *)-1)
        {
          perror("cant shmat");
          exit(EXIT_FAILURE);
        }
        if (writer(semid, number))
        {
          perror("cant write");
          exit(EXIT_FAILURE);
        }
        else
          exit(EXIT_SUCCESS);
      }
  }

  for (int i = WRITER_CNT; i < READER_CNT + WRITER_CNT; i++) 
  {
    if ((chpid[i] = fork()) == -1)
      log_exit("cant fork");
    else if (chpid[i] == 0)
    {
      int *number = shmat(shmid, 0, perms);
      if (number == (int *)-1)
      {
        perror("cant shmat");
        exit(EXIT_FAILURE);
      }
      if (reader(semid, number))
      {
        perror("cant read");
        exit(EXIT_FAILURE);
      }
      else
        exit(EXIT_SUCCESS);
    }       
  }

  if (writer(semid, number))
  {
    perror("cant write");
    exit(EXIT_FAILURE);
  }

  for (int i = 0; i < READER_CNT + WRITER_CNT; i++)
  {
    int status;
    if (waitpid(chpid[i], &status, WUNTRACED) == -1)
      log_exit("cant wait");
    log_status_info(chpid[i], status);
  }

  if (semctl(semid, 1, IPC_RMID, NULL) == -1)
    log_exit("cant semctl");
  // if (shmdt(number) == -1)
  //   log_exit("cant shmdt");
  if (shmctl(shmid, IPC_RMID, NULL) == -1)
    log_exit("cant shmctl");

  return EXIT_SUCCESS;
}
