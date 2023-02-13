FROM mcr.microsoft.com/windows/servercore:ltsc2019

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download and install Go
RUN powershell -Command \
    $ErrorActionPreference = 'Stop'; \
    $ProgressPreference = 'SilentlyContinue'; \
    Invoke-WebRequest -Method Get -Uri https://dl.google.com/go/go1.20.windows-amd64.msi -OutFile go.msi; \
    Start-Process -FilePath .\go.msi -ArgumentList '/quiet', '/passive', '/norestart' -NoNewWindow -Wait; \
    Remove-Item -Force go.msi

# Set the environment variables for Go
ENV GOROOT C:\Go
ENV GOPATH C:\Go\src
ENV PATH C:\Go\bin;%PATH%

# Build the Go application
RUN go build -o main.exe .

# Set the entrypoint to the executable
ENTRYPOINT ["main.exe"]