#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "rsa.h"
#include "sign.h"

int main(int argc, char **argv) {
    if (argc < 2) {
        printf("Invalid number of arguments, must be at least 2.\n");
        return EXIT_FAILURE;
    }

    if (strcmp(argv[1], "generate") == 0) {
        if (argc != 3) {
            printf("Invalid number of arguments, must be 3.\n");
            return EXIT_FAILURE;
        }

        char *private_filename = argv[2];
        char public_filename[128] = { 0 };
        sprintf(public_filename, "%s.pub", private_filename);

        if (rsa_generate_keys(private_filename, public_filename, 32) != EXIT_SUCCESS) {
            printf("Cannot generate RSA keys.\n");
            return EXIT_FAILURE;
        }

        printf("Key pair generated successfully.\n");

        return EXIT_SUCCESS;
    }

    if (strcmp(argv[1], "sign") == 0) {
        if (argc != 4) {
            printf("Invalid number of arguments, must be 4.\n");
            return EXIT_FAILURE;
        }

        char *key_filename = argv[2];
        char *filename = argv[3];
        char sign_filename[100];
        sprintf(sign_filename, "%s.sign", filename);

        if (sign(filename, sign_filename, key_filename) != EXIT_SUCCESS) {
            printf("Cannot sign.\n");
            return EXIT_FAILURE;
        }

        return EXIT_SUCCESS;
    }

    if (strcmp(argv[1], "check") == 0) {
        if (argc != 4) {
            printf("Invalid number of arguments, must be 4.\n");
            return EXIT_FAILURE;
        }

        char *key_filename = argv[2];
        char *filename = argv[3];
        char sign_filename[100];
        sprintf(sign_filename, "%s.sign", filename);

        if (check(filename, sign_filename, key_filename) != EXIT_SUCCESS) {
            printf("Check is failed.\n");
            return EXIT_FAILURE;
        }

        return EXIT_SUCCESS;
    }

    printf("Unknown command.\n");
    return EXIT_FAILURE;
}
