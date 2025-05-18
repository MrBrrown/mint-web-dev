FROM golang:1.24-alpine AS builder

RUN apk add --no-cache ca-certificates 

WORKDIR /app/gateway

COPY common /app/common

COPY gateway/go.mod  gateway/go.sum ./
RUN go mod download


COPY gateway/ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" \
        -o /gateway \
        ./cmd/gateway

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app/gateway/configs /configs

COPY --from=builder /gateway /gateway 

EXPOSE 8080

ENTRYPOINT ["/gateway"]
