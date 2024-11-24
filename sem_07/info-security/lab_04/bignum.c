/*
 * Functions for handling math operations on big numbers
 * credits: https://github.com/pantaloons/RSA/blob/master/multiple.c
 * MIT License
*/

#include "bignum.h"

// Initial capacity for a bignum structure.
#define BIGNUM_DEFAULT_CAPACITY 20
#define RADIX 4294967296UL
#define HALFRADIX 2147483648UL

#define MAX(a,b) ((a) > (b) ? (a) : (b))
#define MIN(a,b) ((a) > (b) ? (b) : (a))

// Allocate memory and initialize a bignum, return a pointer
bignum * bignum_alloc()
{
    bignum* b = (bignum *) malloc(sizeof(bignum));
    b->length = 0;
    b->capacity = BIGNUM_DEFAULT_CAPACITY;
    b->data = (uint32_t *) calloc(BIGNUM_DEFAULT_CAPACITY, sizeof(uint32_t));
    return b;
}

// Free memory used by a bignum.
void bignum_free(bignum* b)
{
    free(b->data);
    free(b);
}

// На вход поступает bignum, под который уже выделена память
// на ошибки при выделении памяти проверки не делаю
int bignum_read(FILE *file, bignum *num) {
    if (fread(&(num->length), sizeof(num->length), 1, file) != 1)
        return EXIT_FAILURE;
    if (fread(&(num->capacity), sizeof(num->capacity), 1, file) != 1)
        return EXIT_FAILURE;
    num->data = (uint32_t *) realloc(num->data, num->capacity * sizeof(uint32_t));
    for (int i = 0; i < num->capacity; i++)
        if (fread(&(num->data[i]), sizeof(num->data[i]), 1, file) != 1)
            return EXIT_FAILURE;
    return EXIT_SUCCESS;
}

int bignum_write(FILE *file, bignum *num) {
    if (fwrite(&(num->length), sizeof(num->length), 1, file) != 1)
        return EXIT_FAILURE;
    if (fwrite(&(num->capacity), sizeof(num->capacity), 1, file) != 1)
        return EXIT_FAILURE;
    for (int i = 0; i < num->capacity; i++)
        if (fwrite(&(num->data[i]), sizeof(num->data[i]), 1, file) != 1)
            return EXIT_FAILURE;
    return EXIT_SUCCESS;
}

// Создать bigint из последовательности байтов
bignum *from_bin(const char *data, int bytes) {
    int len = (bytes + sizeof(uint32_t) - 1) / sizeof(uint32_t);
    bignum* b = (bignum *) malloc(sizeof(bignum));
    b->length = len;
    b->capacity = len;
    b->data = (uint32_t *) calloc(len, sizeof(uint32_t));
    for (int i = 0; i < bytes; i ++)
        ((char *) b->data)[i] = data[i];
    for (int i = bytes; i < len * sizeof(uint32_t); i ++)
        ((char *) b->data)[i] = 0;
    return b;
}

int bignum_iszero(bignum* b)
{
    return b->length == 0 || (b->length == 1 && b->data[0] == 0);
}

int bignum_isodd(bignum *b)
{
    if (b->length == 0)
        return 0;
    return (b->data[0] & 1);
}

int bignum_isequal(bignum* b1, bignum* b2)
{
    int i;
    if(bignum_iszero(b1) && bignum_iszero(b2)) return 1;
    else if(bignum_iszero(b1)) return 0;
    else if(bignum_iszero(b2)) return 0;
    else if(b1->length != b2->length) return 0;
    for(i = 0; i < b1->length; i ++)
        if(b1->data[i] != b2->data[i])
            return 0;
    return 1;
}

int bignum_isgreater(bignum* b1, bignum* b2)
{
    int i;
    if(bignum_iszero(b1) && bignum_iszero(b2)) return 0;
    else if(bignum_iszero(b1)) return 0;
    else if(bignum_iszero(b2)) return 1;
    else if(b1->length != b2->length)
        return b1->length > b2->length;
    for(i = b1->length - 1; i >= 0; i--)
    {
        if(b1->data[i] != b2->data[i])
            return b1->data[i] > b2->data[i];
    }
    return 0;
}

