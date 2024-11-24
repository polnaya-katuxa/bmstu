#ifndef LAB_04_MD5_H
#define LAB_04_MD5_H

#include <stdio.h>
#include <stdint.h>

int md5(char *filename, uint8_t *result);
void md5_print(uint8_t *p);

#endif //LAB_04_MD5_H
