# Use the official AWS Golang image as a base
FROM public.ecr.aws/docker/library/golang:latest

# Set the working directory inside the container
WORKDIR /app

# Install Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy necessary migration files
COPY . .

# Default command to run migrations (adjust as needed)
ENTRYPOINT ["tail", "-f", "/dev/null"]
