#include <stdio.h>
#include <stdlib.h>
#include <memory.h>

#include "md5.h"

typedef struct {
    uint64_t cur_len;
    uint8_t cur_input[64];

    uint32_t parts[4];

    uint8_t digest[16];
} hasher_t;

static uint32_t s[] = {7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
                       5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
                       4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
                       6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21};

static uint32_t k[] = {0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
                       0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
                       0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be,
                       0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
                       0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa,
                       0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
                       0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed,
                       0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
                       0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c,
                       0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
                       0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
                       0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
                       0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039,
                       0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
                       0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1,
                       0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391};

static uint8_t padding[] = {0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00};

uint32_t rotate_left(uint32_t x, uint32_t n) {
    return (x << n) | (x >> (32 - n));
}

void init(hasher_t *h) {
    h->cur_len = (uint64_t) 0;

    h->parts[0] = 0x67452301;
    h->parts[1] = 0xefcdab89;
    h->parts[2] = 0x98badcfe;
    h->parts[3] = 0x10325476;
}

void step(uint32_t *buf, const uint32_t *input) {
    uint32_t a = buf[0];
    uint32_t b = buf[1];
    uint32_t c = buf[2];
    uint32_t d = buf[3];

    for (unsigned int i = 0; i < 64; i++) {
        uint32_t f, g;

        if (i <= 15) {
            f = (b & c) | (~b & d);
            g = i % 16;
        }
        else if (i <= 31) {
            f = (b & d) | (c & ~d);
            g = ((i * 5) + 1) % 16;
        }
        else if (i <= 47) {
            f = (b ^ c ^ d);
            g = ((i * 3) + 5) % 16;
        }
        else {
            f = (c ^ (b | ~d));
            g = (i * 7) % 16;
        }

        f = f + a + k[i] + input[g];
        a = d;
        d = c;
        c = b;
        b = b + rotate_left(f, s[i]);
    }

    buf[0] += a;
    buf[1] += b;
    buf[2] += c;
    buf[3] += d;
}

void update(hasher_t *h, const uint8_t *buf, size_t len) {
    uint32_t input[16];
    uint64_t offset = h->cur_len % 64;
    h->cur_len += (uint64_t) len;

    for (unsigned int i = 0; i < len; i++) {
        h->cur_input[offset++] = buf[i];

        if (offset % 64 == 0) {
            for (unsigned int j = 0; j < 16; ++j) {
                input[j] = (uint32_t) (h->cur_input[(j * 4) + 3]) << 24 |
                           (uint32_t) (h->cur_input[(j * 4) + 2]) << 16 |
                           (uint32_t) (h->cur_input[(j * 4) + 1]) << 8 |
                           (uint32_t) (h->cur_input[(j * 4)]);
            }
            step(h->parts, input);
            offset = 0;
        }
    }
}

void finalize(hasher_t *h) {
    uint32_t input[16];
    unsigned int offset = h->cur_len % 64;
    unsigned int padding_length = offset < 56 ? 56 - offset : (56 + 64) - offset;

    update(h, padding, padding_length);
    h->cur_len -= (uint64_t) padding_length;

    for (unsigned int j = 0; j < 14; ++j) {
        input[j] = (uint32_t) (h->cur_input[(j * 4) + 3]) << 24 |
                   (uint32_t) (h->cur_input[(j * 4) + 2]) << 16 |
                   (uint32_t) (h->cur_input[(j * 4) + 1]) << 8 |
                   (uint32_t) (h->cur_input[(j * 4)]);
    }
    input[14] = (uint32_t) (h->cur_len * 8);
    input[15] = (uint32_t) ((h->cur_len * 8) >> 32);

    step(h->parts, input);

    for (unsigned int i = 0; i < 4; ++i) {
        h->digest[(i * 4) + 0] = (uint8_t) ((h->parts[i] & 0x000000FF));
        h->digest[(i * 4) + 1] = (uint8_t) ((h->parts[i] & 0x0000FF00) >> 8);
        h->digest[(i * 4) + 2] = (uint8_t) ((h->parts[i] & 0x00FF0000) >> 16);
        h->digest[(i * 4) + 3] = (uint8_t) ((h->parts[i] & 0xFF000000) >> 24);
    }
}

void md5_file(FILE *file, uint8_t *result) {
    uint8_t buf[64];
    size_t input_size;

    hasher_t h;
    init(&h);

    while ((input_size = fread(buf, 1, 64, file)) > 0) {
        update(&h, buf, input_size);
    }

    finalize(&h);

    memcpy(result, h.digest, 16);
}

int md5(char *filename, uint8_t *result) {
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        return EXIT_FAILURE;
    }

    md5_file(file, result);
    fclose(file);

    return EXIT_SUCCESS;
}

void md5_print(uint8_t *p) {
    for (unsigned int i = 0; i < 16; ++i) {
        printf("%02x", p[i]);
    }
}
