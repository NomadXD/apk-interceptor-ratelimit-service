# Use the official Golang image as the base image
FROM golang:1.17-alpine3.14 AS build

RUN apk update && \
    apk add --no-cache build-base

# Set the working directory to /app
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod ./

# Download the Go modules
RUN go mod download

# Copy the server.go file to the container
COPY server.go ./

# Build the Go application and statically link the libraries
RUN go build -ldflags "-linkmode external -extldflags '-static'" -o server

# Use the official Alpine Linux image as the base image for the final image
FROM alpine:3.14

# Set the working directory to /app
WORKDIR /app

# Copy the server binary from the build stage to the final image
COPY --from=build /app/server .

# Expose port 8080 to the host
EXPOSE 8080

# Run the server binary when the container starts
CMD ["./server"]