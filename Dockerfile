# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint compared to Ubuntu
FROM golang:1.16-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

# Server Port
EXPOSE 8080

CMD [ "/docker-gs-ping" ]