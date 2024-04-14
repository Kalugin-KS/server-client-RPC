FROM golang:alpine

WORKDIR /server

RUN go env -w GO111MODULE=auto

ADD server.go .

RUN go build server.go

ENTRYPOINT [ "./server" ]