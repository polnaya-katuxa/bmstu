cmake_minimum_required(VERSION 3.26)
project(networks_course C)

set(CMAKE_C_STANDARD 23)

add_executable(networks_course main.c
        http.c
        http.h
        errors.h
        log.c
        log.h)
