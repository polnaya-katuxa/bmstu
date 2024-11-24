#include <stdlib.h>
#include <memory.h>

#include "dict.h"

void init_dict(dict_t *d) {
    d->len = 0;
    for (int i = 0; i < BYTE_VALUES; i++) {
        uint8_t *tmp = calloc(1, 1);
        *tmp = i;
        d->data[i].data = tmp;
        d->data[i].len = 1;
        d->len++;
    }

    for (int i = BYTE_VALUES; i < MAX_LZW_ELEMENTS; i++) {
        d->data[i].data = NULL;
        d->data[i].len = 0;
    }
}

void free_dict(dict_t *d) {
    for (int i = 0; i < MAX_LZW_ELEMENTS; i++) {
        if (d->data[i].len) {
            free(d->data[i].data);
        }
    }
}

int is_in_dict(dict_t *d, bytes_t s) {
    for (int i = BYTE_VALUES; (d->data[i].data != NULL) && (i < MAX_LZW_ELEMENTS); i++) {
        if ((d->data[i].len == s.len) && memcmp(s.data, d->data[i].data, d->data[i].len) == 0) {
            return 1;
        }
        i++;
    }

    return 0;
}

int get_index(dict_t *d, bytes_t s) {
    for (int i = 0; (d->data[i].data != NULL) && (i < MAX_LZW_ELEMENTS); i++) {
        if ((d->data[i].len == s.len) && memcmp(s.data, d->data[i].data, d->data[i].len) == 0) {
            return i;
        }
    }

    return 0;
}

void add_to_dict(dict_t *d, bytes_t s) {
    if (d->len >= MAX_LZW_ELEMENTS)
        return;

    uint8_t *tmp = calloc(s.len, 1);
    memcpy(tmp, s.data, s.len);

    d->data[d->len].data = tmp;
    d->data[d->len].len = s.len;
    d->len++;
}

bytes_t get_from_dict(dict_t *d, int code) {
    bytes_t b = {0};
    copy_bytes(&b, d->data[code]);
    return b;
}

void copy_bytes(bytes_t *dst, bytes_t src) {
    dst->data = realloc(dst->data, src.len);
    memcpy(dst->data, src.data, src.len);
    dst->len = src.len;
}

void copy_data(bytes_t *dst, uint8_t src) {
    dst->data = realloc(dst->data, 1);
    memcpy(dst->data, &src, 1);
    dst->len = 1;
}

void append_byte(bytes_t *b, uint8_t byte) {
    b->data = realloc(b->data, b->len + 1);
    memcpy((b->data + b->len), &byte, 1);
    b->len++;
}

void append_bytes(bytes_t *b, bytes_t bytes) {
    b->data = realloc(b->data, b->len + bytes.len);
    memcpy((b->data + b->len), bytes.data, bytes.len);
    b->len = b->len + bytes.len;
}

void free_bytes(bytes_t *b) {
    if (b->len != 0) {
        free(b->data);
    }
}

