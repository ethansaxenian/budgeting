FROM busybox

WORKDIR /app

ADD --chmod=755 https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-linux-x64 /usr/local/bin/tailwindcss

COPY . .

CMD ["tailwindcss", "-i", "./assets/main.css", "-o", "./assets/dist/tailwind.css", "--minify", "--watch"]
