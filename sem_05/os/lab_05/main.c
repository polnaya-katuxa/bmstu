#include <windows.h>
#include <stdbool.h>
#include <stdio.h>
#include <time.h>
#include <stdbool.h>

#define MINIMUM_READER_DELAY 120
#define MINIMUM_WRITER_DELAY 120
#define MAXIMUM_READER_DELAY 300
#define MAXIMUM_WRITER_DELAY 450

#define READERS_NUMBER 5
#define WRITERS_NUMBER 3

#define COUNT 8

HANDLE can_read;
HANDLE can_write;
HANDLE mutex;

LONG waiting_writers_count = 0;
LONG waiting_readers_count = 0;
LONG active_readers_count = 0;

bool is_writer_active  = false;

HANDLE readerThreads[READERS_NUMBER];
HANDLE writerThreads[WRITERS_NUMBER];

int readersID[READERS_NUMBER];
int writersID[WRITERS_NUMBER];

int readersRand[READERS_NUMBER * COUNT];
int writersRand[READERS_NUMBER * COUNT];

int value = 0;

void StartRead()
{
	InterlockedIncrement(&waiting_readers_count);

	if (is_writer_active  || WaitForSingleObject(can_write, 0) == WAIT_OBJECT_0)
	{
		WaitForSingleObject(can_read, INFINITE);
	}

	WaitForSingleObject(mutex, INFINITE);
	
	InterlockedDecrement(&waiting_readers_count);
	
	InterlockedIncrement(&active_readers_count);
	
	SetEvent(can_read);
	
	ReleaseMutex(mutex);
}

void StopRead()
{
	InterlockedDecrement(&active_readers_count);

	if (!active_readers_count)
	{
		SetEvent(can_write);
	}
}

DWORD WINAPI Reader(CONST LPVOID param)
{
	int id = *(int *)param;
	int sleepTime;
	int begin = id * COUNT;
	for (int i = 0; i < COUNT; i++)
	{
		sleepTime = rand() % MAXIMUM_READER_DELAY + MINIMUM_READER_DELAY;
		Sleep(sleepTime);
		
		StartRead();
		printf("Reader id = %d; value = %d; sleep time = %d.\n", id, value, sleepTime);
		StopRead();
	}
}

void StartWrite()
{
	InterlockedIncrement(&waiting_writers_count);

	if (is_writer_active  || active_readers_count > 0)
	{
		WaitForSingleObject(can_write, INFINITE);
	}

	InterlockedDecrement(&waiting_writers_count);

	is_writer_active  = true;
	ResetEvent(can_write);
}

void StopWrite()
{
	is_writer_active  = false;
	
	if (waiting_readers_count)
	{
		SetEvent(can_read);
	}
	else
	{
		SetEvent(can_write);
	}		
}

DWORD WINAPI Writer(CONST LPVOID param)
{
	int id = *(int *)param;
	int sleepTime;
	int begin = id * COUNT;
	
	for (int i = 0; i < COUNT; i++)
	{
		sleepTime = rand() % MAXIMUM_WRITER_DELAY + MINIMUM_WRITER_DELAY;
		Sleep(sleepTime);

		StartWrite();
		++value;
		printf("Writer id = %d; value = %d; sleep time = %d.\n", id, value, sleepTime);
		StopWrite();
	}
}

int init()
{
	mutex = CreateMutex(NULL, FALSE, NULL);
	if (mutex == NULL)
	{
		perror("Can't create mutex\n");
		return -1;
	}

	if ((can_write = CreateEvent(NULL, TRUE, FALSE, NULL)) == NULL)
	{
		perror("CreateEvent (can_write)");
		return -1;
	}
	
	if ((can_read = CreateEvent(NULL, FALSE, FALSE, NULL)) == NULL)
	{
		perror("CreateEvent (can_read)");
		return -1;
	}
	
	return 0;
}

int CreateThreads()
{
	for (short i = 0; i < WRITERS_NUMBER; i++)
	{
		writersID[i] = i;
		if ((writerThreads[i] = CreateThread(NULL, 0, &Writer, writersID + i, 0, NULL)) == NULL)
		{
			perror("Can't create thread (writer)");
			return -1;
		}
	}
	
	for (short i = 0; i < READERS_NUMBER; i++)
	{
		readersID[i] = i;
		if ((readerThreads[i] = CreateThread(NULL, 0, &Reader, readersID + i, 0, NULL)) == NULL)
		{
			perror("Can't create thread (reader)");
			return -1;
		}
	}

	return 0;
}

void Close()
{
	for (int i = 0; i < READERS_NUMBER; i++)
	{
		CloseHandle(readerThreads[i]);
	}

	for (int i = 0; i < WRITERS_NUMBER; i++)
	{
		CloseHandle(writerThreads[i]);
	}

	CloseHandle(can_read);
	CloseHandle(can_write);
	CloseHandle(mutex);
}

int main(void)
{
	setbuf(stdout, NULL);
	srand(time(NULL));


	int rc = init();
	if (rc)
	{
		return -1;
	}

	rc = CreateThreads();
	if (rc)
	{
		return -1;
	}

	WaitForMultipleObjects(WRITERS_NUMBER, writerThreads, TRUE, INFINITE);
	WaitForMultipleObjects(READERS_NUMBER, readerThreads, TRUE, INFINITE);


	Close();

	printf("\nSUCCESS\n");
	return 0;
}
