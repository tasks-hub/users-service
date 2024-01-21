# Use the official AWS Golang image as a base
FROM public.ecr.aws/docker/library/golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Install Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy necessary migration files
COPY . .

# Use the official AWS Golang image as a base
FROM public.ecr.aws/docker/library/golang:latest

COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Set the working directory inside the container
WORKDIR /app

COPY --from=builder /app/run_migrations.sh .
COPY --from=builder /app/*.sql .

# Default command to run migrations (adjust as needed)
ENTRYPOINT ["tail", "-f", "/dev/null"]
