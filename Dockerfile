FROM golang:latest

WORKDIR /build

COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /server server.go

ENTRYPOINT [ "/server" ]