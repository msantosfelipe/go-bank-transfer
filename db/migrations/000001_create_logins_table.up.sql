BEGIN;

CREATE TABLE IF NOT EXISTS logins(
   cpf VARCHAR (11) PRIMARY KEY,
   secret VARCHAR (1024) NOT NULL
);

INSERT INTO logins (cpf, secret) VALUES 
('87832842067', '$2a$10$DrTEM2HLEZFizo3Xt1lIVe8H8E5mRXPOm7pXH/k06ShbLgf.A5dLK');

COMMIT;