CREATE TABLE user_cred (id SERIAL PRIMARY KEY NOT NULL,
						username VARCHAR(50) NOT NULL,
						password VARCHAR(50) NOT NULL,
					   created_at DATE NOT NULL,
					   updated_at DATE NOT NULL);
ALTER TABLE user_cred ADD created_at DATE NOT NULL;
ALTER TABLE user_cred ADD updated_at DATE NOT NULL;
CREATE TABLE customer (id SERIAL PRIMARY KEY NOT NULL,
					  user_id INT NOT NULL,
					  nik VARCHAR NOT NULL,
					  name VARCHAR (50) NOT NULL,
					  email VARCHAR (50) NOT NULL,
					  phone VARCHAR (20) NOT NULL,
					  address VARCHAR (255) NOT NULL,
					   status VARCHAR NOT NULL,
					  birthdate DATE NOT NULL,
					  FOREIGN KEY (user_id) REFERENCES user_cred(id)
					  );
CREATE TABLE payment (id SERIAL PRIMARY KEY NOT NULL,
						 customer_id INT NOT NULL,
						 paid MONEY NOT NULL,
					 created_by VARCHAR NOT NULL,
					 created_at TIMESTAMP)
SELECT * FROM customer;
SELECT id, username, password FROM user_cred;
SELECT id, user_id,nik, name, email, phone, address, birthdate, status FROM customer;
SELECT id, customer_id, paid, created_by, created_at FROM payment;
