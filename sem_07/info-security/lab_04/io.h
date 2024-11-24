#ifndef LAB_04_IO_H
#define LAB_04_IO_H

int read_file(char* filename, uint8_t** buf, int* size);
int write_file(char* filename, uint8_t* buf, int size);

#endif //LAB_04_IO_H
