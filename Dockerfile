FROM golang:1.16.5-alpine3.14
RUN apk update \
    && apk add build-base
WORKDIR /app
ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download
ADD . /app

CMD ["go", "test", "./tests/..."]