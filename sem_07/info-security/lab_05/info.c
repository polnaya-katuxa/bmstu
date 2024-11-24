#include <stdlib.h>
#include <stdio.h>

#include "info.h"

int get_size(char *filename) {
    FILE* f = fopen(filename, "r");
    if (f == NULL) {
        printf("Cannot open file %s.\n", filename);
        return -1;
    }

    fseek(f, 0, SEEK_END);
    int size = ftell(f);
    fclose(f);

    return size;
}

int get_compress_ratio(char* in_filename, char* compressed_filename, float *ratio) {
    int in_size = get_size(in_filename);
    if (in_size < 0) {
        printf("Cannot get file size for %s.\n", in_filename);
        return EXIT_FAILURE;
    }

    int compressed_size = get_size(compressed_filename);
    if (compressed_size < 0) {
        printf("Cannot get file size for %s.\n", compressed_filename);
        return EXIT_FAILURE;
    }

    if (in_size == 0) {
        printf("File %s is empty.\n", in_filename);
        return EXIT_FAILURE;
    }

    if (compressed_size == 0) {
        printf("File %s is empty.\n", compressed_filename);
        return EXIT_FAILURE;
    }

    *ratio = (float) in_size / (float) compressed_size * 100.0;

    return EXIT_SUCCESS;
}