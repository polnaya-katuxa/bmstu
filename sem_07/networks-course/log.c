#include <stdlib.h>
#include <stdarg.h>
#include <stdio.h>
#include <time.h>
#include "log.h"

#define ANSI_COLOR_RED     "\x1b[31m"
#define ANSI_COLOR_GREEN   "\x1b[32m"
#define ANSI_COLOR_YELLOW  "\x1b[33m"
#define ANSI_COLOR_BLUE    "\x1b[34m"
#define ANSI_COLOR_MAGENTA "\x1b[35m"
#define ANSI_COLOR_CYAN    "\x1b[36m"
#define ANSI_COLOR_RESET   "\x1b[0m"

#define LOG_LEN 1024
#define LOG_TIME_LEN 128

int log_level = LOG_DEBUG;

void log_set_level(log_level_t level) {
    log_level = level;
}

void get_cur_time_str(char* time_buffer) {
    time_t rawtime;
    struct tm *local;

    time(&rawtime);
    local = localtime(&rawtime);
    strftime(time_buffer, LOG_LEN, ANSI_COLOR_BLUE "[%x %X]" ANSI_COLOR_RESET " ", local);
}

void log_debug(char* format, ...) {
    if (log_level < LOG_DEBUG) {
        return;
    }

    va_list arg;
    va_start(arg, format);

    char time_buffer[LOG_TIME_LEN];
    char log_str[LOG_LEN];

    get_cur_time_str(time_buffer);

    sprintf(log_str, "%s" ANSI_COLOR_CYAN "DEBUG" ANSI_COLOR_RESET " %s\n", time_buffer, format);
    vprintf(log_str, arg);
    va_end(arg);
}

void log_info(char* format, ...) {
    if (log_level < LOG_INFO) {
        return;
    }

    va_list arg;
    va_start(arg, format);

    char time_buffer[LOG_TIME_LEN];
    char log_str[LOG_LEN];

    get_cur_time_str(time_buffer);

    sprintf(log_str, "%s" ANSI_COLOR_GREEN "INFO" ANSI_COLOR_RESET " %s\n", time_buffer, format);
    vprintf(log_str, arg);
    va_end(arg);
}

void log_warning(char* format, ...) {
    if (log_level < LOG_WARNING) {
        return;
    }

    va_list arg;
    va_start(arg, format);

    char time_buffer[LOG_TIME_LEN];
    char log_str[LOG_LEN];

    get_cur_time_str(time_buffer);

    sprintf(log_str, "%s" ANSI_COLOR_YELLOW "WARNING" ANSI_COLOR_RESET " %s\n", time_buffer, format);
    vprintf(log_str, arg);
    va_end(arg);
}

void log_error(char* format, ...) {
    if (log_level < LOG_ERROR) {
        return;
    }

    va_list arg;
    va_start(arg, format);

    char time_buffer[LOG_TIME_LEN];
    char log_str[LOG_LEN];

    get_cur_time_str(time_buffer);

    sprintf(log_str, "%s" ANSI_COLOR_RED "ERROR" ANSI_COLOR_RESET " %s\n", time_buffer, format);
    vprintf(log_str, arg);
    va_end(arg);
}

void log_fatal(char* format, ...) {
    if (log_level < LOG_FATAL) {
        return;
    }

    va_list arg;
    va_start(arg, format);

    char time_buffer[LOG_TIME_LEN];
    char log_str[LOG_LEN];

    get_cur_time_str(time_buffer);

    sprintf(log_str, "%s" ANSI_COLOR_MAGENTA "FATAL" ANSI_COLOR_RESET " %s\n", time_buffer, format);
    vprintf(log_str, arg);
    va_end(arg);

    exit(1);
}
