# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY sessions-api/go.mod .
COPY sessions-api/go.sum .

RUN go mod download

COPY sessions-api/*.go ./

RUN go build -o /sessions-api

EXPOSE 8010

CMD [ "/sessions-api" ]