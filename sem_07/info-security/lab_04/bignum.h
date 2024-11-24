#ifndef LAB_04_BIGNUM_H
#define LAB_04_BIGNUM_H

#include <string.h>
#include <limits.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>

typedef struct
{
    int length;
    int capacity;
    uint32_t * data;   //data[0] is LSB, memory allocated at runtime
} bignum;

static uint32_t DATA0[1] = {0};
static uint32_t DATA1[1] = {1};
static uint32_t DATA2[1] = {2};
static uint32_t DATA3[1] = {3};
static uint32_t DATA4[1] = {4};
static uint32_t DATA5[1] = {5};
static uint32_t DATA6[1] = {6};
static uint32_t DATA7[1] = {7};
static uint32_t DATA8[1] = {8};
static uint32_t DATA9[1] = {9};
static uint32_t DATA10[1] = {10};

static bignum NUMS[11] = {{1, 1, DATA0},{1, 1, DATA1},{1, 1, DATA2},
                   {1, 1, DATA3},{1, 1, DATA4},{1, 1, DATA5},
                   {1, 1, DATA6},{1, 1, DATA7},{1, 1, DATA8},
                   {1, 1, DATA9},{1, 1, DATA10}};

bignum * bignum_alloc();
void bignum_free(bignum *b);

int bignum_read(FILE *file, bignum *num);
int bignum_write(FILE *file, bignum *num);
bignum *from_bin(const char *data, int bytes);

void bignum_set_zero(bignum *b);
void bignum_fromint(bignum* b, uint32_t num);
void bignum_fromstring(bignum* b, char* string);
char * bignum_tostring(bignum* b);
void bignum_copy(bignum* source, bignum* dest);
void print_bignum(char *s, bignum *b);

void bignum_random(int bytes, bignum *result);

//compare & testing
int bignum_iszero(bignum* b);
int bignum_isodd(bignum *b);
int bignum_isequal(bignum* b1, bignum* b2);
int bignum_isgreater(bignum* b1, bignum* b2);
int bignum_isless(bignum* b1, bignum* b2);
int bignum_isgeq(bignum* b1, bignum* b2);
int bignum_isleq(bignum* b1, bignum* b2);

//math operations
void bignum_iadd(bignum* source, bignum* add);
void bignum_iadd_2(bignum* source);
void bignum_add(bignum* result, bignum* b1, bignum* b2);
void bignum_isubtract(bignum* source, bignum* add);
void bignum_subtract(bignum* result, bignum* b1, bignum* b2);
void bignum_imultiply(bignum* source, bignum* add);
void bignum_multiply(bignum* result, bignum* b1, bignum* b2);
void bignum_idivide(bignum* source, bignum* div);
void bignum_idivide_2(bignum *source);
// void bignum_idivider(bignum* source, bignum* div, bignum* remainder);
void bignum_mod(bignum* source, bignum *div, bignum* remainder);
void bignum_imod(bignum* source, bignum* modulus);
void bignum_divide(bignum* quotient, bignum* remainder, bignum* b1, bignum* b2);

void bignum_pow_mod(bignum* result, bignum* base, bignum* expo, bignum* mod);

#endif //LAB_04_BIGNUM_H
