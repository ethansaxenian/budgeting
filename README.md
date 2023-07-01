A simple web app to track monthly expenses. Yes, it doesn't look great. I don't care.

Create a `.env` file and customize the following environment variables as you see fit:

```shell
WEB_PORT=8000
API_PORT=8001
DB_PORT=8002
DB_USER=postgres
DB_PASSWORD=postgres
```

The values listed above will be the default.

Make sure you have the docker daemon running, and simply run

```shell
docker-compose up -d
```

to start everything on localhost.

The web, api, and db will be running on the ports you chose above.
