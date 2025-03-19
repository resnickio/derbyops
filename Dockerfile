# Basic Dockerfile for API
FROM golang:1.19 AS builder
WORKDIR /app
COPY api/ .
RUN go mod tidy && go build -o main

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
