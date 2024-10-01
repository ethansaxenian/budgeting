FROM golang:1.23-alpine

WORKDIR /app

RUN apk update && \
  apk add curl tzdata && \
  curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 && \
  chmod +x tailwindcss-linux-x64 && \
  mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss && \
  go install github.com/air-verse/air@latest && \
  go install github.com/a-h/templ/cmd/templ@latest && \
  go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV TZ="America/New_York"

CMD ["air", "-c", ".air.toml"]
