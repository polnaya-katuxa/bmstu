#include <stdio.h>
#include <sys/stat.h>
#include <unistd.h>
#include <string.h>
#include <fcntl.h>

#include "http.h"
#include "errors.h"

#define PAGE_SIZE 4096

/*
 * http_parse_request парсит запрос, получает версию протокола HTTP, метод и URI.\
 *
 * GET /wiki/HTTP HTTP/1.0
 */
int http_parse_request(char *raw, http_request_t *req) {
    char method_str[16] = { 0 };

    if (sscanf(raw, "%s %s %s\n", method_str, req->uri, req->version) != 3) {
        return ERR_INVALID_REQUEST;
    }

    if (strcmp(req->version, "HTTP/1.1") != 0 && strcmp(req->version, "HTTP/1.0") != 0) {
        return ERR_UNSUPPORTED_HTTP_VERSION;
    }

    int len = sizeof(http_methods_map) / sizeof(char*);

    for (int i = 0; i < len; i++) {
        if (strcmp(method_str, http_methods_map[i]) == 0) {
            req->method = i;
            return OK;
        }
    }

    return ERR_UNSUPPORTED_METHOD;
}

char *responses_map[] = {
    [RESPONSE_CODE_OK] = "200 OK",
    [RESPONSE_CODE_FORBIDDEN] = "403 Forbidden",
    [RESPONSE_CODE_NOT_FOUND] = "404 Not Found",
    [RESPONSE_CODE_METHOD_NOT_ALLOWED] = "405 Method Not Allowed",
    [RESPONSE_CODE_BAD_REQUEST] = "400 Bad Request",
    [RESPONSE_CODE_INTERNAL_SERVER_ERROR] = "500 Internal Server Error",
};

/*
 * http_write_status записывает в сокет первую строку ответа.
 *
 * HTTP/1.0 200 OK
 */
int http_write_status(http_request_t *req, http_response_code_t code) {
    char buf[64] = { 0 };
    if (sprintf(buf, "%s %s\n", req->version, responses_map[code]) < 0) {
        return ERR_CANT_SPRINTF;
    }

    if (write(req->fd, buf, strlen(buf)) == -1) {
        return ERR_CANT_WRITE_TO_SOCKET;
    }

    return OK;
}

/*
 * http_write_error выдает ошибочный HTTP-ответ.
 * Отдает статус, заголовки, и тело с текстом проблемы.
 */
int http_write_error(http_request_t *req, http_response_code_t code) {
    int rc;
    rc = http_write_status(req, code);
    if (rc != 0) {
        return rc;
    }

    char *message = responses_map[code];

    char content_length_header[128] = { 0 };
    sprintf(content_length_header, "Content-Length: %lu\n", strlen(message));
    if (write(req->fd, content_length_header, strlen(content_length_header)) == -1) {
        return ERR_CANT_WRITE_TO_SOCKET;
    }

    char *content_type_header = "Content-Type: text/plain\n\n";
    if (write(req->fd, content_type_header, strlen(content_type_header)) == -1) {
        return ERR_CANT_WRITE_TO_SOCKET;
    }

    if (write(req->fd, message, strlen(message)) == -1) {
        return ERR_CANT_WRITE_TO_SOCKET;
    }

    return OK;
}

/*
 * http_copy_segment копирует кусок из одного файла в другой.
 */
int http_copy_segment(int dst, int src) {
    char buf[PAGE_SIZE] = { 0 };

    ssize_t n = read(src, buf, PAGE_SIZE);
    if (n < 0) {
        return ERR_READ;
    }

    if (n == 0) {
        return ERR_EOF;
    }

    if (write(dst, buf, n) < 0) {
        return ERR_WRITE;
    }

    return OK;
}

/*
 * http_copy копирует содержимое одного файла в другой.
 */
int http_copy(int dst, int src) {
    int rc = OK;
    while (rc == OK) {
        rc = http_copy_segment(dst, src);
    }

    if (rc != ERR_EOF) {
        return rc;
    }

    return OK;
}

