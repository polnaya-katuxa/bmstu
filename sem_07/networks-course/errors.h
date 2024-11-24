#ifndef NETWORKS_COURSE_ERRORS_H
#define NETWORKS_COURSE_ERRORS_H

#define OK 0
#define ERR_UNSUPPORTED_HTTP_VERSION 1
#define ERR_INVALID_REQUEST 2
#define ERR_UNSUPPORTED_METHOD 3
#define ERR_CANT_WRITE_TO_SOCKET 4
#define ERR_CANT_SPRINTF 5
#define ERR_FILE_NOT_FOUND 6
#define ERR_FILESYSTEM 7
#define ERR_UNSUPPORTED_FILE 8
#define ERR_FORBIDDEN_PATH 9
#define ERR_EOF 10
#define ERR_READ 11
#define ERR_WRITE 12
#define ERR_ACCEPT 13

static char *error_messages[] = {
    [ERR_UNSUPPORTED_HTTP_VERSION] = "unsupported version of HTTP",
    [ERR_INVALID_REQUEST] = "invalid request",
    [ERR_UNSUPPORTED_METHOD] = "unsupported method, must be GET or HEAD",
    [ERR_CANT_WRITE_TO_SOCKET] = "cannot write to socket",
    [ERR_CANT_SPRINTF] = "cannot sprintf",
    [ERR_FILE_NOT_FOUND] = "file not found",
    [ERR_FILESYSTEM] = "unknown error with filesystem, cant get file",
    [ERR_UNSUPPORTED_FILE] = "unsupported file type",
    [ERR_FORBIDDEN_PATH] = "forbidden file path",
    [ERR_EOF] = "end of file",
    [ERR_READ] = "cannot read",
    [ERR_WRITE] = "cannot write",
    [ERR_ACCEPT] = "cannot accept",
};

#endif //NETWORKS_COURSE_ERRORS_H
