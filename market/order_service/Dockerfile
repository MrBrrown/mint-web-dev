FROM golang:1.24-alpine AS builder

RUN apk add --no-cache ca-certificates 

WORKDIR /app/order_service

COPY common /app/common

COPY order_service/go.mod  order_service/go.sum ./
RUN go mod download

COPY order_service/ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" \
    -o /order_service \
    ./cmd/order_service


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app/order_service/configs /configs

COPY --from=builder /order_service /order_service 

ENTRYPOINT ["/order_service"]
