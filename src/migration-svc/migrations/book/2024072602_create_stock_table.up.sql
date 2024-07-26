CREATE TABLE stock (
    stock_id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(book_id),
    total_quantity INT NOT NULL,
    available_quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);