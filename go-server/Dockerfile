FROM golang:1.12.7-alpine3.10 AS builder
RUN apk add git

RUN mkdir -p /go/src/github.com/moznobkin/life-mf-mock
COPY . /go/src/github.com/moznobkin/life-mf-mock
WORKDIR /go/src/github.com/moznobkin/life-mf-mock
ENV GO111MODULE=on
RUN go build -v -o /go/bin/life-mf-mock.exe

FROM alpine:3.10
LABEL maintainer="Michael Oznobkin <oznobkin@gmail.com>"

COPY --from=builder /go/bin/life-mf-mock.exe /usr/bin/life-mf-mock.exe
COPY --from=builder /go/src/github.com/moznobkin/life-mf-mock/examples/json/  /usr/examples/json/

WORKDIR /usr/bin

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/life-mf-mock.exe"]