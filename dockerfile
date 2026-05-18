# Build stage: Using Golang base image with Alpine
FROM golang:1.20-alpine as builder

# Install git and other dependencies using apk (Alpine package manager)
RUN apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download 

# Copy the source code from the current directory to the working directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage: Using a minimal base image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose the port that your Go app listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
