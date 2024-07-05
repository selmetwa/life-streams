# Use the official Golang image as a base image
FROM golang:1.22.3-alpine AS build

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the application code
COPY . .

# Enable CGO and build the Go application
ENV CGO_ENABLED=1
RUN go build -o main cmd/api/main.go

# Use a minimal image for the final stage
FROM alpine:latest

# Install the necessary C library dependencies
RUN apk add --no-cache libc6-compat

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go binary from the build stage
COPY --from=build /app/main .

# Copy the SQLite database file
COPY database.db /app/database.db

# Expose the port the app runs on
ENV PORT=3000
EXPOSE 3000

# Command to run the application
CMD ["./main"]
