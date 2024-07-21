CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    borrower_id INTEGER NOT NULL,
    principal_amount NUMERIC(15, 2) NOT NULL,
    rate NUMERIC(15, 2) NOT NULL,
    roi NUMERIC(15, 2) NOT NULL,
    agreement_letter TEXT NOT NULL,
    state VARCHAR(255) NOT NULL,
    total_invested NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_loans_borrower_id ON loans(borrower_id);
CREATE INDEX idx_loans_state ON loans(state);
