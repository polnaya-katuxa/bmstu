#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "lzw.h"
#include "info.h"

int usage_cmd(int argc, char **argv) {
    printf("Usage of %s.\n", argv[0]);
    printf("%s compress <input file> — compressing file.\n", argv[0]);
    printf("%s decompress <input file> <output file> — decompressing file.\n", argv[0]);
    return EXIT_FAILURE;
}

int compress_cmd(int argc, char **argv) {
    if (argc < 3) {
        printf("Num of arguments must be 3.\n");
        return usage_cmd(argc, argv);
    }

    int rc;

    char *in_filename = argv[2];
    char compressed_filename[100];
    sprintf(compressed_filename, "%s.lzw", argv[2]);

    if ((rc = lzw_compress(in_filename, compressed_filename)) != EXIT_SUCCESS) {
        printf("Cannot compress via LZW.\n");
        return rc;
    }

    float ratio;
    if ((rc = get_compress_ratio(in_filename, compressed_filename, &ratio)) != EXIT_SUCCESS) {
        printf("Cannot compute compress ratio.\n");
        return rc;
    }

    printf("Compressed successfully.\n");
    printf("Compress ratio = %.4f%%.\n", ratio);

    return EXIT_SUCCESS;
}

int decompress_cmd(int argc, char **argv) {
    if (argc < 4) {
        printf("Num of arguments must be 4.\n");
        return usage_cmd(argc, argv);
    }

    int rc;

    if ((rc = lzw_decompress(argv[2], argv[3])) != EXIT_SUCCESS) {
        printf("Cannot decompress via LZW.\n");
        return rc;
    }

    printf("Decompressed successfully.\n");

    return EXIT_SUCCESS;
}

int main(int argc, char **argv) {
    if (argc < 2) {
        printf("Num of arguments must be at least 2.\n");
        return usage_cmd(argc, argv);
    }

    if (strcmp(argv[1], "compress") == 0) {
        return compress_cmd(argc, argv);
    }

    if (strcmp(argv[1], "decompress") == 0) {
        return decompress_cmd(argc, argv);
    }

    printf("Invalid subcommand.\n");
    return usage_cmd(argc, argv);
}
