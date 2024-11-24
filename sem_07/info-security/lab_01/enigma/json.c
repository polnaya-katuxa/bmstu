//
// Created by Екатерина Карпова on 27.09.2023.
//

#include "enigma.h"
#include <cjson/cJSON.h>

int unmarshal_enigma(char* str, enigma_t* enigma) {
    const cJSON *size_alphabet = NULL;
    const cJSON *size_rotors = NULL;
    const cJSON *rotor = NULL;
    const cJSON *rotors = NULL;
    const cJSON *reflector = NULL;
    const cJSON *panel = NULL;
    const cJSON *code = NULL;

    cJSON *enigma_json = cJSON_Parse(str);
    if (enigma_json == NULL)
    {
        const char *error_ptr = cJSON_GetErrorPtr();
        if (error_ptr != NULL)
        {
            printf("error: %s\n", error_ptr);
        }

        return ERR_JSON_PARSE;
    }

    size_alphabet = cJSON_GetObjectItemCaseSensitive(enigma_json, "size_alphabet");
    if (!cJSON_IsNumber(size_alphabet))
    {
        printf("error: size_alphabet\n");
        cJSON_Delete(enigma_json);
        return ERR_JSON_PARSE;
    }

    enigma->size_alphabet = size_alphabet->valueint;

    size_rotors = cJSON_GetObjectItemCaseSensitive(enigma_json, "size_rotors");
    if (!cJSON_IsNumber(size_rotors))
    {
        printf("error: size_rotors\n");
        cJSON_Delete(enigma_json);
        return ERR_JSON_PARSE;
    }

    enigma->size_rotors = size_rotors->valueint;

    int rc = alloc_enigma(enigma);
    if (rc != EXIT_SUCCESS) {
        return ERR_ALLOC;
    }

    int count = 0;
    rotors = cJSON_GetObjectItemCaseSensitive(enigma_json, "rotors");
    cJSON_ArrayForEach(rotor, rotors)
    {
        cJSON_ArrayForEach(code, rotor)
        {
            if (!cJSON_IsNumber(code)) {
                printf("error: code %d\n", code->valueint);
                cJSON_Delete(enigma_json);
                return ERR_JSON_PARSE;
            }

            if ((code->valueint < 0) && (code->valueint >= enigma->size_alphabet)) {
                printf("error: code %d\n", code->valueint);
                cJSON_Delete(enigma_json);
                return ERR_JSON_PARSE;
            }

            enigma->rotors[count / enigma->size_alphabet][count % enigma->size_alphabet] = code->valueint;

            count++;
        }
    }

    count = 0;
    reflector = cJSON_GetObjectItemCaseSensitive(enigma_json, "reflector");
    cJSON_ArrayForEach(code, reflector)
    {
        if (!cJSON_IsNumber(code)) {
            printf("error: code %d\n", code->valueint);
            cJSON_Delete(enigma_json);
            return ERR_JSON_PARSE;
        }

        if ((code->valueint < 0) && (code->valueint >= enigma->size_alphabet)) {
            printf("error: code %d\n", code->valueint);
            cJSON_Delete(enigma_json);
            return ERR_JSON_PARSE;
        }

        enigma->reflector[count] = code->valueint;

        count++;
    }

    count = 0;
    panel = cJSON_GetObjectItemCaseSensitive(enigma_json, "panel");
    cJSON_ArrayForEach(code, panel)
    {
        if (!cJSON_IsNumber(code)) {
            printf("error: code %d\n", code->valueint);
            cJSON_Delete(enigma_json);
            return ERR_JSON_PARSE;
        }

        if ((code->valueint < 0) && (code->valueint >= enigma->size_alphabet)) {
            printf("error: code %d\n", code->valueint);
            cJSON_Delete(enigma_json);
            return ERR_JSON_PARSE;
        }

        enigma->panel[count] = code->valueint;

        count++;
    }

    cJSON_Delete(enigma_json);

    return EXIT_SUCCESS;
}
