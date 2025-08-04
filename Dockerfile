# Use an official minimal Go image
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o theopneustos .

# Use a minimal base image to run the app
FROM alpine:latest

# Create a working directory
WORKDIR /app

# Copy the binary and static files from the builder stage
COPY --from=builder /app/theopneustos .
COPY --from=builder /app/public ./public
COPY --from=builder /app/bible ./bible

# Expose the port
EXPOSE 7777

# Run the app
CMD ["./theopneustos"]
