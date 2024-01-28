FROM golang:1.21-alpine AS builder

RUN apk update && apk add gcc libc-dev

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go build -o /app/segments

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/segments /app/segments
COPY templates /app/templates
COPY static /app/static

ENV GIN_MODE=release
ENV PORT=8080
EXPOSE 8080
CMD ["./segments"]
