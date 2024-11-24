#ifndef LAB_05_DICT_H
#define LAB_05_DICT_H

#define MAX_LZW_ELEMENTS 4096
#define BYTE_VALUES 256

typedef struct {
    uint8_t* data;   //buffer
    int      len;       //size of data
} bytes_t;

typedef struct {
    bytes_t data[MAX_LZW_ELEMENTS];
    int     len;               //number of data strings generated so far
} dict_t;

void init_dict(dict_t *d);
void free_dict(dict_t *d);
int is_in_dict(dict_t *d, bytes_t s);
int get_index(dict_t *d, bytes_t s);
bytes_t get_from_dict(dict_t* d, int code);
void add_to_dict(dict_t *d, bytes_t s);

void append_byte(bytes_t *b, uint8_t byte);
void append_bytes(bytes_t *b, bytes_t bytes);
void copy_bytes(bytes_t *dst, bytes_t src);
void copy_data(bytes_t *dst, uint8_t src);
void free_bytes(bytes_t *b);

#endif //LAB_05_DICT_H
