CREATE TABLE loan_investments (
    id SERIAL PRIMARY KEY,
    loan_id INTEGER NOT NULL,
    investor_id INTEGER NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    paid_at TIMESTAMP NULL
);

CREATE INDEX idx_loan_investments_loan_id ON loan_investments(loan_id);
CREATE INDEX idx_loan_investments_investor_id ON loan_investments(investor_id);
CREATE INDEX idx_loan_investments_paid_at ON loan_investments(paid_at);
