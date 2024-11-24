#include <stdlib.h>
#include <stdio.h>

#include "des.h"

/*
 * mask делает маску из единиц.
 * mask(3) = 000...000111.
 */
block_t mask(int n) {
    return ((block_t) 1 << n) - 1;
}

block_t get_bit_count_from_right(block_t b, int n) {
    return (b >> n) & 1;
}

/*
 * get_bit_count_from_left получение бита числа, если считать номер бита с единицы, и отсчитывать справа.
 */
block_t get_bit_count_from_left(block_t b, int n, int size) {
    return (b >> (size - n)) & 1;
}

/*
 * set_bit установка бита числа, если считать номер бита с единицы, и отсчитывать справа.
 */
block_t set_bit(block_t b, block_t bit, int n, int size) {
    return b | (bit << (size - n));
}

/*
 * ip производит начальную перестановку блока в соответствии с матрицей.
 * 64 -> 64.
 */
block_t ip(des_t des, block_t block) {
    block_t result = 0;

    for (int i = 1; i <= SIZE; i++) {
        result = set_bit(result, get_bit_count_from_left(block, des.ip[i - 1], SIZE), i, SIZE);
    }

    return result;
}

/*
 * ip_reversed производит конечную перестановку блока в соответствии с инфертированной матрицей.
 * 64 -> 64.
 */
block_t ip_reversed(des_t des, block_t block) {
    block_t result = 0;

    for (int i = 1; i <= SIZE; i++) {
        result = set_bit(result, get_bit_count_from_left(block, des.ip_reversed[i - 1], SIZE), i, SIZE);
    }

    return result;
}

/*
 * g производит начальную перестановку ключа, и выкидывание лишних битов.
 * Если изначально ключ был 64 бита, то станет 56 бит.
 * 64 -> 56.
 */
block_t g(des_t des, block_t key) {
    block_t result = 0;

    for (int i = 1; i <= KEY_SIZE; i++) {
        result = set_bit(result, get_bit_count_from_left(key, des.g[i - 1], SIZE), i, KEY_SIZE);
    }

    return result;
}

/*
 * h производит конечную перестановку ключа.
 * 56 -> 48.
 */
block_t h(des_t des, block_t key) {
    block_t result = 0;

    for (int i = 1; i <= KEY_COMPRESSED_SIZE; i++) {
        result = set_bit(result, get_bit_count_from_left(key, des.h[i - 1], KEY_SIZE), i, KEY_COMPRESSED_SIZE);
    }

    return result;
}

/*
 * p производит перестановку перед выдачей результата функцией f().
 * 32 -> 32.
 */
block_t p(des_t des, block_t r) {
    block_t result = 0;

    for (int i = 1; i <= HALF_SIZE; i++) {
        result = set_bit(result, get_bit_count_from_left(r, des.p[i - 1], HALF_SIZE), i, HALF_SIZE);
    }

    return result;
}

/*
 * e производит расширение r на 48 бит и перемешивание.
 * 32 -> 48.
 */
block_t e(des_t des, block_t r) {
    block_t result = 0;

    for (int i = 1; i <= KEY_COMPRESSED_SIZE; i++) {
        result = set_bit(result, get_bit_count_from_left(r, des.e[i - 1], HALF_SIZE), i, KEY_COMPRESSED_SIZE);
    }

    return result;
}

/*
 * f функция для расчета, используемая при расчете нового значения r.
 * ФУНКЦИЯ ФЕЙСТЕЛЯ!
 */
block_t f(des_t des, block_t r, block_t k) {
    r = e(des, r);

    r = r ^ k;

    block_t b[8];
    for (int i = 8; i > 0; i--) {
        b[i - 1] = r & mask(6);
        r >>= 6;
    }

    for (int i = 8; i > 0; i--) {
        block_t row = get_bit_count_from_right(b[i - 1], 5) << 1 | get_bit_count_from_right(b[i - 1], 0);
        block_t col = (b[i - 1] >> 1) & mask(4);
        b[i - 1] = des.s[i - 1][row][col];
    }

    r = 0;
    for (int i = 1; i <= 8; i++) {
        r <<= 4;
        r |= b[i - 1];
    }

    return p(des, r);
}

