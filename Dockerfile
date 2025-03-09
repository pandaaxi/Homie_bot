# Build stage
FROM golang:1.18 AS builder
WORKDIR /app

# Copy the entire repository into the container
COPY . .

# Download dependencies and build the main binary
RUN go mod tidy && go build -o homie_bot ./cmd/main.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/homie_bot .

# Expose the API port (update if needed)
EXPOSE 8443

CMD ["./homie_bot"]