int bignum_isless(bignum* b1, bignum* b2)
{
    int i;
    if(bignum_iszero(b1) && bignum_iszero(b2)) return 0;
    else if(bignum_iszero(b1)) return 1;
    else if(bignum_iszero(b2)) return 0;
    else if(b1->length != b2->length)
        return b1->length < b2->length;
    for(i = b1->length - 1; i >= 0; i--)
    {
        if(b1->data[i] != b2->data[i])
            return b1->data[i] < b2->data[i];
    }
    return 0;
}

int bignum_isgeq(bignum* b1, bignum* b2)
{
    return !bignum_isless(b1, b2);
}

int bignum_isleq(bignum* b1, bignum* b2)
{
    return !bignum_isgreater(b1, b2);
}

//generate a random number of given bytes
void bignum_random(int bytes, bignum *result)
{
    int i;
    result->length = bytes / 4;
    if(result->capacity < bytes / 4)
    {
        result->capacity = bytes / 4;
        result->data = (uint32_t *) realloc(result->data, sizeof(uint32_t) * result->capacity);
    }
    for(i = 0; i < bytes/4; i++)
    {
        result->data[i] = rand();
    }
}

//add two bignums in place
void bignum_iadd(bignum* source, bignum* add)
{
    bignum* temp = bignum_alloc();
    bignum_add(temp, source, add);
    bignum_copy(temp, source);
    bignum_free(temp);
}

//source += 2
void bignum_iadd_2(bignum* source)
{
    bignum_iadd(source, &NUMS[2]);
}

//result = b1 + b2
void bignum_add(bignum* result, bignum* b1, bignum* b2) {
    uint32_t sum, carry = 0;
    int i;
    int n = MAX(b1->length, b2->length);
    if(n + 1 > result->capacity)
    {
        result->capacity = n + 1;
        result->data = (uint32_t *) realloc(result->data, result->capacity * sizeof(uint32_t));
    }
    for(i = 0; i < n; i++)
    {
        sum = carry;
        if(i < b1->length) sum += b1->data[i];
        if(i < b2->length) sum += b2->data[i];
        result->data[i] = sum; /* Already taken mod 2^32 by unsigned wrap around */

        if(i < b1->length)
        {
            if(sum < b1->data[i])
                carry = 1; /* Result must have wrapped 2^32 */
            else
                carry = 0;
        }
        else
        {
            if(sum < b2->data[i])
                carry = 1; /* Result must have wrapped 2^32 */
            else
                carry = 0;
        }
    }
    if(carry == 1)
    {
        result->length = n + 1;
        result->data[n] = 1;
    }
    else
        result->length = n;
}

//inplace substract
void bignum_isubtract(bignum* source, bignum* sub)
{
    bignum* temp = bignum_alloc();
    bignum_subtract(temp, source, sub);
    bignum_copy(temp, source);
    bignum_free(temp);
}

//result = b1-b2
void bignum_subtract(bignum* result, bignum* b1, bignum* b2)
{
    int length, i;
    uint32_t carry, diff, temp;

    length = carry = 0;
    if(b1->length > result->capacity)
    {
        result->capacity = b1->length;
        result->data = (uint32_t *) realloc(result->data, result->capacity * sizeof(uint32_t));
    }

    for(i = 0; i < b1->length; i++)
    {
        temp = carry;
        if(i < b2->length)
            temp = temp + b2->data[i]; /* Auto wrapped mod RADIX */
        diff = b1->data[i] - temp;
        if(temp > b1->data[i])
            carry = 1;
        else
            carry = 0;
        result->data[i] = diff;
        if(result->data[i] != 0)
            length = i + 1;
    }
    result->length = length;
}

