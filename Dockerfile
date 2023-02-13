# Use the Golang Windows Server Core image as the base image
FROM golang:1.20.0-windowsservercore-1809

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build

# Set the entrypoint to the executable
ENTRYPOINT ["main.exe"]