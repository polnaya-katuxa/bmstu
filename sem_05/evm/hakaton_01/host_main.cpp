#include <iostream>
#include <stdio.h>
#include <stdexcept>
#include <iomanip>
#ifdef _WINDOWS
#include <io.h>
#else
#include <unistd.h>
#endif


#include "experimental/xrt_device.h"
#include "experimental/xrt_kernel.h"
#include "experimental/xrt_bo.h"
#include "experimental/xrt_ini.h"

#include "gpc_defs.h"
#include "leonhardx64_xrt.h"
#include "gpc_handlers.h"

#define BURST 256

union uint64 {
    uint64_t 	u64;
    uint32_t 	u32[2];
    uint16_t 	u16[4];   
    uint8_t 	u8[8];   
};

uint64_t rand64() {
    uint64 tmp;
    tmp.u32[0] =  rand();
    tmp.u32[1] =  rand();
    return tmp.u64;
}

static void usage()
{
	std::cout << "usage: <xclbin> <sw_kernel>\n\n";
}

int main(int argc, char** argv)
{

	unsigned int cores_count = 0;
	float LNH_CLOCKS_PER_SEC;

	__foreach_core(group, core) cores_count++;

	//Assign xclbin
	if (argc < 3) {
		usage();
		throw std::runtime_error("FAILED_TEST\nNo xclbin specified");
	}

	//Open device #0
	leonhardx64 lnh_inst = leonhardx64(0,argv[1]);
	__foreach_core(group, core)
	{
		lnh_inst.load_sw_kernel(argv[2], group, core);
	}


	// /*
	//  *
	//  * Запись множества из BURST key-value и его последовательное чтение через Global Memory Buffer 
	//  *
	//  */


	//Выделение памяти под буферы gpc2host и host2gpc для каждого ядра и группы
	uint64_t *host2gpc_buffer[LNH_GROUPS_COUNT][LNH_MAX_CORES_IN_GROUP];
	__foreach_core(group, core)
	{
		host2gpc_buffer[group][core] = (uint64_t*) malloc(2*BURST*sizeof(uint64_t));
	}
	uint64_t *gpc2host_buffer[LNH_GROUPS_COUNT][LNH_MAX_CORES_IN_GROUP];
	__foreach_core(group, core)
	{
		gpc2host_buffer[group][core] = (uint64_t*) malloc(2*BURST*sizeof(uint64_t));
	}

	//Создание массива ключей и значений для записи в lnh64
	__foreach_core(group, core)
	{
		for (int i=0;i<BURST;i++) {
			//Первый элемент массива uint64_t - key
            uint64_t key = rand64() % 1048577;
			host2gpc_buffer[group][core][2*i] = key * key;
			//Второй uint64_t - value
			host2gpc_buffer[group][core][2*i+1] = key;
            //printf("key %llu, value %llu\n", key, key * key);
		}
	}

	//Запуск обработчика insert_burst
	__foreach_core(group, core) {
		lnh_inst.gpc[group][core]->start_async(__event__(insert_burst));
	}

	//DMA запись массива host2gpc_buffer в глобальную память
	__foreach_core(group, core) {
		lnh_inst.gpc[group][core]->buf_write(BURST*2*sizeof(uint64_t),(char*)host2gpc_buffer[group][core]);
	}

	//Ожидание завершения DMA
	__foreach_core(group, core) {
		lnh_inst.gpc[group][core]->buf_write_join();
	}

	//Передать количество key-value
	__foreach_core(group, core) {
		lnh_inst.gpc[group][core]->mq_send(BURST);
	}
	
	uint64_t* user_key = (uint64_t*)malloc(sizeof(uint64_t));
    user_key[0] = 158404;
    
    //DMA ?????? ??????? host2gpc_buffer ? ?????????? ??????
	__foreach_core(group, core) {
		lnh_inst.gpc[group][core]->buf_write(sizeof(uint64_t),(char*)user_key);
	}

	//???????? ?????????? DMA
	__foreach_core(group, core) {
		lnh_inst.gpc[group][core]->buf_write_join();
	}

	//Запуск обработчика для последовательного обхода множества ключей
	__foreach_core(group, core) {
		lnh_inst.gpc[group][core]->start_async(__event__(search_burst));
	}

	//Получить количество ключей
	unsigned int count[LNH_GROUPS_COUNT][LNH_MAX_CORES_IN_GROUP];

	__foreach_core(group, core) {
		count[group][core] = lnh_inst.gpc[group][core]->mq_receive();
        
	}


	//Прочитать количество ключей
	__foreach_core(group, core) {
		lnh_inst.gpc[group][core]->buf_read(count[group][core]*2*sizeof(uint64_t),(char*)gpc2host_buffer[group][core]);
	}

	//Ожидание завершения DMA
	__foreach_core(group, core) {
		lnh_inst.gpc[group][core]->buf_read_join();
	}
	
	

	bool error = false;
	//Проверка целостности данных
	__foreach_core(group, core) {
		uint64_t value = gpc2host_buffer[group][core][0];
        printf("flag %llu, val %llu\n",count[group][core], value);
	}


	__foreach_core(group, core) {
		free(host2gpc_buffer[group][core]);
		free(gpc2host_buffer[group][core]);
	}

	if (!error)
		printf("Тест пройден успешно!\n");
	else
		printf("Тест завершен с ошибкой!\n");


	return 0;
}
