FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache tzdata supervisor

ADD --chmod=755 https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 /usr/local/bin/tailwindcss

RUN go install github.com/air-verse/air@v1.60.0
RUN go install github.com/a-h/templ/cmd/templ@v0.2.778
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.22.1

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY prefix-log /usr/local/bin/prefix-log

CMD ["/usr/bin/supervisord", "-c", "supervisord.conf"]
