# Use official Golang image as a base
FROM golang:1.24

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files, and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o payment-service

# Expose the application's port
EXPOSE 8081

# Run the application
CMD ["./payment-service"]
