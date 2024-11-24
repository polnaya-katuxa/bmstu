//
// Created by Екатерина Карпова on 18.11.2023.
//

#ifndef LAB3_OFB_H
#define LAB3_OFB_H

#define BLOCK_SIZE 16

int ofb_encrypt(uint8_t **buf, int len, uint8_t *key, uint8_t *iv, int *new_len);

int ofb_decrypt(uint8_t *buf, int len, uint8_t *key, uint8_t *iv, int *new_len);

#endif //LAB3_OFB_H
