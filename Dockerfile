# Use the official AWS Golang image as a base
FROM public.ecr.aws/docker/library/golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o users-service ./cmd

# Final image
FROM public.ecr.aws/docker/library/golang:latest 
# Set the working directory inside the container
WORKDIR /app

# Copy only the binary from the builder stage
COPY --from=builder /app/users-service .

# Expose the port on which the application will run
EXPOSE 8080

# Command to execute the application
CMD ["./users-service"]
