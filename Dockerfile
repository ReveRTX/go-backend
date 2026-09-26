# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -o backend main.go

# Runtime stage
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/backend .

EXPOSE 8000

CMD ["./backend"]