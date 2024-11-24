//
// Created by Екатерина Карпова on 27.09.2023.
//

#include "enigma.h"

char* rome_num_10(int n) {
    char* rome[10] = {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"};

    return rome[n % 10];
}

void print_enigma(enigma_t* enigma) {
    printf("ENIGMA:\n");
    printf("counter: %d\n", enigma->counter);
    printf("size_alphabet: %zu\n", enigma->size_alphabet);
    printf("size_rotors: %zu\n", enigma->size_rotors);

    printf("rotors:\n");
    for (int i = 0; i < enigma->size_rotors; i++) {
        printf("%7s ", rome_num_10(i + 1));
    }
    printf("\n");

    for (size_t i = 0; i < enigma->size_alphabet; i++) {
        for (size_t j = 0; j < enigma->size_rotors; j++) {
            printf("%3zu:%-3d ", i, enigma->rotors[j][i]);
        }
        printf("\n");
    }

    printf("reflector:\n");
    for (size_t i = 0; i < enigma->size_alphabet; i++) {
        printf("%zu:%d ", i, enigma->reflector[i]);
    }
    printf("\n");

    printf("panel:\n");
    for (size_t i = 0; i < enigma->size_alphabet; i++) {
        printf("%zu:%d ", i, enigma->panel[i]);
    }
    printf("\n");
}

void free_enigma(enigma_t* enigma) {
    for (size_t i = 0; i < enigma->size_rotors; i++) {
        if (enigma->rotors[i] != NULL) {
            free(enigma->rotors[i]);
        }
    }

    if (enigma->rotors != NULL) {
        free(enigma->rotors);
    }

    if (enigma->reflector != NULL) {
        free(enigma->reflector);
    }

    if (enigma->panel != NULL) {
        free(enigma->panel);
    }

    enigma->size_alphabet = 0;
    enigma->size_rotors = 0;
    enigma->counter = 0;
}

int alloc_enigma(enigma_t* enigma)
{
    if (enigma->size_alphabet == 0 || enigma->size_rotors == 0) {
        printf("zero size\n");
        return ERR_ALLOC;
    }

    enigma->rotors = malloc(sizeof(int*) * enigma->size_rotors);
    if (enigma->rotors == NULL) {
        printf("rotors malloc\n");
        free_enigma(enigma);
        return ERR_ALLOC;
    }

    for (size_t i = 0; i < enigma->size_rotors; i++) {
        enigma->rotors[i] = malloc(sizeof(int) * enigma->size_alphabet);
        if (enigma->rotors[i] == NULL) {
            printf("rotor malloc\n");
            free_enigma(enigma);
            return ERR_ALLOC;
        }
    }

    enigma->reflector = malloc(sizeof(int) * enigma->size_alphabet);
    if (enigma->reflector == NULL) {
        printf("reflector malloc\n");
        free_enigma(enigma);
        return ERR_ALLOC;
    }

    enigma->panel = malloc(sizeof(int) * enigma->size_alphabet);
    if (enigma->panel == NULL) {
        printf("panel malloc\n");
        free_enigma(enigma);
        return ERR_ALLOC;
    }

    return EXIT_SUCCESS;
}

int create_enigma(enigma_t* enigma, char* cfg_file)
{
    char* str;
    size_t s;
    int rc = read_file(cfg_file, &str, &s);
    if (rc != EXIT_SUCCESS) {
        return rc;
    }

    rc = unmarshal_enigma(str, enigma);
    if (rc != EXIT_SUCCESS) {
        return rc;
    }

    return rc;
}

void shift_by_one(int* rotor, size_t size) {
    int last = rotor[size - 1];

    for (size_t i = size - 1; i >= 1; i--) {
        rotor[i] = rotor[i - 1];
    }

    rotor[0] = last;
}

size_t get_index(const int* rotor, size_t size, int val) {
    for (size_t i = 0; i < size; i++) {
        if (rotor[i] == val) {
            return i;
        }
    }

    return -1;
}

unsigned char* encode_enigma(enigma_t* enigma, unsigned char* msg, size_t size) {
    unsigned char* res = calloc(size, sizeof(unsigned char));

    enigma->counter = 1;
    for (size_t i = 0; i < size; i++, enigma->counter++) {
        int code = msg[i];

        code = enigma->panel[code];

        for (size_t j = 0; j < enigma->size_rotors; j++) {
            code = enigma->rotors[j][code];
        }

        code = enigma->reflector[code];

        for (int j = enigma->size_rotors - 1; j >= 0; j--) {
            code = get_index(enigma->rotors[j], enigma->size_alphabet, code);
        }

        code = enigma->panel[code];

        res[i] = code;

        shift_by_one(enigma->rotors[0], enigma->size_alphabet);
        for (size_t j = 1; j < enigma->size_rotors; j++) {
            if (enigma->counter % (enigma->size_alphabet * j) == 0) {
                shift_by_one(enigma->rotors[j], enigma->size_alphabet);
            }
        }
    }

    return res;
}
