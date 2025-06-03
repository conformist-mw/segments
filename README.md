# Segments accounting

This project allows you doing segments (rectangles) accounting by size, type or color. Removed segments doesn't remove from database just mark them as inactive. Also removed segments have order number. 

## Installation

To build and run the application locally use Go:

```shell
go build -o segments
./segments
```

### Docker

The application can also be built and run with Docker:

```shell
docker build -t segments .
docker run -p 8080:8080 segments
```

## Regular users

Users are not allowed to login but if it is the need, they should be in the `users` group.

