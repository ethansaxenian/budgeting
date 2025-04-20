FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go tool sqlc generate

CMD ["go", "tool", "air", "-c", ".air-sqlc.toml"]
