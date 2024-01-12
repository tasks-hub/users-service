# Use the official Golang image as a base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o users-service ./cmd

# Expose the port on which the application will run
EXPOSE 8080

# Command to execute the application
CMD ["./users-service"]