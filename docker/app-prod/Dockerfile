# syntax=docker/dockerfile:1

FROM golang:1.18 AS builder
WORKDIR /app/
COPY /internal/go.mod ./
COPY /internal/go.sum ./
RUN go mod download
COPY /internal/ ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /usr/local/bin/main -a -ldflags '-linkmode external -extldflags "-static"' .
# Download the static build of Litestream directly into the path & make it executable.
# This is done in the builder and copied as the chmod doubles the size.
ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.8/litestream-v0.3.8-linux-amd64-static.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz

FROM alpine
COPY --from=builder /usr/local/bin/main /usr/local/bin/main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/local/bin/litestream /usr/local/bin/litestream
COPY --from=builder /app/docs/ /docs/
RUN apk add bash
RUN mkdir -p /data
COPY service-account.json ./

# Copy Litestream configuration file & startup script.
COPY etc/litestream.yml /etc/litestream.yml
COPY scripts/run.sh /scripts/run.sh

EXPOSE 8080

CMD [ "/scripts/run.sh" ]