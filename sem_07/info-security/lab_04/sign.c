#include <stdlib.h>
#include <stdio.h>

#include "sign.h"
#include "rsa.h"
#include "io.h"
#include "md5.h"

int sign(char *filename, char *sign_filename, char *key_filename) {
    uint8_t hash[1024] = { 0 };
    if (md5(filename, hash) != EXIT_SUCCESS) {
        printf("Cannot compute MD5.\n");
        return EXIT_FAILURE;
    }

    printf("File checksum is ");
    md5_print(hash);
    printf(".\n");

    uint8_t sign_content[1024] = { 0 };
    int size;
    if ((size = rsa(hash, 16, key_filename, sign_content)) < 0) {
        printf("Cannot compute RSA.\n");
        return EXIT_FAILURE;
    }

    if (write_file(sign_filename, sign_content, size) != EXIT_SUCCESS) {
        printf("Cannot write file.\n");
        return EXIT_FAILURE;
    }

    printf("Sign file is %s.\n", sign_filename);

    return EXIT_SUCCESS;
}

int check(char *filename, char *sign_filename, char *key_filename) {
    uint8_t hash[16] = { 0 };
    if (md5(filename, hash) != EXIT_SUCCESS) {
        printf("Cannot compute MD5.\n");
        return EXIT_FAILURE;
    }
    printf("File checksum is ");
    md5_print(hash);
    printf(".\n");

    uint8_t* sign_content;
    int sign_size;
    if (read_file(sign_filename, &sign_content, &sign_size) != EXIT_SUCCESS) {
        printf("Read sign file.\n");
        return EXIT_FAILURE;
    }

    uint8_t hash_from_sign[1024] = { 0 };
    if (rsa(sign_content, sign_size, key_filename, hash_from_sign) < 0) {
        printf("Cannot compute RSA.\n");
        return EXIT_FAILURE;
    }

    printf("Checksum from sign ");
    md5_print(hash_from_sign);
    printf(".\n");

    for (int i = 0; i < 16; i++) {
        if (hash_from_sign[i] != hash[i]) {
            printf("File is corrupted.\n");
            return EXIT_FAILURE;
        }
    }

    printf("Check is successful.\n");

    return EXIT_SUCCESS;
}
