# Base Image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code from the host to the container
COPY go.mod .

# Build the Go app
RUN go mod tidy

COPY . .

RUN go build -o client client/

# Command to run the executable
CMD ["./client"]
