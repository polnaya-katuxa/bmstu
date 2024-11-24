cmake_minimum_required(VERSION 3.26)
project(lab_02 C)

set(CMAKE_C_STANDARD 23)

add_executable(lab_02 main.c
        des.c
        des.h
        cbc.c
        cbc.h
        io.c
        io.h)
