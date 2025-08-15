# Строим бинарник
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .

# Финальный образ
FROM debian:bullseye-slim
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]