#ifndef LAB_04_RSA_H
#define LAB_04_RSA_H

#include "bignum.h"

int rsa_generate_keys(char *private_filename, char *public_filename, int bytes);
int rsa(const char *buf, int bytes, char *key_filename, char *result);

#endif //LAB_04_RSA_H
