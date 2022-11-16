# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /golang-docker

COPY function.go function.go
COPY golang-listener.py golang-listener.py

RUN go build .

CMD [ "./golang"]