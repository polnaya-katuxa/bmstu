#include <stdio.h>
#include <stdlib.h>

#include "io.h"
#include "aes.h"
#include "ofb.h"

int main() {
    char src[100] = {0};
    char dst[100] = {0};
    uint8_t key[KEY_SIZE_BYTES + 1] = {0};
    uint8_t iv[16 + 1] = {0};

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

    printf("Input key: ");
    if (scanf("%s", &key) != 1) {
        printf("Cannot read key.\n");
        return EXIT_FAILURE;
    }

    printf("Input iv: ");
    if (scanf("%s", &iv) != 1) {
        printf("Cannot read iv.\n");
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

    int new_len;
    int rc;
    if (mode == 'e') {
        rc = ofb_encrypt(&buf, size, key, iv, &new_len);
    } else if (mode == 'd') {
        rc = ofb_decrypt(buf, size, key, iv, &new_len);
    } else {
        free(buf);
        printf("Unknown mode.\n");
        return EXIT_FAILURE;
    }

    if (rc == EXIT_FAILURE) {
        free(buf);
        printf("Cannot make result.\n");
        return EXIT_FAILURE;
    }

    if (io_write(dst, buf, new_len) != EXIT_SUCCESS) {
        free(buf);
        printf("Cannot write.\n");
        return EXIT_FAILURE;
    }

    free(buf);

    return 0;
}
