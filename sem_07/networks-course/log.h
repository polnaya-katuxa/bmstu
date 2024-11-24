#ifndef NETWORKS_COURSE_LOG_H
#define NETWORKS_COURSE_LOG_H

typedef enum {
    LOG_FATAL = 0,
    LOG_ERROR = 1,
    LOG_WARNING = 2,
    LOG_INFO = 3,
    LOG_DEBUG = 4,
} log_level_t;

void log_set_level(log_level_t level);

void log_debug(char* format, ...);
void log_info(char* format, ...);
void log_warning(char* format, ...);
void log_error(char* format, ...);
void log_fatal(char* format, ...);

#endif //NETWORKS_COURSE_LOG_H
