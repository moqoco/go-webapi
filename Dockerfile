FROM golang:1.16-alpine

WORKDIR /go/app
COPY ./app .

RUN apk upgrade --update && \
    apk --no-cache add git curl

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air


CMD ["air", "-c", ".air.toml"]