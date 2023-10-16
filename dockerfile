# Use an official Go runtime as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire content of your application to the container
COPY . .

# Build the Go application inside the container
RUN go build -o main

# Expose the port that the application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
