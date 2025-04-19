FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache tzdata

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go tool templ generate
RUN go tool sqlc generate

CMD ["go", "tool", "air", "-c", ".air-server.toml"]
