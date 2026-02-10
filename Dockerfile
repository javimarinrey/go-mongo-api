# ---------- BUILD ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Dependencias
COPY go.mod go.sum ./
RUN go mod download

# Código
COPY . .

# Build optimizado
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o app ./cmd/api

# ---------- RUNTIME ----------
FROM alpine:3.19

WORKDIR /app

# Certificados TLS
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