void bignum_imultiply(bignum* source, bignum* mult)
{
    bignum* temp = bignum_alloc();
    bignum_multiply(temp, source, mult);
    bignum_copy(temp, source);
    bignum_free(temp);
}

void bignum_multiply(bignum* result, bignum* b1, bignum* b2)
{
    int i, j, k;
    uint32_t carry, temp;
    uint64_t prod; //intermediate product
    if(b1->length + b2->length > result->capacity)
    {
        result->capacity = b1->length + b2->length;
        result->data = (uint32_t *) realloc(result->data, result->capacity * sizeof(uint32_t));
    }
    for(i = 0; i < b1->length + b2->length; i++)
        result->data[i] = 0;

    for(i = 0; i < b1->length; i++)
    {
        for(j = 0; j < b2->length; j++)
        {
            prod = (b1->data[i] * (uint64_t)b2->data[j]) + (uint64_t)(result->data[i+j]);
            carry = (uint32_t)(prod / RADIX);

            /* Add carry to the next uint32_t over, but this may cause further overflow.. propogate */
            k = 1;
            while(carry > 0)
            {
                temp = result->data[i+j+k] + carry;
                if(temp < result->data[i+j+k]) carry = 1;
                else carry = 0;
                result->data[i+j+k] = temp; /* Already wrapped in unsigned arithmetic */
                k++;
            }

            prod = (result->data[i+j] + b1->data[i] * (uint64_t)b2->data[j]) % RADIX; /* Again, should not overflow... */
            result->data[i+j] = prod; /* Add */
        }
    }
    if(b1->length + b2->length > 0 && result->data[b1->length + b2->length - 1] == 0)
        result->length = b1->length + b2->length - 1;
    else
        result->length = b1->length + b2->length;
}

//source = source / div
void bignum_idivide(bignum *source, bignum *div)
{
    bignum *q = bignum_alloc();
    bignum *r = bignum_alloc();
    bignum_divide(q, r, source, div);
    bignum_copy(q, source);
    bignum_free(q);
    bignum_free(r);
}

//source /= 2
void bignum_idivide_2(bignum *source)
{
    bignum_idivide(source, &NUMS[2]);
}


//remainder = source % div
void bignum_mod(bignum* source, bignum *div, bignum* remainder)
{
    bignum *q = bignum_alloc();
    bignum_divide(q, remainder, source, div);
    bignum_free(q);
}

//source = source/div. remainder = source - source/div
void bignum_idivider(bignum* source, bignum* div, bignum* rem)
{
    bignum *q = bignum_alloc();
    bignum_divide(q, rem, source, div);
    bignum_copy(q, source);
    bignum_free(q);
}

//source = source % modulus
void bignum_imod(bignum* source, bignum* modulus) {
    bignum *q = bignum_alloc();
    bignum *r = bignum_alloc();
    bignum_divide(q, r, source, modulus);
    bignum_copy(r, source);
    bignum_free(q);
    bignum_free(r);
}

