FROM winamd64/golang:1.20

WORKDIR /usr/src/cyclops

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
COPY . .
RUN go build -v -o app .

CMD ["./app"]