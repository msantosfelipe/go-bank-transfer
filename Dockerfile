FROM golang:1.21.0-alpine as builder

ENV GOPATH="$HOME/go"

WORKDIR $GOPATH/src/github.com/msantosfelipe/go-bank-transfer

COPY . $GOPATH/src/github.com/msantosfelipe/go-bank-transfer

RUN go build -o go-bank-transfer

RUN go test ./app/...

FROM alpine:latest

ENV GOPATH="$HOME/go"

WORKDIR /app

RUN apk update

COPY --from=builder $GOPATH/src/github.com/msantosfelipe/go-bank-transfer .

ENTRYPOINT ["./go-bank-transfer"]
