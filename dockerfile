# Start from Node.js base image for building the Nuxt.js project
FROM node:alpine as nuxt-builder

# Set the working directory
WORKDIR /nuxt-app

# Copy the Nuxt.js project files
COPY ./view/package*.json ./
COPY ./view/ ./

# Install dependencies and build the Nuxt.js project
RUN npm install && npm run generate

# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Denies Kresna <denieskresna@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Copy the built Nuxt.js project from the previous stage
COPY --from=nuxt-builder /nuxt-app/.output/public ./view/.output/public

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# Copy the Nuxt.js built files
COPY --from=builder /app/view/.output ./view/.output

# Copy  TLS file
# COPY --from=builder /app/cert/server-cert.pem ./cert/server-cert.pem
# COPY --from=builder /app/cert/server-key.pem ./cert/server-key.pem
# COPY --from=builder /app/cert/ca-cert.pem ./cert/ca-cert.pem
# COPY --from=builder /app/cert/ca-key.pem ./cert/ca-key.pem
# COPY --from=builder /app/cert/server-req.pem ./cert/server-req.pem
# COPY --from=builder /app/cert/server-ext.cnf ./cert/server-ext.cnf

# Expose port to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]