int is_file_exists(char *filename) {
    return access(filename, F_OK) == 0;
}

int is_path_safe(char *filename) {
    return strstr(filename, "../") == NULL && filename[0] != '/';
}

typedef struct {
    char *extension;
    char *content_type;
} content_type_mapping_t;

content_type_mapping_t content_type_map[] = {
    {
        .extension = ".html",
        .content_type = "text/html",
    },
    {
        .extension = ".css",
        .content_type = "text/css",
    },
    {
        .extension = ".js",
        .content_type = "text/javascript",
    },
    {
        .extension = ".png",
        .content_type = "image/png",
    },
    {
        .extension = ".jpg",
        .content_type = "image/jpeg",
    },
    {
        .extension = ".jpeg",
        .content_type = "image/jpeg",
    },
    {
        .extension = ".swf",
        .content_type = "application/vnd.adobe.flash-movie",
    },
    {
        .extension = ".gif",
        .content_type = "image/gif",
    },
};

char *http_get_content_type(char *filename) {
    int len = sizeof(content_type_map) / sizeof(content_type_mapping_t);

    for (int i = 0; i < len; i++) {
        char *suffix = filename + strlen(filename) - strlen(content_type_map[i].extension);
        if (strcmp(suffix, content_type_map[i].extension) == 0) {
            return content_type_map[i].content_type;
        }
    }

    return NULL;
}

int http_write_file_headers(http_request_t *req, char *filename, struct stat info) {
    char content_length_header[128] = { 0 };
    sprintf(content_length_header, "Content-Length: %ld\n", info.st_size);
    if (write(req->fd, content_length_header, strlen(content_length_header)) == -1) {
        return ERR_CANT_WRITE_TO_SOCKET;
    }

    char *content_type = http_get_content_type(filename);
    if (content_type != NULL) {
        char content_type_header[128] = { 0 };
        sprintf(content_type_header, "Content-Type: %s\n", content_type);
        if (write(req->fd, content_type_header, strlen(content_type_header)) == -1) {
            return ERR_CANT_WRITE_TO_SOCKET;
        }
    }

    return OK;
}

/*
 * http_serve_file раздает статический файл с диска.
 */
int http_serve_file(http_request_t *req, char *filename) {
    if (!is_path_safe(filename)) {
        http_write_error(req, RESPONSE_CODE_FORBIDDEN);
        return ERR_FORBIDDEN_PATH;
    }

    if (!is_file_exists(filename)) {
        http_write_error(req, RESPONSE_CODE_NOT_FOUND);
        return ERR_FILE_NOT_FOUND;
    }

    struct stat info;
    if (stat(filename, &info) < 0) {
        http_write_error(req, RESPONSE_CODE_INTERNAL_SERVER_ERROR);
        return ERR_FILESYSTEM;
    }

    if (!S_ISREG(info.st_mode)) {
        http_write_error(req, RESPONSE_CODE_BAD_REQUEST);
        return ERR_UNSUPPORTED_FILE;
    }

    http_write_status(req, RESPONSE_CODE_OK);

    int rc = http_write_file_headers(req, filename, info);
    if (rc != 0) {
        return rc;
    }

    if (write(req->fd, "\n", 1) == -1) {
        return ERR_CANT_WRITE_TO_SOCKET;
    }

    if (req->method == METHOD_HEAD) {
        return OK;
    }

    int fd = open(filename, O_RDONLY);
    if (fd < 0) {
        http_write_error(req, RESPONSE_CODE_INTERNAL_SERVER_ERROR);
        return ERR_FILESYSTEM;
    }

    rc = http_copy(req->fd, fd);
    if (rc != 0) {
        close(fd);
        return rc;
    }

    close(fd);

    return OK;
}

/*
 * http_handle обрабатывает HTTP-запрос.
 */
int http_handle(http_request_t *req) {
    if (strcmp(req->uri, "/") == 0) {
        return http_serve_file(req, "index.html");
    }

    // /hello.txt -> hello.txt
    char *filename = &(req->uri[1]);

    return http_serve_file(req, filename);
}

