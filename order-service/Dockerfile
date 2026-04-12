FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o order-service .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/order-service /app/order-service

EXPOSE 8082

CMD ["/app/order-service"]
