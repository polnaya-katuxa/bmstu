//
// Created by Екатерина Карпова on 18.11.2023.
//

#include <stdlib.h>
#include <stdint.h>

#include "aes.h"
#include "galois.h"

static uint8_t s_table[256] = {
        0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76,
        0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0,
        0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15,
        0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75,
        0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84,
        0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf,
        0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8,
        0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2,
        0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73,
        0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb,
        0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79,
        0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08,
        0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a,
        0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e,
        0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf,
        0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16
};

void sub_bytes(uint8_t *state) {
    for (uint8_t i = 0; i < WORD_SIZE; i++) {
        for (uint8_t j = 0; j < N_B; j++) {
            state[N_B * i + j] = s_table[state[N_B * i + j]];
        }
    }
}

void shift_rows(uint8_t *state) {
    uint8_t s, tmp;
    for (uint8_t i = 1; i < WORD_SIZE; i++) {
        s = 0;
        while (s < i) {
            tmp = state[N_B * i + 0];
            for (uint8_t k = 1; k < N_B; k++) {
                state[N_B * i + k - 1] = state[N_B * i + k];
            }
            state[N_B * i + N_B - 1] = tmp;
            s++;
        }
    }
}

void mul_column(const uint8_t *matrix, const uint8_t *column, uint8_t *res) {
    res[0] = galois_mul(matrix[0], column[0]) ^ galois_mul(matrix[3], column[1]) ^ galois_mul(matrix[2], column[2]) ^
             galois_mul(matrix[1], column[3]);
    res[1] = galois_mul(matrix[1], column[0]) ^ galois_mul(matrix[0], column[1]) ^ galois_mul(matrix[3], column[2]) ^
             galois_mul(matrix[2], column[3]);
    res[2] = galois_mul(matrix[2], column[0]) ^ galois_mul(matrix[1], column[1]) ^ galois_mul(matrix[0], column[2]) ^
             galois_mul(matrix[3], column[3]);
    res[3] = galois_mul(matrix[3], column[0]) ^ galois_mul(matrix[2], column[1]) ^ galois_mul(matrix[1], column[2]) ^
             galois_mul(matrix[0], column[3]);
}

void mix_columns(uint8_t *state) {
    uint8_t matrix[] = {0x02, 0x01, 0x01, 0x03};

    uint8_t column[WORD_SIZE], res[WORD_SIZE];

    for (uint8_t j = 0; j < N_B; j++) {
        for (uint8_t i = 0; i < WORD_SIZE; i++) {
            column[i] = state[N_B * i + j];
        }

        mul_column(matrix, column, res);

        for (uint8_t i = 0; i < WORD_SIZE; i++) {
            state[N_B * i + j] = res[i];
        }
    }
}

void add_round_key(uint8_t *state, const uint8_t *expanded_key, uint8_t round_num) {
    for (uint8_t c = 0; c < N_B; c++) {
        state[N_B * 0 + c] = state[N_B * 0 + c] ^ expanded_key[WORD_SIZE * N_B * round_num + WORD_SIZE * c + 0];
        state[N_B * 1 + c] = state[N_B * 1 + c] ^ expanded_key[WORD_SIZE * N_B * round_num + WORD_SIZE * c + 1];
        state[N_B * 2 + c] = state[N_B * 2 + c] ^ expanded_key[WORD_SIZE * N_B * round_num + WORD_SIZE * c + 2];
        state[N_B * 3 + c] = state[N_B * 3 + c] ^ expanded_key[WORD_SIZE * N_B * round_num + WORD_SIZE * c + 3];
    }
}

void sub_word(uint8_t *key) {
    for (uint8_t i = 0; i < WORD_SIZE; i++) {
        key[i] = s_table[key[i]];
    }
}

void rot_word(uint8_t *key) {
    uint8_t tmp;
    tmp = key[0];
    for (uint8_t i = 0; i < 3; i++) {
        key[i] = key[i + 1];
    }
    key[3] = tmp;
}

void xor_word(const uint8_t *a, const uint8_t *b, uint8_t *result) {
    result[0] = a[0] ^ b[0];
    result[1] = a[1] ^ b[1];
    result[2] = a[2] ^ b[2];
    result[3] = a[3] ^ b[3];
}

uint8_t r_con_first(uint8_t i) {
    if (i == 1) {
        return 0x01;
    } else if (i > 1) {
        int n = 0x02;

        i--;
        while (i > 1) {
            n = galois_mul(n, 0x02);
            i--;
        }

        return n;
    }

    return 0;
}

void expand_key(const uint8_t *key, uint8_t *out) {
    uint8_t tmp[WORD_SIZE];

    for (uint8_t i = 0; i < N_K; i++) {
        out[WORD_SIZE * i + 0] = key[WORD_SIZE * i + 0];
        out[WORD_SIZE * i + 1] = key[WORD_SIZE * i + 1];
        out[WORD_SIZE * i + 2] = key[WORD_SIZE * i + 2];
        out[WORD_SIZE * i + 3] = key[WORD_SIZE * i + 3];
    }
    for (int i = N_K; i < N_B * (N_R + 1); i++) {
        tmp[0] = out[WORD_SIZE * (i - 1) + 0];
        tmp[1] = out[WORD_SIZE * (i - 1) + 1];
        tmp[2] = out[WORD_SIZE * (i - 1) + 2];
        tmp[3] = out[WORD_SIZE * (i - 1) + 3];
        if (i % N_K == 0) {
            rot_word(tmp);
            sub_word(tmp);

            uint8_t r_con[] = {r_con_first(i / N_K), 0x00, 0x00, 0x00};

            xor_word(tmp, r_con, tmp);
        } else if (N_K > 6 && i % N_K == WORD_SIZE) {
            sub_word(tmp);
        }

        out[WORD_SIZE * i + 0] = out[WORD_SIZE * (i - N_K) + 0] ^ tmp[0];
        out[WORD_SIZE * i + 1] = out[WORD_SIZE * (i - N_K) + 1] ^ tmp[1];
        out[WORD_SIZE * i + 2] = out[WORD_SIZE * (i - N_K) + 2] ^ tmp[2];
        out[WORD_SIZE * i + 3] = out[WORD_SIZE * (i - N_K) + 3] ^ tmp[3];
    }
}

void block_to_state(const uint8_t *in, uint8_t *state) {
    for (uint8_t i = 0; i < WORD_SIZE; i++) {
        for (uint8_t j = 0; j < N_B; j++) {
            state[N_B * i + j] = in[i + WORD_SIZE * j];
        }
    }
}

void block_from_state(const uint8_t *state, uint8_t *out) {
    for (uint8_t i = 0; i < WORD_SIZE; i++) {
        for (uint8_t j = 0; j < N_B; j++) {
            out[i + WORD_SIZE * j] = state[N_B * i + j];
        }
    }
}

void aes_encrypt(const uint8_t *in, uint8_t *out, uint8_t *key) {
    uint8_t state[WORD_SIZE * N_B];
    block_to_state(in, state);

    uint8_t expanded_key[WORD_SIZE * N_B * (N_R + 1)];
    expand_key(key, expanded_key);
    add_round_key(state, expanded_key, 0);

    for (uint8_t r = 1; r < N_R; r++) {
        sub_bytes(state);
        shift_rows(state);
        mix_columns(state);
        add_round_key(state, expanded_key, r);
    }

    sub_bytes(state);
    shift_rows(state);
    add_round_key(state, expanded_key, N_R);

    block_from_state(state, out);
}

