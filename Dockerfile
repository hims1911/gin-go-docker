#build stage
ARG GO_VERSION=1.17.2

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api


COPY . .
RUN go mod download

RUN go build -o ./app ./src/main.go


#final stage
FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /api
COPY --from=builder /api/app .
COPY --from=builder /api/test.db .

ENTRYPOINT ["./app"]
LABEL Name=golangexample1 Version=0.0.1
EXPOSE 3000
