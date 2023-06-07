# Use the official Go image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . . 

# Build the Go application
RUN go build -o main .

# Expose the port on which the application listens and the Mysql port
EXPOSE 8000
EXPOSE 3306

# Run the Go application
CMD ["./main"]
