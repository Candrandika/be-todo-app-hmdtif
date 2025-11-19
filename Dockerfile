FROM golang:1.25-alpine AS builder

WORKDIR /app

# install air for hot reload
RUN go install github.com/air-verse/air@latest

# Copy file dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
# RUN go build -o main ./cmd/server

# Stage 2: Run
# FROM alpine:latest
# WORKDIR /root/
# COPY --from=builder /app/main .

# EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
