//
// Created by Екатерина Карпова on 18.11.2023.
//

#ifndef LAB3_IO_H
#define LAB3_IO_H

int io_read(char *filename, uint8_t **buf, int *size);

int io_write(char *filename, uint8_t *buf, int size);

#endif //LAB3_IO_H
