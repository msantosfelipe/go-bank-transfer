# go-bank-transfer

## Description
API for bank transfers from a digital bank

## Requirements
- Docker - run application in a docker container
- Make - run Makefile commands
- Golang 1.21 - run locally

## Dependencies
- Web: `gin`
- Database - Postgres driver: `pgx`
- Log: `logrus`
- Hash: `bcrypt`
- Env variables: `go-env` and `gotenv`
- More on `go.mod`

# How to run
#### Run app with Docker
- Start: `make run` or `docker-compose up`
- Stop: `make stop` or `docker-compose-down`

#### Development
- Run tests - `make tests`
- Documentation (Swagger) - http://localhost:8081/go-bank-transfer/swagger/index.html
- Migrate (https://github.com/golang-migrate/migrate) - if you want to run migrations in database - Run `make db-migrate`
- SQLC (https://sqlc.dev/) - if you want to convert sql queries into golang code - Run `make db-queries`
- Swagger (https://github.com/swaggo/gin-swagger) - if you want to update swagger files Run `make swagger`

## Features
- Accounts:
    - Get list of accounts
    - Get account balance
    - Create an account
- Login:
    - Authenticate a user
        - The default logins/secrets are:
            - `87832842067` / `LetsGo321@`
- Transfer:
    - Get authenticated user transfers
    - Make a transfer between accounts

#### Future features
- Login - Forgot password feature
- Cover repositories and handlers with unit tests
- Cover project with integration tests

## Examples
```
func main() {
    fmt.Println("Hello, world!")
}
```

## Architecture


## License
- MIT

## Credits
- Felipe Maia Santos: https://github.com/msantosfelipe