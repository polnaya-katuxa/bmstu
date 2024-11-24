#include "enigma/enigma.h"
#include "ioutils/ioutils.h"

#define ERR_INVALID_ARGS 1

int main(int argc, char* argv[]) {
    if (argc != 4) {
        printf("expected more arguments: config file, input file, output file paths");
        return ERR_INVALID_ARGS;
    }

    enigma_t enigma;
    int rc = create_enigma(&enigma, argv[1]);
    if (rc != EXIT_SUCCESS) {
        printf("create enigma error");
        free_enigma(&enigma);
        return rc;
    }

    print_enigma(&enigma);

    unsigned char* msg;
    size_t s;
    rc = read_file(argv[2], &msg, &s);
    if (rc != EXIT_SUCCESS) {
        printf("read file error");
        free_enigma(&enigma);
        return rc;
    }

    unsigned char* encoded = encode_enigma(&enigma, msg, s);

    rc = write_file(argv[3], encoded, s);
    if (rc != EXIT_SUCCESS) {
        printf("write file error");
        free_enigma(&enigma);
        return rc;
    }

    free_enigma(&enigma);
    free(encoded);

    return EXIT_SUCCESS;
}
