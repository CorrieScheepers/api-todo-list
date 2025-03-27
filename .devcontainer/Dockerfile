# Use a Go image as the base image
FROM mcr.microsoft.com/devcontainers/go:latest

# Install Docker and Docker Compose
RUN apt-get update && \
    apt-get install -y curl lsb-release && \
    curl -fsSL https://get.docker.com -o get-docker.sh && \
    sh get-docker.sh && \
    curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \
    chmod +x /usr/local/bin/docker-compose

# Set the working directory inside the container
WORKDIR /workspace

# Install Go dependencies (if needed)
# RUN go mod tidy

# Start the shell to allow interaction
CMD ["/bin/bash"]
