/*
 * This is sample code generated by rpcgen.
 * These are only templates and you can use them
 * as a guideline for developing your own functions.
 */
 
#include <pthread.h>

#include "bakery.h"

#define N 50

int choosing[N] = { 0 };
int num[N] = { 0 };
int id = 0;
int data = 'a';

struct arg_t
{
	int id;
	int num;
};

void get_num(struct arg_t *arg)
{	
	sleep(rand() % 3);
	
	while (choosing[id]);
	
	choosing[id] = 1;
	int i = id;
	id++;
	arg->id = i;
	int max = 0;
	for (int j = 0; j < N; j++)
		if (num[j] > max)
			max = num[j];
	num[i] = max + 1;
	arg->num = num[i];
	choosing[i] = 0;	
}

int lexical_less(int a1, int a2, int b1, int b2) {
	if (a1 < b1) 
		return 1;
	if (a1 > b1)
		return 0;
	return a2 < b2;
}

int get_data(struct arg_t *arg)
{
	sleep(rand() % 3);
	int i = arg->id;
	for (int j = 0; j < N; j++)
	{
		while (choosing[j]);
		while (num[j] != 0 && lexical_less(num[i], i, num[j], j));
	}
	int d = data;
	data++;
	//num[i] = 0;
	return d;
}

bool_t
bakery_proc_1_svc(struct BAKERY *argp, struct BAKERY *result, struct svc_req *rqstp)
{
	bool_t retval;

	struct arg_t arg;


	switch (argp->op)
	{
		case GET_NUM:
		{
			get_num(&arg);
			result->id = arg.id;
			result->num = arg.num;
			break;
		}
		case OPEN_CRIT_SECTION:
		{
			arg.id = argp->id;
			result->res = get_data(&arg);
			break;
		}
	}

	return retval;
}

int
bakery_prog_1_freeresult (SVCXPRT *transp, xdrproc_t xdr_result, caddr_t result)
{
	xdr_free (xdr_result, result);

	/*
	 * Insert additional freeing code here, if needed
	 */

	return 1;
}