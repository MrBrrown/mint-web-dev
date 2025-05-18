FROM golang:1.24-alpine AS builder

RUN apk add --no-cache ca-certificates 

WORKDIR /app/product_service

COPY common /app/common

COPY product_service/go.mod  product_service/go.sum ./
RUN go mod download

COPY product_service/ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" \
        -o /product_service \
        ./cmd/product_service

        
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app/product_service/configs /configs

COPY --from=builder /product_service /product_service 

ENTRYPOINT ["/product_service"]
