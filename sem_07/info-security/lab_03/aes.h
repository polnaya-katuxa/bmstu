//
// Created by Екатерина Карпова on 18.11.2023.
//

#ifndef LAB3_AES_H
#define LAB3_AES_H

#define KEY_SIZE 128
#define KEY_SIZE_BYTES (KEY_SIZE / 8)

// Количество слов в блоке.
#define N_B 4

// Количество слов в ключе.
#define N_K (KEY_SIZE / 32)

// Размер слова в байтах.
#define WORD_SIZE 4

// Количество раундов.
#if KEY_SIZE == 128
#define N_R 10
#elif KEY_SIZE == 192
#define N_R 12
#elif KEY_SIZE == 256
#define N_R 14
#endif

void aes_encrypt(const uint8_t *in, uint8_t *out, uint8_t *key);

#endif //LAB3_AES_H
