FROM golang:1.18.0 as builder
WORKDIR /app/
COPY /internal/go.mod ./
COPY /internal/go.sum ./
RUN go mod download
COPY /internal/ ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /main
COPY /service-account-file.json ./
ENV SERVICE_ACCOUNT_FILE "./service-account-file.json"
ENV GCP_PROJECT "jerens-app"
ENV SQLITE_DB "./sqlite.db"
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/main"]