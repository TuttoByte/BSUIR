FROM golang:1.25.3-bookworm AS builder

WORKDIR /app
COPY . .

RUN go build -o server

FROM alpine
WORKDIR /app
COPY --from=builder server server
CMD["./server"]
