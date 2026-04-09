FROM golang:1.25.5-alpine AS builder

WORKDIR /app

ENV GONOSUMDB="*"
ENV GOINSECURE="*"
ENV GOPRIVATE="*"

# Install git and certificates for fetching dependencies
RUN apk add --no-cache git ca-certificates

# Download dependencies first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the executable
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .
# Copy .env if useful, but we prefer passing env vars via compose 
# So we use wildcard to avoid error if .env doesn't exist
COPY --from=builder /app/.env* ./

EXPOSE 8085

# Execute the binary with 'start' argument since user uses 'go run main.go start'
CMD ["./main", "start"]
