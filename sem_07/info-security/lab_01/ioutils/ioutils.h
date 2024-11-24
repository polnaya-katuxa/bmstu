//
// Created by Екатерина Карпова on 27.09.2023.
//

#ifndef LAB_01_IOUTILS_H
#define LAB_01_IOUTILS_H

#define OPEN_ERR 1
#define WRITE_ERR 3

#include "stdio.h"
#include "stdlib.h"
#include "string.h"
#include <sys/types.h>
#include <sys/stat.h>


int read_file(char* filename, unsigned char** str, size_t* s);
int write_file(char* filename, unsigned char* str, size_t s);

#endif //LAB_01_IOUTILS_H
