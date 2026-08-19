# ------- Stage 1 : Build -------
  FROM golang:alpine AS builder
  WORKDIR /app
  COPY go.mod go.sum ./
  RUN go mod download 
  COPY . .
  RUN go build -o hexago-backend main.go

# ------- Stage 2 : Run -------
  FROM alpine:latest
  WORKDIR /root/
  COPY --from=builder /app/hexago-backend .
  EXPOSE 8080
  CMD ["./hexago-backend"]