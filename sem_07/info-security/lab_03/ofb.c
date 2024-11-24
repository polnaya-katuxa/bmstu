//
// Created by Екатерина Карпова on 18.11.2023.
//

#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

#include "ofb.h"
#include "aes.h"

#define DIVIDER 0x50

void ofb(uint8_t *data, int num_blocks, uint8_t *iv, uint8_t *key) {
    uint8_t iv_copy[BLOCK_SIZE + 1];

    memcpy(iv_copy, iv, BLOCK_SIZE + 1);

    for (int i = 0; i < num_blocks; i++) {
        aes_encrypt(iv_copy, iv_copy, key);

        for (int j = 0; j < BLOCK_SIZE; j++) {
            data[BLOCK_SIZE * i + j] = data[BLOCK_SIZE * i + j] ^ iv_copy[j];
        }
    }
}

uint8_t *round_buf(uint8_t *buf, int len, int *num_blocks) {
    int new_len = len - len % BLOCK_SIZE + BLOCK_SIZE;
    *num_blocks = new_len / BLOCK_SIZE;

    uint8_t *new_buf = realloc(buf, new_len * sizeof(char));
    if (new_buf == NULL) {
        printf("Cannot realloc buf.\n");
        return NULL;
    }

    new_buf[len] = DIVIDER;
    for (int i = len + 1; i < new_len; i++) {
        new_buf[i] = 0;
    }

    return new_buf;
}

int ofb_encrypt(uint8_t **buf, int len, uint8_t *key, uint8_t *iv, int *new_len) {
    int num_blocks;

    uint8_t *new_buf = round_buf(*buf, len, &num_blocks);
    if (new_buf == NULL) {
        printf("Cannot round buf.\n");
        return EXIT_FAILURE;
    }

    *new_len = num_blocks * BLOCK_SIZE;

    ofb(new_buf, num_blocks, iv, key);

    *buf = new_buf;

    return EXIT_SUCCESS;
}

int ofb_decrypt(uint8_t *buf, int len, uint8_t *key, uint8_t *iv, int *new_len) {
    int num_blocks = len / BLOCK_SIZE;

    ofb(buf, num_blocks, iv, key);

    *new_len = len;

    while (buf[*new_len] != DIVIDER) {
        (*new_len)--;
    }

    buf[*new_len] = '\0';

    return EXIT_SUCCESS;
}
