.PHONY: openapi_http
openapi_http:
	oapi-codegen -generate types -o internal/ports/openapi_types.gen.go -package ports api/openapi/notification.yml
	oapi-codegen -generate chi-server -o internal/ports/openapi.gen.go -package ports api/openapi/notification.yml
	oapi-codegen -generate types -o internal/tests/client/openapi_types.gen.go -package client api/openapi/notification.yml
	oapi-codegen -generate client -o internal/tests/client/openapi_client.gen.go -package client api/openapi/notification.yml

.PHONY: openapi_documentation
openapi_documentation:
	docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/api/openapi/notification.yml \
  -g html2 \
  -o /local/documentation/openapi

.PHONY: diagram_c4
diagram_c4:
	docker run -d -p 8080:8080 plantuml/plantuml-server:jetty

.PHONY: key_generate
key_generate:
	base64 service-account-key.json > service-account-key.base64