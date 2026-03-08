CREATE TABLE accounts (
    id SERIAL PRIMARY KEY, -- not as optimal as uuidV7, but will use for simplicity
    document_number TEXT NOT NULL UNIQUE
);

CREATE TABLE operation_types (
    id SERIAL PRIMARY KEY,
    description TEXT
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY, -- not as optimal as uuidV7, but will use for simplicity
    account_id INTEGER NOT NULL,
    operation_type_id INTEGER NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    event_date TIMESTAMP NOT NULL,
    CONSTRAINT fk_account
        FOREIGN KEY(account_id) 
        REFERENCES accounts(id),
    CONSTRAINT fk_operation_type
        FOREIGN KEY(operation_type_id) 
        REFERENCES operation_types(id)
);
