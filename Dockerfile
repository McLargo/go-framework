# Build the Go binary
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the source from the current directory
COPY . ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -tags=viper_bind_struct -o go_framework ./cmd/main.go

#--------------------------#
# Build the final image
FROM alpine:3.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/go_framework .

# Install curl for healthcheck
RUN apk --update --no-cache add curl

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the Go binary
CMD ["./go_framework"]
