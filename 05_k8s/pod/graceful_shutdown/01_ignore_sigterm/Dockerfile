# Stage 1: Build the application
FROM golang:1.20 as builder
WORKDIR /app
# Enable static compilation
ENV CGO_ENABLED=0 GOOS=linux
# Copy all application files
COPY . .
# Build the Go application
RUN go build -o main main.go

# Stage 2: Create a lightweight image for the application
FROM alpine:latest
WORKDIR /root/
# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .
# Expose the port on which the application will run
EXPOSE 8080
# Set the entrypoint
ENTRYPOINT ["/root/main"]