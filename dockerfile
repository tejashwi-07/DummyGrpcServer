# Use the official Go module image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /go/src/engine


# Copy the entire server directory to the container's working directory
COPY . .

# Build the Go binary of the gRPC server
RUN go build -o engine-server

# Expose port 10000 for gRPC server
EXPOSE 10000

# Set the command to run the gRPC server
CMD ["./engine-server"]
