FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
RUN go mod download
RUN go build -o /app/segments

FROM scratch

WORKDIR /app
COPY --from=builder /app/segments /app/segments
COPY templates /app/templates
COPY static /app/static

ENV GIN_MODE=release
ENV PORT=8080
EXPOSE 8080
CMD ["./segments"]
