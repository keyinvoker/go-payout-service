DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payout_status')
    THEN
        CREATE TYPE payout_status AS ENUM (
            'PENDING',
            'CALCULATION_FAILED',
            'READY_TO_PAYOUT',
            'ON_PROCESS',
            'PAYOUT_FAILED',
            'PAID_OUT'
        );
    END IF;
END $$;

CREATE TABLE payouts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP NULL,
    loan_id INT NOT NULL,
    user_id INT NOT NULL,
    payout_status payout_status NOT NULL DEFAULT 'PENDING',
    payout_date TIMESTAMP NULL,
    total DECIMAL(18,2) NOT NULL DEFAULT 0,
    principal DECIMAL(18,2) NOT NULL DEFAULT 0,
    interest DECIMAL(18,2) NOT NULL DEFAULT 0,
    fine DECIMAL(18,2) NOT NULL DEFAULT 0,
    description TEXT NULL,
    CONSTRAINT payouts_loan_id_key INDEX (loan_id),
    CONSTRAINT payouts_user_id_key INDEX (user_id)
);
