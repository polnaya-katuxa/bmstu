#ifndef LAB_02_IO_H
#define LAB_02_IO_H

int io_read(char* filename, unsigned char** buf, int* size);
int io_write(char* filename, unsigned char* buf, int size);

#endif //LAB_02_IO_H
