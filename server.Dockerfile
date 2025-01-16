FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache tzdata

RUN go install github.com/air-verse/air@v1.60.0
RUN go install github.com/a-h/templ/cmd/templ@v0.2.778
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.22.1

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["air", "-c", ".air-server.toml"]
