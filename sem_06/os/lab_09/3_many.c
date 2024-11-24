#include <fcntl.h>
#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

#define FILENAME "tmp2.txt"

void* thread1()
{
  FILE* f = fopen(FILENAME, "w");

  for (char c = 'a'; c <= 'z'; c += 2)
  {
    fprintf(f, "%c", c);
  }

  fclose(f);

  return NULL;
}

void* thread2() 
{
  FILE* f = fopen(FILENAME, "w");

  for (char c = 'b'; c <= 'z'; c += 2)
  { 
    fprintf(f, "%c", c);
  }

  fclose(f);

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