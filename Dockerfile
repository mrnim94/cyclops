FROM winamd64/golang:1.20

WORKDIR /usr/src/cyclops

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY nim .
RUN go build -v -o /usr/local/bin/cyclops ./...

CMD ["cyclops"]