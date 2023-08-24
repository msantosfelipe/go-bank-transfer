BEGIN;

CREATE TABLE IF NOT EXISTS accounts(
   id uuid PRIMARY KEY,
   name VARCHAR (1024) NOT NULL,
   cpf VARCHAR (11) NOT NULL,
   balance NUMERIC NOT NULL DEFAULT 0,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   CONSTRAINT fk_accounts_logins_cpf
        FOREIGN KEY (cpf) REFERENCES logins(cpf)
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO accounts (id, name, cpf) VALUES 
(gen_random_uuid(), 'James Bond', '87832842067');

COMMIT;