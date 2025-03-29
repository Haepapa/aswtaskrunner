# Use the official Golang image with the Bookworm version
FROM golang:1.24.1-bookworm

# Set the working directory inside the container
WORKDIR /app

# Install the required Go package
RUN go install github.com/sabhiram/go-wol/cmd/wol@latest

# Set the entrypoint to execute the wol command with the specified MAC address
ENTRYPOINT ["wol", "wake", "<mac>"]