/*
 * cycle_shift_left циклический сдвиг влево числа размера size бит на shift бит.
 * Пример:
 * Пусть будет блок размером 8 бит (block_t), но фактический размер хранимого числа должен быть 3 бита.
 * (Актуально, так как в коде есть место, где в 64-битной переменной лежит 28-битное значение).
 * 0000 0011, size = 3.
 * Мы хотим сдвинуть влево на 2 бита.
 * По идее, 011 << 1 = 110, 110 << 1 = 101 => 011 << 2 = 101.
 * То есть нужно отрезать те биты, которые выйдут за границу числа при сдвиге, потом сделать сдвиг, отрезав эти биты.
 * Затем справа дописать эти биты.
 *
 * То есть в данном случае:
 * 0000 0011, в результате сдвига на 2 (0000 0011 << 2 = 0000 1100) отрежется такая часть: 0000 1. Сохраним ее в
 * буфер overflow. overflow = 0000 0001.
 *
 * Из просто сдвинутого числа мы overflow можем получить просто сдвинув его вправо на размер 0000 1100 >> 3 = 0000 0001.
 * Но из изначального числа мы можем его посчитать, сдвинув вправо на (размер - сдвиг).
 * 0000 0011 >> (3 - 2) = 0000 0001.
 *
 * Затем сдвинуть число на сдвиг: 0000 0011 << 2 = 0000 1100.
 * Зачем отрезать биты, которые вышли за границу с помощью маски: 0000 1100 & 0000 0111 = 0000 0100.
 *
 * Затем к полученному справа добавляем то, что вышло за границу:
 * 0000 0100 & overflow = 0000 0100 & 0000 0001 = 0000 0101.
 */
block_t cycle_shift_left(block_t block, int size, int shift) {
    // Так как сдвиг циклический, то сдвигать больше, чем на размер не имеет смысла.
    shift %= size;

    block_t overflow = block >> (size - shift); // 0000 0011 >> (3 - 2) = 1.
    block <<= shift; // 0000 0011 << 2 = 0000 1100.

    block &= mask(size); // 0000 1100 & 111 = 0000 0100.
    block |= overflow; // 0000 0100 | 0000 0001 = 0000 0101.

    return block;
}

/*
 * fill_keys расчитывает все ключи Ki, нужные для шифрование/расшифрования.
 */
void fill_keys(des_t des, block_t k[17], block_t raw_key) {
    k[0] = g(des, raw_key);

    block_t c[17], d[17];
    c[0] = (k[0] >> HALF_KEY_SIZE) & mask(HALF_KEY_SIZE);
    d[0] = k[0] & mask(HALF_KEY_SIZE);

    for (int i = 1; i <= 16; i++) {
        // des.shifts[i-1] так как в нем индексация с нуля.
        c[i] = cycle_shift_left(c[i-1], HALF_KEY_SIZE, des.shifts[i-1]);
        d[i] = cycle_shift_left(d[i-1], HALF_KEY_SIZE, des.shifts[i-1]);

        k[i] = h(des, d[i] | (c[i] << HALF_KEY_SIZE));
    }
}

block_t des_encrypt(des_t des, block_t t, block_t key) {
    block_t t0 = ip(des, t);

    block_t r[17], l[17];
    r[0] = t0 & mask(HALF_SIZE);
    l[0] = (t0 >> HALF_SIZE) & mask(HALF_SIZE);

    block_t k[17];
    fill_keys(des, k, key);

    for (int i = 1; i <= 16; i++) {
        l[i] = r[i-1];
        r[i] = l[i-1] ^ f(des, r[i-1], k[i]);
    }

    block_t t16 = l[16] | (r[16] << HALF_SIZE);

    return ip_reversed(des, t16);
}

block_t des_decrypt(des_t des, block_t t, block_t key) {
    block_t t16 = ip(des, t);

    block_t r[17], l[17];
    l[16] = t16 & mask(HALF_SIZE);
    r[16] = (t16 >> HALF_SIZE) & mask(HALF_SIZE);

    block_t k[17];
    fill_keys(des, k, key);

    for (int i = 16; i > 0; i--) {
        r[i-1] = l[i];
        l[i-1] = r[i] ^ f(des, l[i], k[i]);
    }

    block_t t0 = r[0] | (l[0] << HALF_SIZE);

    return ip_reversed(des, t0);
}