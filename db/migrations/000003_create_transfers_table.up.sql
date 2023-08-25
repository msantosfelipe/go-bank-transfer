CREATE TABLE IF NOT EXISTS transfers(
   id uuid PRIMARY KEY,
   account_origin_id uuid NOT NULL,
   account_destination_id uuid NOT NULL,
   amount BIGINT NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   
   CONSTRAINT fk_transfers_account_origin_id
        FOREIGN KEY (account_origin_id) REFERENCES accounts(id),
   CONSTRAINT fk_transfers_account_destination_id
        FOREIGN KEY (account_origin_id) REFERENCES accounts(id)
);
