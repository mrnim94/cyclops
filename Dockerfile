# Use the Golang Windows Server Core image as the base image
FROM golang:windowsservercore-ltsc2022

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build

# Set the entrypoint to the executable
ENTRYPOINT ["main.exe"]