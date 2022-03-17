.PHONY: openapi_http
openapi_http:
	~/oapi-codegen -generate types -o internal/ports/openapi_types.gen.go -package ports api/openapi/notification.yml
	~/oapi-codegen -generate chi-server -o internal/ports/openapi.gen.go -package ports api/openapi/notification.yml