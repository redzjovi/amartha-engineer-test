CREATE TABLE loan_approvals (
    id SERIAL PRIMARY KEY,
    loan_id INTEGER NOT NULL,
    validator_id INTEGER NOT NULL,
    picture_proof TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_loan_approvals_loan_id ON loan_approvals(loan_id);
CREATE INDEX idx_loan_approvals_validator_id ON loan_approvals(validator_id);
