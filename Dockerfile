# Build stage
FROM golang:1.18 AS builder
WORKDIR /app

# Copy the project-root folder content into the container.
# (Since the Dockerfile is inside project-root, use '.' as the context.)
COPY . .

# Download dependencies and build the main binary.
RUN go mod tidy && go build -o homie_bot ./cmd/main.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/homie_bot .

# Expose the API port (as defined in config, e.g. 8443)
EXPOSE 8443

CMD ["./homie_bot"]
