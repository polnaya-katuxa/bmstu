cmake_minimum_required(VERSION 3.26)
project(lab_03 C)

set(CMAKE_C_STANDARD 99)

add_executable(lab_03 main.c
        aes.c
        aes.h
        galois.h
        ofb.c
        ofb.h
        io.c
        io.h)
