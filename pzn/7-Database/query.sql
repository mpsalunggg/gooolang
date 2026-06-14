-- INSERT INTO customer(id, name)
-- VALUES('john', 'John Doe');

-- DELETE FROM customer;

-- ALTER TABLE customer
--     ADD COLUMN email VARCHAR(100),
--     ADD COLUMN balance INTEGER DEFAULT 0,
--     ADD COLUMN rating DOUBLE DEFAULT 0.0,
--     ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     ADD COLUMN birth_date DATE,
--     ADD COLUMN married BOOLEAN DEFAULT false;
    
-- INSERT INTO customer (id, name, email, balance, rating, created_at, birth_date, married) VALUES
--       ('id1', 'Alice Smith', 'alice@example.com', 10000, 4.5, CURRENT_TIMESTAMP, '1995-02-10', false),
--       ('id2', 'Bob Jones', 'bob@example.com', 5000, 3.8, CURRENT_TIMESTAMP, '1988-06-23', true),
--       ('id3', 'Charlie Evans', 'charlie@example.com', 25000, 4.9, CURRENT_TIMESTAMP, '1990-09-15', false);

INSERT INTO customer (id, name, email, balance, rating, created_at, birth_date, married) VALUES
    ('id4', 'Dewi Kurnia', NULL, NULL, NULL, CURRENT_TIMESTAMP, NULL, NULL);
 
 