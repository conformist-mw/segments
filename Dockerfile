FROM golang:1.21 AS builder

RUN apt update && apt install -y gcc-aarch64-linux-gnu

WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CGO_CFLAGS="-D_LARGEFILE64_SOURCE" CC=aarch64-linux-gnu-gcc go build -o /app/segments

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/segments /app/segments
COPY templates /app/templates
COPY static /app/static

ENV GIN_MODE=release
ENV PORT=8080
EXPOSE 8080
CMD ["./segments"]
