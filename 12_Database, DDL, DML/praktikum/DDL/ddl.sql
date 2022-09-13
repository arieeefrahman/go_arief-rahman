/* 1. Create database alta_online_shop */
CREATE DATABASE alta_online_shop;

USE alta_online_shop;

/* 2. a. Create table user*/
CREATE TABLE users ( 
    user_id INT PRIMARY KEY AUTO_INCREMENT, 
    full_name VARCHAR(255) NOT NULL, 
    address VARCHAR(255) NOT NULL, 
    birthday DATE NOT NULL,
    user_status VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


/* 2. b. Create table product, product type, operators, product description, payment_method*/
CREATE TABLE product (
    product_id INT PRIMARY KEY AUTO_INCREMENT,
	product_name VARCHAR(255),
    product_description VARCHAR(255),
    product_price INT,
    operators VARCHAR(255),
    payment_method INT
);

/* 2. c. Create table transaction, transaction detail*/
CREATE TABLE transaction (
    transaction_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);

CREATE TABLE transaction_detail (
    transaction_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);

/* 3. Create tabel kurir dengan field id, name, created_at, updated_at */
CREATE TABLE kurir (
	kurirID INT PRIMARY KEY AUTO_INCREMENT,
    kurirName VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

/* 4. Tambahkan ongkos_dasar column di tabel kurir */
ALTER TABLE kurir ADD ongkos_dasar INT;

/* 5. Rename tabel kurir menjadi shipping */
ALTER TABLE kurir RENAME TO shipping;

/* 6. Hapus / drop tabel shipping karena ternyata tidak dibutuhkan */
DROP TABLE shipping;

/* 7. a. 1-to-1: payment method description */
CREATE TABLE payment_method_description (
	payment_id INT PRIMARY KEY AUTO_INCREMENT,
    payment_name VARCHAR(255),
    payment_description VARCHAR(255),
    transaction_id INT,
    FOREIGN KEY (transaction_id) REFERENCES transaction(transaction_id)
);

/* 7. b. 1-to-many: user dengan alamat */
CREATE TABLE alamat_detail (
	user_id INT,
    jalan VARCHAR(255),
    kecamatan VARCHAR(255),
    kota VARCHAR(255),
    provinsi VARCHAR(255),
    negara VARCHAR(255),
    
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

/* 7. c. many-to-many: user dengan payment method menjadi user_payment_method_detail */
CREATE TABLE user_payment_method_detail (
	user_id INT,
    payment_id INT,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (payment_id) REFERENCES payment_method_description(payment_id)
);