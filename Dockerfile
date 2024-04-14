FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application with optimizations
RUN cd cmd && go build -ldflags="-s -w" -trimpath -o raven


# Use a minimal base image for the final image
FROM alpine:3.14

# Copy the built binary from the builder stage
COPY --from=builder /app/cmd/raven /cmd/raven

# Set the working directory inside the container
WORKDIR /cmd

# Set the entry point for the container
ENTRYPOINT [ "/cmd/raven" ]