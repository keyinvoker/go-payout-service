DROP TABLE IF EXISTS payouts;

DO $$ 
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payout_status') THEN
        DROP TYPE payout_status;
    END IF;
END $$;
