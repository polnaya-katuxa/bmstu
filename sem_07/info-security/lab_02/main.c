#include <stdio.h>
#include <stdlib.h>

#include "des.h"
#include "cbc.h"
#include "io.h"

int main() {
    char src[100];
    char dst[100];
    block_t key, iv;

    printf("Input source file: ");
    if (scanf("%s", src) != 1) {
        printf("Cannot read src file name.\n");
        return EXIT_FAILURE;
    }

    printf("Input destination file: ");
    if (scanf("%s", dst) != 1) {
        printf("Cannot read dst file name.\n");
        return EXIT_FAILURE;
    }

    printf("Input key, iv: ");
    if (scanf("%llu %llu", &key, &iv) != 2) {
        printf("Cannot read keys.\n");
        return EXIT_FAILURE;
    }

    fflush(stdin);

    char mode;
    printf("Input mode, 'e' for encrypt, 'd' for decrypt: ");
    if (scanf("%c", &mode) != 1) {
        printf("Cannot read mode.\n");
        return EXIT_FAILURE;
    }

    unsigned char *buf;
    int size;
    if (io_read(src, &buf, &size) != EXIT_SUCCESS) {
        printf("Cannot read file.\n");
        return EXIT_FAILURE;
    }

    unsigned char *result = NULL;
    int new_len;
    if (mode == 'e') {
        result = cbc_encrypt(buf, size, key, iv, &new_len);
    } else if (mode == 'd') {
        result = cbc_decrypt(buf, size, key, iv, &new_len);
    } else {
        free(buf);
        printf("Unknown mode.\n");
        return EXIT_FAILURE;
    }

    if (result == NULL) {
        free(buf);
        printf("Cannot make result.\n");
        return EXIT_FAILURE;
    }

    if (io_write(dst, result, new_len) != EXIT_SUCCESS) {
        free(buf);
        free(result);
        printf("Cannot write.\n");
        return EXIT_FAILURE;
    }

    free(buf);
    free(result);

    return 0;
}
