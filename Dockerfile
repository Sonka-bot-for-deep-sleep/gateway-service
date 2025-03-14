FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o gateway_service ./cmd/gateway_service/main.go

FROM alpine:3.21
WORKDIR /app

COPY --from=builder /app/gateway_service .
ENTRYPOINT ["./gateway_service"]
