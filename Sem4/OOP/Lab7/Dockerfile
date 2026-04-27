FROM golang:1.26.2-bookworm AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

FROM alpine
WORKDIR /app
COPY --from=builder server server
COPY --from=builder /app/.env .env
CMD ["./server"]
