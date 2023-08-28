# go-bank-transfer

## Description
API for bank transfers from a digital bank

## Requirements
- Docker - run application in a docker container
- Make - run Makefile commands
- Golang 1.21 - run app locally

## Go project dependencies
- Web: `gin`
- Database - Postgres driver: `pgx`
- Log: `logrus`
- Hash: `bcrypt`
- Env variables: `go-env` and `gotenv`
- More on `go.mod`

#### Local development dependencies
- Postgres
- Migrate (https://github.com/golang-migrate/migrate) - if you want to run migrations in database - Run `make db-migrate`
- SQLC (https://sqlc.dev/) - if you want to convert sql queries into golang code - Run `make db-queries`
- Swagger (https://github.com/swaggo/gin-swagger) - if you want to update swagger files Run `make swagger`

# How to run
## Run go-bank-transfer with docker-compose
- To start just run: `make run` or `docker-compose up`
    - Usage:
        - See 'Examples' tag bellow in this file
        - Download file `Go Bank Transfer.postman_collection.json` and import on Postman
        - Access documentation (Swagger) - http://localhost:8080/go-bank-transfer/swagger/index.html
- Stop: `make stop` or `docker-compose-down`
- Clean all data: `make clean-docker` - remove docker images and db volume

#### Run in local development (db container and local go app)
- Run app:
    - `make init-dev-db`
    - `go run .`
- Run tests: `make tests`

## Features
- Accounts:
    - Get list of accounts
    - Get account balance
    - Create an account
        - Password must have between 6 and 16 characteres
        - Accounts are created with a default balance value of 5123.56
- Login:
    - Authenticate a user
        - There app is started with a default login/secret created:
            - `87832842067` / `LetsGo321@`
- Transfer:
    - Get authenticated user transfers
    - Make a transfer between accounts

#### Future features
- Login - Forgot password feature
- Cover repositories and handlers with unit tests
- Cover project with integration tests

## Architecture
```
go-bank-transfer/
├── app/
|   ├── delivery/
|       ├── http/
|           ├── ...     # REST routers and handlers
|   ├── repository/
|       ├── ...         # Database interactions
|   ├── usecase/    
|       ├── ...         # App logic methods
├── config/
|   |── ...             # Environment variables config
├── db/
|   ├── migrations/
|       ├── ...         # Dabase migration files
|   ├── queries/
|       ├── ...         # Database SQL queries
├── domain/
|   ├── ...             # Domain entities
├── infrastructure/
|   ├── ...             # External tools used (crypto, jwt, swagger...)
├── main.go             # App entry point
├── ...                 # Other files
```

## Examples
- Login
```
Request:
curl -L 'http://localhost:8080/go-bank-transfer/login' -H 'Content-Type: application/json' --data-raw '{"cpf":"87832842067","secret":"LetsGo321@"}'

Response:
{"token":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X29yaWdpbl9pZCI6ImQwNTM3YTBlLTE2YzMtNDc0YS1hOGQ3LTdhNWZjNmExYzc5YyIsImV4cCI6MTY5MzE1NzkwNn0.NygThhB0G5jhjoRjn4yV3r_DmQfHakpxF6x4f_dxVZ0"}
```

- Account - Create user
```
Request: 
curl -L 'http://localhost:8080/go-bank-transfer/accounts' -H 'Content-Type: application/json' -d '{"name":"User0","cpf":"77788899900","secret":"TestePass"}'

Response:
{"id":"7a95de4b-0936-437f-b48e-9e898f4c99af"}
```

- Account - Get accounts
```
Request:
curl -L 'http://localhost:8080/go-bank-transfer/accounts'

Response:
{"accounts":[{"id":"d0537a0e-16c3-474a-a8d7-7a5fc6a1c79c","string":"James Bond","cpf":"87832842067","created_at":"27/08/2023 15:54:37"},{"id":"7a95de4b-0936-437f-b48e-9e898f4c99af","string":"User 0","cpf":"77788899900","created_at":"27/08/2023 16:30:45"}]}
```

- Transfer - Create transfer
```
Request:
curl -L 'http://localhost:8080/go-bank-transfer/transfers' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X29yaWdpbl9pZCI6ImQwNTM3YTBlLTE2YzMtNDc0YS1hOGQ3LTdhNWZjNmExYzc5YyIsImV4cCI6MTY5MzE1NzkwNn0.NygThhB0G5jhjoRjn4yV3r_DmQfHakpxF6x4f_dxVZ0' -H 'Content-Type: application/json' -d '{"account_destination_id":"7a95de4b-0936-437f-b48e-9e898f4c99af","amount":100.1}'

Response:
{"id":"6db5c431-0484-45f1-a8ae-25e6e4813c1a","old_account_origin_balance":5123.56,"new_account_origin_balance":5023.46,"old_account_destination_balance":5123.56,"new_account_destination_balance":5223.66}
```

- Trnasfer - Get transfers list
```
Request:
curl -L 'http://localhost:8080/go-bank-transfer/transfers' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X29yaWdpbl9pZCI6ImQwNTM3YTBlLTE2YzMtNDc0YS1hOGQ3LTdhNWZjNmExYzc5YyIsImV4cCI6MTY5MzE1NzkwNn0.NygThhB0G5jhjoRjn4yV3r_DmQfHakpxF6x4f_dxVZ0'

Response:
{"transfers":[{"id":"6db5c431-0484-45f1-a8ae-25e6e4813c1a","account_origin_id":"7a95de4b-0936-437f-b48e-9e898f4c99af","account_destination_id":"7a95de4b-0936-437f-b48e-9e898f4c99af","amount":100.1,"created_at":"27/08/2023 16:40:45"}]}
```

- Account - Get balance
```
Request:
curl -L 'http://localhost:8080/go-bank-transfer/accounts/d0537a0e-16c3-474a-a8d7-7a5fc6a1c79c/balance'

Response:
{"balance":5023.46}
```

## License
- MIT

## Credits
- Felipe Maia Santos: https://github.com/msantosfelipe