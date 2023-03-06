FROM golang:alpine AS builder

RUN apk update

WORKDIR $GOPATH/src/fizzbuzz/
COPY go.mod ./
COPY go.sum ./
COPY main.go main.go
COPY server.go server.go
COPY constants constants
COPY controllers controllers
COPY redis redis
COPY services services
COPY types types
COPY utils utils

RUN go mod tidy
RUN go mod download
RUN go mod verify

RUN go build -o /go/bin/fizzbuzz .

FROM alpine
RUN mkdir /fizzbuzz
COPY --from=builder /go/bin/fizzbuzz /go/bin/fizzbuzz
ENTRYPOINT [ "/go/bin/fizzbuzz" ]
