#include <stdlib.h>
#include <limits.h>
#include <stdio.h>

#include "cbc.h"

#define CHARS_IN_BLOCK (64 / CHAR_BIT)
#define DIVIDER 0x50

block_t* cbc_encrypt_blocks(block_t* p, int len, block_t key, block_t iv) {
    block_t* c = calloc(len, sizeof(block_t));
    if (c == NULL) {
        return NULL;
    }

    c[0] = des_encrypt(des_default, p[0] ^ iv, key);
    for (int i = 1; i < len; i++) {
        c[i] = des_encrypt(des_default, c[i - 1] ^ p[i], key);
    }

    return c;
}

block_t* cbc_decrypt_blocks(block_t* c, int len, block_t key, block_t iv) {
    block_t* p = calloc(len, sizeof(block_t));
    if (p == NULL) {
        return NULL;
    }

    p[0] = des_decrypt(des_default, c[0], key) ^ iv;
    for (int i = 1; i < len; i++) {
        p[i] = des_decrypt(des_default, c[i], key) ^ c[i-1];
    }

    return p;
}

// Массив символов в массив блоков.
void to_blocks(block_t *blocks, unsigned char* buf, int len) {
    int cur = 0;
    for (int i = 0; i < len; i += CHARS_IN_BLOCK) {
        block_t block = 0;
        for (int j = i; j < i + CHARS_IN_BLOCK; j++) {
            block <<= CHAR_BIT;
            block |= buf[j];
        }

        blocks[cur] = block;
        cur++;
    }
}

// Массив блоков в массив символов.
void to_chars(block_t *blocks, unsigned char* buf, int len) {
    int cur = 0;
    for (int i = 0; i < len; i += CHARS_IN_BLOCK) {
        block_t block = blocks[cur];
        for (int j = i + CHARS_IN_BLOCK - 1; j >= i; j--) {
            buf[j] = (char) block;
            block >>= CHAR_BIT;
        }

        cur++;
    }
}

block_t* to_blocks_round(unsigned char* buf, int len, int *num_blocks) {
    int new_len = len - len % CHARS_IN_BLOCK + CHARS_IN_BLOCK;

    unsigned char* new_buf = realloc(buf, new_len * sizeof(char));
    if (new_buf == NULL) {
        printf("Cannot realloc buf.\n");
        return NULL;
    }

    new_buf[len] = DIVIDER;
    for (int i = len + 1; i < new_len; i++) {
        new_buf[i] = 0;
    }

    // В последнем блоке храним то, на сколько мы увеличили размер строки, чтобы она стала кратна
    // размеру блока.
    *num_blocks = new_len / CHARS_IN_BLOCK;
    block_t* blocks = calloc(*num_blocks, sizeof(block_t));
    if (blocks == NULL) {
        printf("Cannot calloc buf.\n");
        return NULL;
    }

    to_blocks(blocks, new_buf, new_len);

    return blocks;
}

unsigned char* cbc_encrypt(unsigned char* buf, int len, block_t key, block_t iv, int *new_len) {
    int num_blocks = 0;
    block_t* blocks = to_blocks_round(buf, len, &num_blocks);
    if (blocks == NULL) {
        return NULL;
    }

    block_t* encrypted = cbc_encrypt_blocks(blocks, num_blocks, key, iv);
    if (encrypted == NULL) {
        free(blocks);
        return NULL;
    }

    *new_len = num_blocks * CHARS_IN_BLOCK;
    unsigned char* result = calloc(*new_len, sizeof(unsigned char));
    if (result == NULL) {
        free(blocks);
        free(encrypted);
        return NULL;
    }

    to_chars(encrypted, result, num_blocks * CHARS_IN_BLOCK);
    return result;
}

unsigned char* cbc_decrypt(unsigned char* buf, int len, block_t key, block_t iv, int *new_len) {
    int num_blocks = len / CHARS_IN_BLOCK;
    block_t* blocks = calloc(num_blocks, sizeof(block_t));
    if (blocks == NULL) {
        return NULL;
    }

    to_blocks(blocks, buf, len);

    block_t* decrypted = cbc_decrypt_blocks(blocks, num_blocks, key, iv);
    if (decrypted == NULL) {
        free(blocks);
        return NULL;
    }

    *new_len = (num_blocks) * CHARS_IN_BLOCK;
    unsigned char* result = calloc(*new_len, sizeof(unsigned char));
    if (result == NULL) {
        free(blocks);
        free(decrypted);
        return NULL;
    }

    to_chars(decrypted, result, *new_len);

    while (result[*new_len] != DIVIDER) {
        (*new_len)--;
    }

    result[*new_len] = '\0';

    return result;
}

