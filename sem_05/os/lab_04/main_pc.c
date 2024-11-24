#include <stdlib.h>
#include <sys/shm.h>
#include <sys/stat.h>
#include <sys/wait.h>
#include <limits.h>
#include <time.h>
#include <stdio.h>
#include <sys/sem.h>
#include <unistd.h>

#define PROD_CNT 3
#define CONS_CNT 3

#define SEM_BIN 0
#define SEM_FULL 1
#define SEM_EMPTY 2

struct sembuf scons_sbuf[] = {
  {SEM_BIN, -1, 0},
  {SEM_FULL, -1, 0},
};
struct sembuf econs_sbuf[] = {
  {SEM_BIN, 1, 0},
  {SEM_EMPTY, 1, 0},
};

int consumer(int semid, char **caddr, void *buf) 
{
  srand(getpid());
  for (int i = 0; i < 4; i++)
  {
    usleep(rand() % 100000);
    char sym;
    if (semop(semid, scons_sbuf, 2) == -1) {
      perror("cant semop");
      if (shmdt(buf) == -1)
        return EXIT_FAILURE;
      return EXIT_FAILURE;
    }
    printf("consumer consumed \'%c\'\n", **caddr);
    (*caddr)++;
    if (semop(semid, econs_sbuf, 2) == -1) {
      perror("cant semop");
      if (shmdt(buf) == -1)
        return EXIT_FAILURE;
      return EXIT_FAILURE;
    }
  }
  if (shmdt(buf) == -1)
        return EXIT_FAILURE;
  return EXIT_SUCCESS;
}

struct sembuf sprod_sbuf[] = {
  {SEM_BIN, -1, 0},
  {SEM_EMPTY, -1, 0},
};
struct sembuf eprod_sbuf[] = {
  {SEM_BIN, 1, 0},
  {SEM_FULL, 1, 0},
};

int producer(int semid, char **paddr, char *alpha, void *buf)
{
  srand(getpid());
  for (int i = 0; i < 4; i++)
  {
    usleep(rand() % 80000);
    if (semop(semid, sprod_sbuf, 2) == -1) {
      perror("cant semop");
      if (shmdt(buf) == -1)
        return EXIT_FAILURE;
      return EXIT_FAILURE;
    }

    **paddr = *alpha;
    printf("producer produced \'%c\'\n", **paddr);
    (*paddr)++;
    (*alpha)++;
    if (semop(semid, eprod_sbuf, 2) == -1) {
      perror("cant semop");
      if (shmdt(buf) == -1)
        return EXIT_FAILURE;
      return EXIT_FAILURE;
    }
  }
  if (shmdt(buf) == -1)
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
  int shmid;
  int perms = S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH;

  void* buf;
  char** paddr;
  char** caddr;
  char* alpha;

  key_t shmkey = ftok("/tmp", 1);
  if (shmkey == -1)
    log_exit("cant ftok");
  if ((shmid = shmget(shmkey, (3 + 1024) * sizeof(char), IPC_CREAT | perms)) == -1)
    log_exit("cant allocate_buffer");

  buf = shmat(shmid, 0, perms);
  if (buf == (char *)-1)
    log_exit("can't attach shared memory");
  
  paddr = buf;
  caddr = buf + sizeof(char*);
  alpha = buf + 2 * sizeof(char*);
  *alpha = 'a';

  *paddr = (void*) alpha + sizeof(char);
  *caddr = *paddr;

  key_t semkey = ftok("/tmp", 2);
  if (semkey == -1)
    log_exit("cant ftok");
  int semid = semget(semkey, 3, IPC_CREAT | perms);
  if (semid == -1)
    log_exit("cant semget");

  if (semctl(semid, SEM_BIN, SETVAL, 1) == -1)
    log_exit("cant semctl");
  if (semctl(semid, SEM_EMPTY, SETVAL, 1024) == -1)
    log_exit("cant semctl");
  if (semctl(semid, SEM_FULL, SETVAL, 0) == -1)
    log_exit("cant semctl");

  pid_t chpid[PROD_CNT + CONS_CNT];
  for (int i = 0; i < PROD_CNT; i++)
  {
    if ((chpid[i] = fork()) == -1)
      log_exit("cant fork");
    if (chpid[i] == 0) 
    {
      void *buf = shmat(shmid, 0, perms);
      if (buf == (char *)-1)
      { 
        perror("cant shmat");
        exit(EXIT_FAILURE);
      }
      if (producer(semid, paddr, alpha, buf))
      { 
        perror("cant produce");
        exit(EXIT_FAILURE);
      }
      else
      {
        exit(EXIT_SUCCESS);
      }
    }
  }

  if (producer(semid, paddr, alpha, buf))
  { 
      perror("cant produce");
      exit(EXIT_FAILURE);
  }

  for (int i = PROD_CNT; i < CONS_CNT + PROD_CNT; i++)
  {
    if ((chpid[i] = fork()) == -1)
      log_exit("cant fork");
    if (chpid[i] == 0)
    {
      void *buf = shmat(shmid, 0, perms);
      if (buf == (char *)-1)
      { 
        perror("cant shmat");
        exit(EXIT_FAILURE);
      }
      if (consumer(semid, caddr, buf))
      { 
        perror("cant consume");
        exit(EXIT_FAILURE);
      }
      else
      {
        exit(EXIT_SUCCESS);
      }
    }
  }

  for (int i = 0; i < PROD_CNT + CONS_CNT; i++)
  {
    int status;
    if (waitpid(chpid[i], &status, WUNTRACED) == -1)
      log_exit("cant wait");
    log_status_info(chpid[i], status);
  }
  
  if (semctl(semid, 1, IPC_RMID, NULL) == -1)
    log_exit("cant free sem");
  if (shmctl(shmid, IPC_RMID, NULL) == -1)
    log_exit("cant free sem");
  return EXIT_SUCCESS;
}
