# Use a Go base image
FROM golang:1.23.5-alpine

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8070 (or your backend's port)
EXPOSE 8070

# Run the Go application
CMD ["./main"]