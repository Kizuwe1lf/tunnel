# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine


WORKDIR /app
COPY . ./
RUN go mod download

EXPOSE 1333
ENTRYPOINT go run .