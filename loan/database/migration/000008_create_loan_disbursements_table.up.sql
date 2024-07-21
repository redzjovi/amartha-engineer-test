CREATE TABLE loan_disbursements (
    id SERIAL PRIMARY KEY,
    loan_id INTEGER NOT NULL,
    field_officer_id INTEGER NOT NULL,
    agreement_letter TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_loan_disbursements_loan_id ON loan_disbursements(loan_id);
CREATE INDEX idx_loan_disbursements_field_officer_id ON loan_disbursements(field_officer_id);
