# Build Stage
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install necessary packages
RUN apk update && apk add --no-cache git

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o /bin/api-cache-store ./cmd/api-cache-store

# Final Stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Install ca-certificates to handle HTTPS
RUN apk --no-cache add ca-certificates

# Copy the binary from the build stage
COPY --from=builder /bin/api-cache-store /app/api-cache-store

# Copy the .env file
COPY .env /app/.env

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["/app/api-cache-store"]