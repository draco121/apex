# Stage 1: Build stage
FROM golang:1.22.2-bookworm AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

RUN go get -u all

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o projectmanagerservice .

# Stage 2: Run stage
FROM debian:bookworm

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files from the build stage
COPY --from=builder /app/projectmanagerservice /app/projectmanagerservice

# Expose the port the application runs on
EXPOSE 80

# Command to run the executable
CMD ["./projectmanagerservice"]
