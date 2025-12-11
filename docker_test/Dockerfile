FROM golang:latest AS builder

WORKDIR /app


COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:latest

WORKDIR /app


# Copy built binary from builder
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Run app
CMD ["./server"]