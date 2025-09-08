CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL UNIQUE,
                                     name TEXT,
                                     email TEXT UNIQUE,
                                     age INT,
                                     is_active BOOLEAN,
                                     created_at TIMESTAMP,
                                     PRIMARY KEY (id)
    );


CREATE TABLE orders (
                        id SERIAL UNIQUE,
                        user_id INT,
                        product TEXT,
                        amount FLOAT,
                        created_at TIMESTAMP,
                        PRIMARY KEY(id),
                        FOREIGN KEY (user_id)REFERENCES users(id)
);


SELECT
    users.name,
    orders.amount
FROM
    users
        INNER JOIN orders ON users.id = orders.user_id;