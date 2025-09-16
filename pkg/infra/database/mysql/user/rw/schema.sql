-- Database: test_db
-- Table: orders

CREATE TABLE orders (
    order_id VARCHAR(36) PRIMARY KEY,
    order_qty INT NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);