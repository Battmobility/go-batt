# go-batt

Reusable go packages.

## Keycloak client code

Since recently, keycloak offers [an openapi spec](https://www.keycloak.org/docs-api/latest/rest-api/openapi.json).
The generation of the go code can be done using the openapitools/openapi-generator-cli docker image.
In the January 2025 version, this document has a few issues that `openapitools/openapi-generator-cli` detects.
The go code is therefore generated using the extra option  `--skip-validate-spec`.

```bash
cd pkg/keycloak/
wget https://www.keycloak.org/docs-api/latest/rest-api/openapi.json
docker run --rm -v "${PWD}:/local" --user "$(id -u):$(id -g)" openapitools/openapi-generator-cli generate -i local/openapi.json -g go -o /local/openapi --global-property apiTests=false  --skip-validate-spec
```

A few minor changes to the generated code are needed to get it compiled (see "// EDITED for compilation" in the code).
Don't commit the generated `.openapi-generator` or `.gitigore` file.
