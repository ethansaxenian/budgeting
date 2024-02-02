A simple web app to track monthly expenses. Yes, it doesn't look great. I don't care.
This is just me messing around with go templ, tailwind, and htmx.

Requirements:
- go 1.21.5
- docker-compose
- tailwindcss
- make (technically not necessary)

To run the thing, just use `make`. That will:
1. install dependencies
2. start the database
3. migrate the database
4. compile tailwindcss and templ
5. run the app


Run `make help` to see everything that you can do.
