# build
FROM golang:1.24-bookworm AS builder

COPY . /build

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o controller main.go

# alpine
FROM alpine:3.21

RUN apk add ca-certificates

COPY --from=builder /build/controller /controller

WORKDIR /

CMD ["./controller"]