A simple web app to track monthly expenses.

Requirements:
- `make`
- `docker-compose`

### Running the app
1. Create a `.env` file with the following contents:
```sh
APP_PORT=
DB_PORT=
DB_USER=
DB_PASSWORD=
```
Any variables you set in `.env` will be set in the docker environment, so you can set the timezone with `TZ`, for example.

2. Run `make start`. This will build and run 6 docker containers:
    - `db`: The postgresql database, accessible locally on port `DB_PORT`.
    - `server`: The main go webserver. Automatically reloads when app code is changed.
    - `proxy`: A proxy for the web app that automatically reload the browser when changes are made to `.templ` files. This container exposes the app on port `APP_PORT`.
    - `tailwind`: Rebuilds the main css file when changes are made.
    - `sqlc`: Regenerates the sqlc files when changes are made to `queries/*.sql`.
    - `sync_assets`: Sends a reload event to `proxy` to reload the browser when changes are made to the `assets/dist` folder
3. Run `make migrate` to set up the database.
4. Visit `http://localhost:<APP_PORT>` in your web browser.

