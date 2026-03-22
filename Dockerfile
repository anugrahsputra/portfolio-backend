# Stage 1: Build
FROM golang:1.25-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Install necessary build tools (optional, but good for some dependencies)
RUN apk add --no-cache git gcc musl-dev

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
# .dockerignore should handle excluding unnecessary files
COPY . .

# Build the Go app
# Use CGO_ENABLED=0 to ensure a static binary that works in alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# Stage 2: Final (Run)
FROM alpine:latest

# Set environment variables
ENV ENV=production

# Install ca-certificates (needed for HTTPS requests) and tzdata
RUN apk add --no-cache ca-certificates tzdata

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary from the previous stage
COPY --from=builder /app/main .

# Copy the SQL schema file (required for INIT_SCHEMA=true)
COPY --from=builder /app/sql/schema/schema.sql sql/schema/schema.sql

# Expose port 8082 to the outside world
EXPOSE 8082

# Command to run the executable
ENTRYPOINT ["./main"]
