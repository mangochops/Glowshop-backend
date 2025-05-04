# syntax=docker/dockerfile:1
# Build Stage
FROM golang:1.20.13-bullseye as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Final Stage - Create the minimal image
FROM debian:bullseye-slim

# Set the working directory inside the container
WORKDIR /app

# Install security updates and remove package lists to reduce image size
RUN apt-get update && apt-get upgrade -y --no-install-recommends && rm -rf /var/lib/apt/lists/*

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 (the port your Go app listens on)
EXPOSE 8080

# Set the default command to run the Go binary
CMD ["./main"]

