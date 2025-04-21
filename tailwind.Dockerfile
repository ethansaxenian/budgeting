FROM cosmtrek/air:latest

WORKDIR /app

ADD --chmod=755 https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-linux-x64 /usr/local/bin/tailwindcss

COPY . .
