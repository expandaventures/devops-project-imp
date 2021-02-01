FROM golang:latest

LABEL maintainer="g_marco1_antonio@hotmail.com"

WORKDIR /app

COPY src .

RUN go get -u github.com/gorilla/mux

RUN go build

RUN go test -v

CMD ["./exam"]