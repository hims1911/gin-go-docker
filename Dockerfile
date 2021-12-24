#build stage
ARG GO_VERSION=1.17.2

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

WORKDIR /app
ADD . /app

RUN go build -o ./out/crud-app .

EXPOSE 8080

ENTRYPOINT [ "./out/crud-app" ]

LABEL Name=CRUD_APP Version=0.0.1

