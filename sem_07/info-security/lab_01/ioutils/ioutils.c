//
// Created by Екатерина Карпова on 27.09.2023.
//

#include "ioutils.h"

int read_file(char* filename, unsigned char** str, size_t* s) {
    FILE* f;

    f = fopen(filename, "rb");
    if (f == NULL)
    {
        printf("cannot open the file\n");
        return OPEN_ERR;
    }

    fseek(f, 0L, SEEK_END);
    size_t size = ftell(f);
    rewind(f);

    *s = size;
    *str = malloc(size + 1);

    char c;
    for (size_t i = 0; i < size; i++) {
        c = fgetc(f);
        (*str)[i] = c;
    }

    (*str)[size] = '\0';

    fclose(f);

    return EXIT_SUCCESS;
}

int write_file(char* filename, unsigned char* str, size_t s) {
    FILE* f;

    f = fopen(filename, "wb");
    if (f == NULL)
    {
        printf("cannot open the file\n");
        return OPEN_ERR;
    }

    size_t n = fwrite(str, sizeof(unsigned char), s, f);
    if (n != s) {
        return WRITE_ERR;
    }

    fclose(f);

    return EXIT_SUCCESS;
}
