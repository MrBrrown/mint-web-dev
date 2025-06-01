FROM golang:1.24-alpine AS builder

RUN apk add --no-cache ca-certificates 

WORKDIR /app/auth_service

COPY common /app/common

COPY auth_service/go.mod  auth_service/go.sum ./
RUN go mod download

COPY auth_service/ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" \
    -o /auth_service \
    ./cmd/auth_service


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app/auth_service/configs /configs

COPY --from=builder /auth_service /auth_service 

ENTRYPOINT ["/auth_service"]
