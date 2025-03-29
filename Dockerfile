# Builder stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copiar primero los módulos para aprovechar el cache de Docker
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Compilar el binario estático
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /laliga-tracker

# Stage final
FROM alpine:3.20

WORKDIR /app

# Copiar el binario desde el builder
COPY --from=builder /laliga-tracker /app/laliga-tracker

# Instalar certificados CA y timezone data
RUN apk add --no-cache ca-certificates tzdata

# Puerto expuesto
EXPOSE 8080

# Comando de ejecución
CMD ["/app/laliga-tracker"]