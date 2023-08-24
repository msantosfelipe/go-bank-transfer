# go-bank-transfer

## Description
API for bank transfers from a digital bank

## Requirements
- Golang 1.21
- Docker - if you want to run local database
- Migrate (https://github.com/golang-migrate/migrate) - if you want to run migrations in database
- SQLC (https://sqlc.dev/) - if you want to convert sql queries into golang code
- Swagger (https://github.com/swaggo/gin-swagger) - if you want to update swagger files

## Dependencies
- Web: `gin`
- DB - Postgres driver: `pgx`
- Log: `logrus`
- Hash: `bcrypt`
- Env variables: `go-env` and `gotenv`
- More on `go.mod`

## How to run
- Run tests - `make tests`
- Documentation (Swagger) - http://localhost:8081/go-bank-transfer/swagger/index.html

## Features
- Accounts
    - Get list of accounts - TODO
    - Get account balance - TODO
    - Create an account
- Login
    - Authenticate a user - TODO
        - The default logins/secrets are:
            - `87832842067` / `LetsGo321@`
- Transfer
    - Get authenticated user transfers - TODO
    - Make a transfer between accounts - TODO

## License
- MIT

## Credits
- Felipe Maia Santos: https://github.com/msantosfelipe