//quotient = b1 // b2
//remainder = b1 % b2
void bignum_divide(bignum* quotient, bignum* remainder, bignum* b1, bignum* b2) {
    bignum *b2copy = bignum_alloc(), *b1copy = bignum_alloc();
    bignum *temp = bignum_alloc(), *temp2 = bignum_alloc(), *temp3 = bignum_alloc();
    bignum* quottemp = bignum_alloc();
    uint32_t carry = 0;
    uint64_t factor = 1;
    uint64_t gquot, gtemp, grem;
    int n, m, i, j, length = 0;

    if(bignum_isless(b1, b2)) { /* Trivial case, b1/b2 = 0 iff b1 < b2. */
        quotient->length = 0;
        bignum_copy(b1, remainder);
    }
    else if(bignum_iszero(b1)) { /* 0/x = 0.. assuming b2 is nonzero */
        quotient->length = 0;
        bignum_fromint(remainder, 0);
    }
    else if(b2->length == 1) { /* Division by a single limb means we can do simple division */
        if(quotient->capacity < b1->length) {
            quotient->capacity = b1->length;
            quotient->data = (uint32_t *) realloc(quotient->data, quotient->capacity * sizeof(uint32_t));
        }
        for(i = b1->length - 1; i >= 0; i--) {
            gtemp = carry * RADIX + b1->data[i];
            gquot = gtemp / b2->data[0];
            quotient->data[i] = gquot;
            if(quotient->data[i] != 0 && length == 0) length = i + 1;
            carry = gtemp % b2->data[0];
        }
        bignum_fromint(remainder, carry);
        quotient->length = length;
    }
    else
    { /* do long division */
        n = b1->length + 1;
        m = b2->length;
        if(quotient->capacity < n - m) {
            quotient->capacity = n - m;
            quotient->data = (uint32_t *) realloc(quotient->data, (n - m) * sizeof(uint32_t));
        }
        bignum_copy(b1, b1copy);
        bignum_copy(b2, b2copy);
        /* Normalize.. multiply by the divisor by 2 until MSB >= HALFRADIX. This ensures fast
         * convergence when guessing the quotient below. We also multiply the dividend by the
         * same amount to ensure the result does not change. */
        while(b2copy->data[b2copy->length - 1] < HALFRADIX)
        {
            factor *= 2;
            bignum_imultiply(b2copy, &NUMS[2]);
        }
        if(factor > 1)
        {
            bignum_fromint(temp, factor);
            bignum_imultiply(b1copy, temp);
        }
        /* Ensure the dividend is longer than the original (pre-normalized) divisor. If it is not
         * we introduce a dummy zero uint32_t to artificially inflate it. */
        if(b1copy->length != n)
        {
            b1copy->length++;
            if(b1copy->length > b1copy->capacity) {
                b1copy->capacity = b1copy->length;
                b1copy->data = (uint32_t *) realloc(b1copy->data, b1copy->capacity * sizeof(uint32_t));
            }
            b1copy->data[n - 1] = 0;
        }

        /* Process quotient by long division */
        for(i = n - m - 1; i >= 0; i--)
        {
            gtemp = RADIX * b1copy->data[i + m] + b1copy->data[i + m - 1];
            gquot = gtemp / b2copy->data[m - 1];
            if(gquot >= RADIX) gquot = UINT_MAX;
            grem = gtemp % b2copy->data[m - 1];
            while(grem < RADIX && gquot * b2copy->data[m - 2] > RADIX * grem + b1copy->data[i + m - 2])
            { /* Should not overflow... ? */
                gquot--;
                grem += b2copy->data[m - 1];
            }
            quottemp->data[0] = gquot % RADIX;
            quottemp->data[1] = (gquot / RADIX);
            if(quottemp->data[1] != 0)
                quottemp->length = 2;
            else
                quottemp->length = 1;

            bignum_multiply(temp2, b2copy, quottemp);
            if(m + 1 > temp3->capacity)
            {
                temp3->capacity = m + 1;
                temp3->data = (uint32_t *) realloc(temp3->data, temp3->capacity * sizeof(uint32_t));
            }
            temp3->length = 0;
            for(j = 0; j <= m; j++)
            {
                temp3->data[j] = b1copy->data[i + j];
                if(temp3->data[j] != 0) temp3->length = j + 1;
            }
            if(bignum_isless(temp3, temp2)) {
                bignum_iadd(temp3, b2copy);
                gquot--;
            }

            bignum_isubtract(temp3, temp2);
            for(j = 0; j < temp3->length; j++)
                b1copy->data[i + j] = temp3->data[j];
            for(j = temp3->length; j <= m; j++)
                b1copy->data[i + j] = 0;
            quotient->data[i] = gquot;
            if(quotient->data[i] != 0)
                quotient->length = i;
        }

        if(quotient->data[b1->length - b2->length] == 0)
            quotient->length = b1->length - b2->length;
        else
            quotient->length = b1->length - b2->length + 1;

        /* Divide by factor now to find final remainder */
        carry = 0;
        for(i = b1copy->length - 1; i >= 0; i--)
        {
            gtemp = carry * RADIX + b1copy->data[i];
            b1copy->data[i] = gtemp/factor;
            if(b1copy->data[i] != 0 && length == 0)
                length = i + 1;
            carry = gtemp % factor;
        }
        b1copy->length = length;
        bignum_copy(b1copy, remainder);
    }
    bignum_free(temp);
    bignum_free(temp2);
    bignum_free(temp3);
    bignum_free(b1copy);
    bignum_free(b2copy);
    bignum_free(quottemp);
}

