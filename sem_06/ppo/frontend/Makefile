gen_client:
    docker run \
      --rm \
      -v "${PWD}:/frontend" \
      openapitools/openapi-generator-cli generate \
      -i /frontend/api/swagger.yml \
      -g typescript-axios \
      -o /frontend/src/api/ts
