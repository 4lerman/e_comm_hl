CREATE TYPE payment_status AS ENUM ('success', 'failed');

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    userId INT NOT NULL,
    orderId INT NOT NULL,
    amount INT NOT NULL,
    paymentDate TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status payment_status NOT NULL,
    
    FOREIGN KEY (orderId) REFERENCES orders(id),
    FOREIGN KEY (userId) REFERENCES users(id)
)