# Use the official Golang image as a base image
FROM golang:1.24.1-bookworm

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o api-todo-list .

# Expose the port your gRPC server will run on
EXPOSE 50051

# Run the application
CMD ["./api-todo-list"]
