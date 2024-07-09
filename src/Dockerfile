# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

LABEL name="app-image-builder"
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o task-manager ./cmd/main.go

# Stage 2: Run the Go application
FROM alpine:latest

LABEL name="app-image"

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/config ./config

CMD ["./main"]
