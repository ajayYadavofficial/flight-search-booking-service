# Build stage
FROM golang:1.25.5-alpine AS builder

# Install git and other dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o flight-booking-service .

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Create app directory
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/flight-booking-service .

# Expose port
EXPOSE 8080

# Set environment variables (can be overridden at runtime)
ENV MONGO_URI=mongodb://localhost:27017
ENV DB_NAME=flight_booking_db
ENV PORT=8080

# Run the application
CMD ["./flight-booking-service"]
