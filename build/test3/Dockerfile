FROM golang:1.17.1-alpine3.14 AS builder
ADD . /usr/src/app
WORKDIR /usr/src/app
RUN go build -v -o /usr/local/bin/app ./cmd/test3

FROM alpine:3.14
COPY --from=builder /usr/local/bin/app /usr/local/bin
ENTRYPOINT ["app"]