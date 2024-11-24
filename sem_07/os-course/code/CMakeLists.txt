cmake_minimum_required(VERSION 3.26)
project(code C)

set(CMAKE_C_STANDARD 99)

add_executable(code my_driver.c
        client.c
        client.h)
