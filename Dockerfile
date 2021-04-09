FROM golang:1.16.3-buster AS builder


WORKDIR /stark
COPY . .

RUN go install -v ./...

FROM alpine:3.8

COPY --from=builder /go/bin/stark /stark

ENTRYPOINT ["/stark"]
