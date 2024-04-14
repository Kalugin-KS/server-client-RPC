FROM golang:alpine

WORKDIR /client

RUN go env -w GO111MODULE=auto

ADD client.go .

RUN go build client.go

ENTRYPOINT [ "./client" ]