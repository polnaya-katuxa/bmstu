//
// Created by Екатерина Карпова on 27.09.2023.
//

#ifndef LAB_01_ENIGMA_H
#define LAB_01_ENIGMA_H

#include "stdlib.h"
#include "ioutils.h"
#include "string.h"

#define ERR_JSON_PARSE 1
#define ERR_ALLOC 2

typedef struct {
    int counter;
    size_t size_alphabet;
    size_t size_rotors;
    int *panel;
    int *reflector;
    int **rotors;
} enigma_t;

int unmarshal_enigma(char* str, enigma_t* enigma);
int create_enigma(enigma_t* enigma, char* cfg_file);
int alloc_enigma(enigma_t* enigma);
void print_enigma(enigma_t* enigma);
void free_enigma(enigma_t* enigma);
unsigned char* encode_enigma(enigma_t* enigma, unsigned char* msg, size_t size);

#endif //LAB_01_ENIGMA_H
