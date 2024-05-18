# Genesis Software Engineering School 4.0 // Кейс

Остап Фецко &bull; o.fetsko@gmail.com

---

Stack: Go (Gin, Gorm) + PostrgeSQL

For the currency rate a mock API is used (in `.mock-api`). It mimics [Minfin API](https://minfin.com.ua/ua/developers/api/)

When the app launches, it:
1. Tries to connect to the database and run auto-migrations (if nedded).
2. Gets the initial rate from the currency-API and schedules regular updates of the rate.
3. Sets up the emails service, schedules sending of emails with current rate.
4. Sets up the API route handlers (with additional `/subs` route that returns list of subscribed emails -- for dev purposes).

## Setup

```sh
# create .env file
cp .env-example .env

# set env variables: DB-config, currency-api url, email-settings
```

### prod
```sh
docker compose up --build
```

### dev

Use [Air](https://github.com/cosmtrek/air) to restart server automatically

Use docker compose profiles to launch specific services

```sh
# install dependencies
go get .

# start DB and mock-api [+app] [+adminer]
docker compose -f compose.yaml -f compose.dev.yaml [--profile app] [--profile adminer] up --build

# run server
go run .
# or
air
```