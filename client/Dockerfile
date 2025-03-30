# Use Golang base image
FROM golang:1.24.1-bookworm AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project (including proto and repository)
COPY . .

# Build the server application, output as 'api-todo-list'
RUN go build -o api-todo-list .

# Expose the server's port
EXPOSE 4000

# Command to run the server
CMD ["./api-todo-list"]
