#include <stdio.h>
#include <stdlib.h>

#include "rsa.h"
#include "bignum.h"

typedef struct {
    bignum *exponent;
    bignum *modulus;
} rsa_key_t;

void gcd(bignum *aa, bignum *bb, bignum *result) {
    bignum *a = bignum_alloc();
    bignum *b = bignum_alloc();
    bignum_copy(aa, a);
    bignum_copy(bb, b);

    bignum *tmp = bignum_alloc();

    while (!bignum_iszero(b)) {
        bignum_copy(b, tmp);  // tmp = b
        bignum_imod(a, b); // a = a % b
        bignum_copy(a, b);    // b = a
        bignum_copy(tmp, a);  // a = tmp
    }

    bignum_copy(a, result);
}

// Modified extended Euclidean algorithm from Knuth.
void find_d(bignum *a, bignum *mod, bignum *result) {
    bignum *u1 = bignum_alloc();
    bignum *u3 = bignum_alloc();
    bignum *v1 = bignum_alloc();
    bignum *v3 = bignum_alloc();
    bignum *t1 = bignum_alloc();
    bignum *t3 = bignum_alloc();
    bignum *q = bignum_alloc();

    int iter = 1;
    bignum_fromint(u1, 1);
    bignum_copy(a, u3);
    bignum_copy(mod, v3);

    while (!bignum_iszero(v3)) {
        bignum_divide(q, t3, u3, v3);
        bignum_imultiply(q, v1);
        bignum_add(t1, u1, q);

        bignum_copy(v1, u1);
        bignum_copy(t1, v1);
        bignum_copy(v3, u3);
        bignum_copy(t3, v3);

        iter = -iter;
    }

    if (iter < 0)
        bignum_subtract(result, mod, u1);
    else
        bignum_copy(u1, result);

    bignum_free(t1);
    bignum_free(u1);
    bignum_free(v1);
    bignum_free(v3);
    bignum_free(t3);
    bignum_free(u3);
    bignum_free(q);
}

// Find result < phi such that gcd(result, phi) == 1.
void find_e(bignum *phi, bignum *result) {
    bignum *x = bignum_alloc();
    uint32_t e = 3;

    while (1) {
        bignum_fromint(result, e);
        gcd(result, phi, x);
        if (x->length == 1 && x->data[0] == 1) { // GCD == 1
            break;
        }

        e++;
    }

    bignum_free(x);
}

// Miller-Rabin test for primality.
int is_prime(bignum *n, int repeat) {
    bignum *n_1 = bignum_alloc();
    bignum *a = bignum_alloc();
    bignum *q = bignum_alloc();
    bignum *one = bignum_alloc();
    bignum *x = bignum_alloc();
    uint32_t k = 0;
    uint32_t i;
    int result = 1;

    one->length = 1;
    one->data[0] = 1;

    bignum_subtract(n_1, n, one);
    bignum_copy(n_1, q);
    //n-1 = 2^k x q
    while (!bignum_isodd(q)) {
        bignum_idivide_2(q);
        k++;
    }

    while (repeat--) {
        bignum_random(10 * n_1->length, a);
        bignum_imod(a, n_1);
        bignum_pow_mod(x, a, q, n);
        if (bignum_isequal(x, one))
            continue;   //a is not well chosen
        i = 0;
        while (!bignum_isequal(x, n_1)) {
            // printf("i=%u\n", i);
            if (i == (k - 1)) {
                return 0;
            } else {
                i++;
                bignum_imultiply(x, x);
                bignum_imod(x, n);
            }
        }
    }

    bignum_free(a);
    bignum_free(q);
    bignum_free(n_1);
    bignum_free(x);
    bignum_free(one);
    return result;
}

void random_prime(int bytes, bignum *result) {
    bignum_random(bytes, result);
    if (!bignum_isodd(result)) {
        result->data[0] += 1;
    }

//    printf("Generating big random prime number");
    while (1) {
//        printf(".");
        fflush(stdout);
        if (is_prime(result, 100))
            break;
        bignum_iadd_2(result);
    }
//    printf("\n");
}

int rsa_read_key(char *filename, rsa_key_t *key) {
    FILE *f;
    if (!(f = fopen(filename, "r"))) {
        return EXIT_FAILURE;
    }

    key->exponent = bignum_alloc();
    key->modulus = bignum_alloc();
    if (bignum_read(f, key->modulus) != EXIT_SUCCESS) {
        return EXIT_FAILURE;
    }

    if (bignum_read(f, key->exponent) != EXIT_SUCCESS) {
        return EXIT_FAILURE;
    }

    return EXIT_SUCCESS;
}

int rsa_write_key(char *filename, rsa_key_t *key) {
    FILE *f;
    if (!(f = fopen(filename, "w"))) {
        return EXIT_FAILURE;
    }

    if (bignum_write(f, key->modulus) != EXIT_SUCCESS) {
        return EXIT_FAILURE;
    }

    if (bignum_write(f, key->exponent) != EXIT_SUCCESS) {
        return EXIT_FAILURE;
    }

    return EXIT_SUCCESS;
}

int rsa_generate_keys(char *private_filename, char *public_filename, int bytes) {
    bignum *p = bignum_alloc();
    bignum *q = bignum_alloc();
    bignum *n = bignum_alloc();
    bignum *d = bignum_alloc();
    bignum *e = bignum_alloc();
    bignum *phi = bignum_alloc();

    random_prime(bytes, p);
    random_prime(bytes, q);

    bignum_multiply(n, p, q);
    bignum_isubtract(p, &NUMS[1]);
    bignum_isubtract(q, &NUMS[1]);
    bignum_multiply(phi, p, q);

    find_e(phi, e);
    find_d(e, phi, d);

    rsa_key_t private = {
        .exponent = e,
        .modulus = n,
    };

    rsa_key_t public = {
        .exponent = d,
        .modulus = n,
    };

    if (rsa_write_key(private_filename, &private) != EXIT_SUCCESS) {
        return EXIT_FAILURE;
    }

    if (rsa_write_key(public_filename, &public) != EXIT_SUCCESS) {
        return EXIT_FAILURE;
    }

    bignum_free(p);
    bignum_free(q);
    bignum_free(phi);

    return EXIT_SUCCESS;
}

void rsa_free_key(rsa_key_t *key) {
    bignum_free(key->exponent);
    bignum_free(key->modulus);
    free(key);
}

int rsa_with_key(const char *buf, int bytes, rsa_key_t *key, char *result) {
    bignum *res, *plain;
    int enc_bytes;

    plain = from_bin(buf, bytes);
    res = bignum_alloc();
    bignum_pow_mod(res, plain, key->exponent, key->modulus);

    enc_bytes = res->length * sizeof(uint32_t);

    for (int i = 0; i < enc_bytes; i++)
        result[i] = ((char *) res->data)[i];

    bignum_free(res);
    bignum_free(plain);

    return enc_bytes;
}

int rsa(const char *buf, int bytes, char *key_filename, char *result) {
    rsa_key_t key;
    if (rsa_read_key(key_filename, &key) != EXIT_SUCCESS) {
        printf("Can't read key.\n");
        return -1;
    }

    return rsa_with_key(buf, bytes, &key, result);
}
