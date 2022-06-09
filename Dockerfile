# syntax=docker/dockerfile:1

FROM golang:1.18 AS builder

WORKDIR /app/
COPY /internal/go.mod ./
COPY /internal/go.sum ./
RUN go mod download
COPY /internal/ ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /main -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch
WORKDIR /app/
COPY --from=builder /main /main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV GCP_PROJECT "jerens-app"
ENV SQLITE_DB "./sqlite.db"
ENV CORS_ALLOWED_ORIGINS "https://www.jerenslensun.com"
ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["/main"]