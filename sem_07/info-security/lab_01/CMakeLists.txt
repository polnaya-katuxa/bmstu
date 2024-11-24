cmake_minimum_required(VERSION 3.26)
project(lab_01 C)

set(CMAKE_C_STANDARD 99)

add_executable(lab_01 main.c
        enigma/enigma.h
        enigma/enigma.c
        ioutils/ioutils.c
        ioutils/ioutils.h
        enigma/json.c
)

include_directories(ioutils)

find_package(cJSON)
if(cJSON_FOUND)
    include_directories(${CJSON_INCLUDE_DIRS})
    target_link_libraries(${PROJECT_NAME} ${CJSON_LIBRARIES})
    message(WARNING "yra sobralos ${PROJECT_NAME})))")
else()
    message(FATAL_ERROR "aaa ne sobralos(((")
endif()

set(CPACK_PROJECT_NAME ${PROJECT_NAME})
set(CPACK_PROJECT_VERSION ${PROJECT_VERSION})
include(CPack)
