DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
        CREATE TYPE order_status AS ENUM (
            'pending',
            'processing',
            'shipped',
            'delivered',
            'cancelled'
        );
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY NOT NULL,
    userId INTEGER NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    status order_status NOT NULL DEFAULT 'pending',
    address TEXT NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES users(id)
);
