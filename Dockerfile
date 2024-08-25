# Use the official Golang image as a build stage
FROM golang:1.22.5 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Set environment variable for specifying .env file path
ENV ENV_PATH=".env"

# Build the Go application with optimized flags
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app ./cmd

# Use a minimal Alpine image for the final stage
FROM alpine:3.14

# Set timezone to UTC
ENV TZ=UTC

# Create a directory for the application
RUN mkdir /app

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary and config file from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/config.yaml .

# Copy the .env file into the container (assuming it's in the root directory of your project)
COPY .env .

## Expose port 2244 to the outside world
#EXPOSE 2244

# Command to run the executable
CMD ["./app", "--config", "config.yaml"]
