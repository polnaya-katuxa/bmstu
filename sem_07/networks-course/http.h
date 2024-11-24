#ifndef NETWORKS_COURSE_HTTP_H
#define NETWORKS_COURSE_HTTP_H

typedef enum {
    METHOD_GET = 0,
    METHOD_HEAD = 1,
} http_method_t;

typedef enum {
    RESPONSE_CODE_OK = 0, // 200
    RESPONSE_CODE_FORBIDDEN = 1, // 403
    RESPONSE_CODE_NOT_FOUND = 2, // 404
    RESPONSE_CODE_METHOD_NOT_ALLOWED = 3, // 405
    RESPONSE_CODE_BAD_REQUEST = 4, // 400
    RESPONSE_CODE_INTERNAL_SERVER_ERROR = 5, // 500
} http_response_code_t;

typedef struct {
    http_method_t method;
    char uri[10000];
    char version[10];
    int fd;
} http_request_t;

static char *http_methods_map[] = {
    [METHOD_GET] = "GET",
    [METHOD_HEAD] = "HEAD",
};

int http_parse_request(char *raw, http_request_t *req);
int http_write_error(http_request_t *req, http_response_code_t code);
int http_handle(http_request_t *req);

#endif //NETWORKS_COURSE_HTTP_H