//dest = source
void bignum_copy(bignum* source, bignum* dest)
{
    dest->length = source->length;
    if(source->capacity > dest->capacity)
    {
        dest->capacity = source->capacity;
        dest->data = (uint32_t *) realloc(dest->data, dest->capacity * sizeof(uint32_t));
    }
    memcpy(dest->data, source->data, dest->length * sizeof(uint32_t));
}

//b = 0
void bignum_set_zero(bignum *b)
{
    int i;
    for(i = 0; i < b->length; i++)
        b->data[i] = 0;
    b->length = 0;
}

//load a bignum from an unsigned integer.
void bignum_fromint(bignum* b, uint32_t num)
{
    bignum_set_zero(b);
    b->length = 1;
    if(b->capacity < 1)
    {
        b->capacity = 1;
        b->data = (uint32_t *) realloc(b->data, sizeof(uint32_t));
    }
    b->data[0] = num;
}

//load bignum from a string ended with '\0'
void bignum_fromstring(bignum* b, char* string)
{
    int i, len = 0;
    len = strlen(string);
    bignum_set_zero(b);
    for(i = 0; i < len; i++)
    {
        if(i != 0)
            bignum_imultiply(b, &NUMS[10]);
        bignum_iadd(b, &NUMS[string[i] - '0']);
    }
}

//convert bignum to a string
char * bignum_tostring(bignum* b)
{
    int cap = 200;
    int len = 0;
    int i;
    char* buffer;
    bignum *copy;
    bignum *remainder;

    uint32_t tmp;

    buffer = (char *) malloc(cap * sizeof(char));
    if(bignum_iszero(b))
    {
        buffer[0] = '0';
        buffer[1] = '\0';
    }
    else
    {
        copy = bignum_alloc();
        remainder = bignum_alloc();

        bignum_copy(b, copy);
        while(!bignum_iszero(copy))
        {
            bignum_idivider(copy, &NUMS[10], remainder);
            buffer[len++] = remainder->data[0] + '0';
            if(len >= cap)
            {
                cap *= 2;
                buffer = (char *) realloc(buffer, cap * sizeof(char));
            }
        }
        //flip around
        for(i = 0; i < (len / 2); i++)
        {
            tmp = buffer[i];
            buffer[i] = buffer[len - i - 1];
            buffer[len - i - 1] = tmp;
        }
        buffer[len] = '\0';
        bignum_free(copy);
        bignum_free(remainder);
    }
    return buffer;
}

void print_bignum(char *s, bignum *b)
{
    char *k;
    k = bignum_tostring(b);
    printf("%s = %s\n", s, k);
    free(k);
}

//result = base^exp % mod
void bignum_pow_mod(bignum* result, bignum* base, bignum* expo, bignum* mod)
{
    bignum *a = bignum_alloc(), *b = bignum_alloc();
    bignum *tmp = bignum_alloc();

    bignum_copy(base, a);
    bignum_copy(expo, b);
    bignum_fromint(result, 1);

    while(!bignum_iszero(b))
    {
        if(b->data[0] & 1)
        {
            bignum_imultiply(result, a);
            bignum_imod(result, mod);
        }
        bignum_idivide_2(b);
        bignum_copy(a, tmp);
        bignum_imultiply(a, tmp);
        bignum_imod(a, mod);
    }
    bignum_free(a);
    bignum_free(b);
    bignum_free(tmp);
}
