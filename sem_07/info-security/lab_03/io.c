//
// Created by Екатерина Карпова on 18.11.2023.
//

#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>

#include "io.h"

int io_read(char *filename, uint8_t **buf, int *size) {
    FILE *f;

    f = fopen(filename, "rb");
    if (f == NULL) {
        printf("Cannot open file.\n");
        return EXIT_FAILURE;
    }

    fseek(f, 0L, SEEK_END);
    *size = (int) ftell(f);
    rewind(f);

    *buf = malloc(*size + 1);

    for (size_t i = 0; i < *size; i++) {
        uint8_t c = (uint8_t) fgetc(f);
        (*buf)[i] = c;
    }
    (*buf)[*size] = '\0';

    fclose(f);

    return EXIT_SUCCESS;
}

int io_write(char *filename, uint8_t *buf, int size) {
    FILE *f;

    f = fopen(filename, "wb");
    if (f == NULL) {
        printf("Cannot open file.\n");
        return EXIT_FAILURE;
    }

    size_t n = fwrite(buf, sizeof(uint8_t), size, f);
    if (n != size) {
        fclose(f);
        printf("Cannot write.\n");
        return EXIT_FAILURE;
    }

    fclose(f);

    return EXIT_SUCCESS;
}


