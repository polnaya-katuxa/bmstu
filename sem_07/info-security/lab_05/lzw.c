#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "lzw.h"
#include "dict.h"

uint8_t get_lsb_byte(short code) {
    short tmp = code & 0x00FF;
    return (uint8_t) tmp;
}

uint8_t get_msb_byte(short code) {
    short tmp = code & 0x0FF0;
    tmp = tmp >> 4;
    return (uint8_t) tmp;
}

uint8_t get_msb_nibble(short code) {
    short tmp = code & 0x0F00;
    tmp = tmp >> 8;
    return (uint8_t) tmp;
}

uint8_t get_lsb_nibble(short code) {
    short tmp = code & 0x000F;
    return (uint8_t) tmp;
}

uint8_t create_byte_from_nibble(uint8_t high_nibble, uint8_t low_nibble) {
    short tmp = high_nibble;
    tmp = tmp << 4;
    tmp = tmp | low_nibble;
    return (uint8_t) tmp;
}

void write_byte_data(FILE *f, short code, int finish) {
    static int remaining_bits = 0;
    // If some data is left out from previous op, then it would always be lsb nibble.
    static uint8_t remaining_data = 0;

    // We write MSB first and then LSB.
    uint8_t c1 = get_msb_byte(code);
    uint8_t c2 = get_lsb_nibble(code);

    if (finish == 1) {
        // Write remaining data if any.
        if (remaining_bits) {
            uint8_t tmp = create_byte_from_nibble(get_lsb_nibble(remaining_data), get_msb_nibble(code));
            fwrite(&tmp, 1, 1, f);
        }
        return;
    }
    if (remaining_bits == 0) {
        // Write lsb as it is.
        fwrite(&c1, 1, 1, f);
        remaining_data = c2;
        remaining_bits = 4;
    } else {
        uint8_t tmp = create_byte_from_nibble(get_lsb_nibble(remaining_data), get_msb_nibble(code));
        fwrite(&tmp, 1, 1, f);
        tmp = get_lsb_byte(code);
        fwrite(&tmp, 1, 1, f);
        remaining_data = 0;
        remaining_bits = 0;
    }
}

void encode_lzw_data(FILE *f_in, FILE *f_out, dict_t *d) {
    uint8_t C = '\0';
    bytes_t P = {0};

    // Read data from file and create table dynamically.
    if (fread(&C, 1, 1, f_in)) {
        append_byte(&P, C);
        while (!feof(f_in)) {
            if (fread(&C, 1, 1, f_in)) {
                bytes_t temp = {0};
                append_bytes(&temp, P);
                append_byte(&temp, C);
                if (is_in_dict(d, temp)) {
                    copy_bytes(&P, temp);
                } else {
                    // Output the data for P.
                    short code = get_index(d, P);
                    write_byte_data(f_out, code, 0);
                    add_to_dict(d, temp);
                    copy_data(&P, C);
                }

                free_bytes(&temp);
            }
        }
        int code = get_index(d, P);
        write_byte_data(f_out, code, 0);
        // Write remaining data if any.
        write_byte_data(f_out, 0, 1);
    }
    free_bytes(&P);
}

// This function takes the input file and compresses it using the LZW algorithm and saves the output file.
int lzw_compress(char *infile, char *outfile) {
    FILE *f_in = fopen(infile, "rb");
    if (f_in == NULL) {
        return EXIT_FAILURE;
    }

    FILE *f_out = fopen(outfile, "wb");
    if (f_out == NULL) {
        fclose(f_in);
        return EXIT_FAILURE;
    }

    dict_t d;
    init_dict(&d);

    encode_lzw_data(f_in, f_out, &d);

    free_dict(&d);
    fclose(f_in);
    fclose(f_out);

    return EXIT_SUCCESS;
}

short get_msb_nibble_as_short(uint8_t c) {
    short t = c;
    t = t & 0xF0;
    return t;
}

uint8_t get_lsb_nibble_as_char(uint8_t c) {
    short t = c;
    t = t & 0x0F;
    return (uint8_t) t;
}

short make_code(uint8_t high_byte, uint8_t low_byte) {
    short code = high_byte & 0xFF;
    short low_byte_code = low_byte & 0xFF;

    code = code << 4;
    code = code | low_byte_code;
    return code;
}

short get_lzw_code(FILE *f) {
    short code = -1;
    static short remaining_bits = 0;
    // If some data is left out from previous op, then it would always be msb nibble.
    static uint8_t remaining_data = 1;
    uint8_t c1, c2;

    if (feof(f)) {
        return code;
    }

    // Read first 1 or 2 bytes depending upon remaining_bits.
    if (remaining_bits) {
        if (fread(&c1, 1, 1, f)) {
            code = make_code(remaining_data << 4, c1);
        }

        remaining_data = 0;
        remaining_bits = 0;
    } else if (fread(&c1, 1, 1, f)) {
        if (feof(f)) {
            return code;
        }

        if (fread(&c2, 1, 1, f)) {
            short c3 = get_msb_nibble_as_short(c2);
            code = make_code(c1, c3 >> 4);
            remaining_bits = 4;
            remaining_data = get_lsb_nibble_as_char(c2);
        }
    }

    return code;
}

void decode_lzw_data(FILE *f_in, FILE *f_out, dict_t *d) {
    short new;

    short old = get_lzw_code(f_in);
    if (old < 0) {
        return;
    }

    bytes_t tmp = get_from_dict(d, old);
    bytes_t C = {0};
    bytes_t S = {0};

    fwrite(tmp.data, 1, tmp.len, f_out);
    copy_data(&C, tmp.data[0]);

    while ((new = get_lzw_code(f_in)) != -1) {
        if (d->data[new].len == 0) {
            S = get_from_dict(d, old);
            append_bytes(&S, C);
        } else {
            S = get_from_dict(d, new);
        }

        fwrite(S.data, 1, S.len, f_out);
        copy_data(&C, S.data[0]);
        tmp = get_from_dict(d, old);
        append_bytes(&tmp, C);
        add_to_dict(d, tmp);
        old = new;
    }
}

int lzw_decompress(char *infile, char *outfile) {
    FILE *f_in = fopen(infile, "rb");
    if (f_in == NULL) {
        return EXIT_FAILURE;
    }

    FILE *f_out = fopen(outfile, "wb");
    if (f_out == NULL) {
        fclose(f_in);
        return EXIT_FAILURE;
    }

    dict_t d;
    init_dict(&d);

    decode_lzw_data(f_in, f_out, &d);

    free_dict(&d);
    fclose(f_in);
    fclose(f_out);

    return EXIT_SUCCESS;
}
