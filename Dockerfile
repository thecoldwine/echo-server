FROM golang:1.21-alpine3.19 AS build

WORKDIR /build
COPY main.go /build/
COPY go.mod /build/

RUN go build -o echo-server .

FROM alpine:3.19

RUN apk add ca-certificates && update-ca-certificates

COPY --from=build /build/echo-server /usr/local/bin/echo-server

EXPOSE 8080

CMD [ "echo-server" ]