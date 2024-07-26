CREATE TABLE borrowing (
    borrowing_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id),
    book_id INT NOT NULL,
    borrow_date DATE NOT NULL,
    due_date DATE NOT NULL,
    return_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);