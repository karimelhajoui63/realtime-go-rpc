# Use an official Go runtime as a parent image
# Intentionnaly not the alpine version to have acces to bash
FROM golang:1.21

# # Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# ADd script to ensure that RabbitMQ server is running before trying to connect to it
RUN wget -O ./wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x ./wait-for-it.sh

# Build
RUN go build -o bin/rpcserver cmd/rpcserver/main.go

# Expose the port
EXPOSE 8080

# Run
CMD ["bin/rpcserver"]