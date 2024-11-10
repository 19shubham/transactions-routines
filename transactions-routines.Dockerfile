# Stage 1: Build the Go application
FROM golang:1.18-alpine AS builder

# Set environment variables
ENV HOME /app
ENV GO111MODULE=on
WORKDIR ${HOME}

# Install dependencies and build the application
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Stage 2: Run the Go application
FROM alpine:latest

# Set environment variables
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./main"]