#ifndef LAB_02_CBC_H
#define LAB_02_CBC_H

#include "des.h"

unsigned char* cbc_encrypt(unsigned char* buf, int len, block_t key, block_t iv, int *new_len);
unsigned char* cbc_decrypt(unsigned char* buf, int len, block_t key, block_t iv, int *new_len);

block_t* cbc_encrypt_blocks(block_t* p, int len, block_t key, block_t iv);
block_t* cbc_decrypt_blocks(block_t* c, int len, block_t key, block_t iv);

#endif //LAB_02_CBC_H
