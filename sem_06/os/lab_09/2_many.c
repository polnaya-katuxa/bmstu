#include <fcntl.h>
#include <pthread.h>
#include <unistd.h>
#include <stdio.h>

#define BUF_SIZE 20
#define FILENAME "alphabet.txt"

void* thread1()
{
  int fd = open(FILENAME, O_RDONLY);

  char c;
  while ((read(fd, &c, 1)) == 1) write(1, &c, 1);

  return NULL;
}

void* thread2()
{
  int fd = open(FILENAME, O_RDONLY);

  char c;
  while ((read(fd, &c, 1)) == 1) write(1, &c, 1);

  return NULL;
}

int main()
{
  pthread_t t1, t2;

  pthread_create(&t1, NULL, thread1, NULL);
  pthread_create(&t2, NULL, thread2, NULL);

  pthread_join(t1, NULL);
  pthread_join(t2, NULL);

  return 0;
}