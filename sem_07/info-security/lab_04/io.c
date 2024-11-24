#include <stdlib.h>
#include <stdio.h>

#include "io.h"

int read_file(char* filename, uint8_t** buf, int* size) {
    FILE* f;

    f = fopen(filename, "rb");
    if (f == NULL)
    {
        printf("fopen()\n");
        return 1;
    }

    fseek(f, 0L, SEEK_END);
    *size = (int) ftell(f);
    rewind(f);

    *buf = malloc(*size + 1);

    for (int i = 0; i < *size; i++) {
        char c = (char) fgetc(f);
        (*buf)[i] = c;
    }
    (*buf)[*size] = 0;

    fclose(f);

    return 0;
}

int write_file(char* filename, uint8_t* buf, int size) {
    FILE* f;

    f = fopen(filename, "wb");
    if (f == NULL)
    {
        printf("fopen()\n");
        return 1;
    }

    size_t n = fwrite(buf, sizeof(uint8_t), size, f);
    if (n != size) {
        fclose(f);
        printf("fwrite()\n");
        return 1;
    }

    fclose(f);

    return 0;
}

