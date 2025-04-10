CREATE TABLE IF NOT EXISTS category (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO category (name) VALUES 
('Electronics'),
('Books'),
('Clothing'),
('Groceries');

CREATE TABLE IF NOT EXISTS product (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    quantity INT UNSIGNED NOT NULL DEFAULT 0,
    price INT UNSIGNED NOT NULL,
    category_id BIGINT UNSIGNED,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE SET NULL
);

INSERT INTO product (name, quantity, price, category_id) VALUES
('Smartphone', 10, 69900, 1),
('Jeans', 25, 4500, 3),
('Apples (1kg)', 50, 900, 4),
('Book: Golang Basics', 100, 3500, 2);

