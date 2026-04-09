FROM golang:1.25.5-alpine AS builder

WORKDIR /app

ENV GONOSUMDB="*"
ENV GOINSECURE="*"
ENV GOPRIVATE="*"

# Install git and certificates for fetching dependencies
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN go build -o main .

# --- Final stage: minimal runtime image ---
FROM alpine:latest

WORKDIR /app

# Create uploads directory for product images
RUN mkdir -p uploads

COPY --from=builder /app/main .
COPY --from=builder /app/.env* ./

EXPOSE 8081

CMD ["./main", "start"]
