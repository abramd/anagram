FROM golang:1.12.7 as build

WORKDIR /go/src/github.com/abramd/anagram
COPY ./ ./

RUN go build -a -o /go/bin/app ./cmd/anagram

#FROM scratch
FROM debian:stretch-slim

COPY --from=build /go/bin/app /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /
ENTRYPOINT ["/app"]