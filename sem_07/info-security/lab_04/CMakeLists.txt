cmake_minimum_required(VERSION 3.26)
project(lab_04 C)

set(CMAKE_C_STANDARD 23)

add_executable(lab_04 main.c
        bignum.c
        bignum.h
        md5.c
        md5.h
        io.c
        io.h
        io.h
        rsa.c
        rsa.h
        sign.c
        sign.h
)
