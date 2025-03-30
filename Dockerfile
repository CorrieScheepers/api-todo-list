# Use a Go image as the base image
FROM mcr.microsoft.com/devcontainers/go:latest AS base

# Set the working directory inside the container
WORKDIR /workspace

EXPOSE 8080

FROM mcr.microsoft.com/devcontainers/go:latest AS build
WORKDIR /workspace
COPY ["*", "."]

# Install Go dependencies (if needed)
RUN go mod tidy

# Start the shell to allow interaction
#CMD ["/bin/bash"]
