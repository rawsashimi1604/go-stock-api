# go-stock-api

## Connection String

Please use this connection string in your `.env` file.
```
POSTGRES_URL="postgres://postgres:mysecretpassword@127.0.0.1:5432/stocksdb?sslmode=disable"
```

Start a docker container and create the stocks database.
```
docker pull postgres:latest
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres
docker exec -it some-postgres bash
su postgres
psql
CREATE DATABASE stocksdb;
```

## Running with hot reload using Air
1. Install air `https://github.com/cosmtrek/air`
2. Run using cmd `air`
3. Make sure `air` is in `PATH`
