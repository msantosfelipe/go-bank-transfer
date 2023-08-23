# go-bank-transfer

## Description
API for bank transfers from a digital bank

## Dependencies
* Configuration
    - Install:
        - Golang 1.21
        - Docker - run local database
        - Migrate (https://github.com/golang-migrate/migrate) - run migrations in database
        - SQLC (https://sqlc.dev/) - convert sql queries into golang code

* Golang libraries:
    - Env variables: `go-env` and `gotenv`
    - More on `go.mod`

## How to run


## Features
- Accounts
    - Get list of accounts - TODO
    - Get account balance - TODO
    - Create an account - TODO
- Login
    - Authenticate a user - TODO
        - The default logins/secrets are:
            - `87832842067` / `FirstString123@`
- Transfer
    - Get authenticated user transfers - TODO
    - Make a transfer between accounts - TODO

## License
- MIT

## Credits
- Felipe Maia Santos: https://github.com/msantosfelipe