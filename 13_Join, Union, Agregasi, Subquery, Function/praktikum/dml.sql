USE alta_online_shop;

/* 1a */
INSERT INTO operators(id, name)
VALUES(1, "Ben");
INSERT INTO operators(name)
VALUES("Ash"),("Mike"),("Jasmine"),("Sharon");

/* 1b */
INSERT INTO product_types(id, name)
VALUES(1, "Electronics");
INSERT INTO product_types(name)
VALUES("Books"),("Clothes");
ALTER TABLE

/*1c*/
ALTER TABLE
    products AUTO_INCREMENT = 1;
INSERT INTO products(
    product_type_id,
    operator_id,
    code,
    name,
	status
)
VALUES(
    1,
    3,
    "PQ206/18",
    "Philips Travel Shaver",
    1
),(
    1,
    3,
    "C16VNGNC",
    "DDR4 SODIMM 2400MHz",
    1
);

/* 1d */
INSERT INTO products(product_type_id, operator_id, code, name, status) 
VALUES
(2, 1, "9780465050659", "The Design of Everyday Things", 1),
(2, 1, "0133970779", "Fundamentals of Database Systems", 1),
(2, 1, "1942788762", "The Unicorn Project", 1);

/* 1e */
INSERT INTO products(product_type_id, operator_id, code, name, status) 
VALUES
(3, 4, "MEN1TD", "Tie Dye By Bathing Ape Tee", 1,),
(3, 4, "MEN2TCCS", "Text Code Camo Side Big Ape Head Tee", 1,),
(3, 4, "MEN3MCC", "Military Crazy College Tee", 1,);

/* 1f */
ALTER TABLE
    `product_description` AUTO_INCREMENT = 1;
INSERT INTO `product_description`(`description`)
VALUES(
    'Alat cukur elektrik Philips PQ206 memadukan sistem bercukur yang bersih dengan kepala pencukur mengambang terpisah. Yakinlah, Anda akan tampil dengan gaya terbaik - setiap hari.'
),(
    'CORSAIR high performance Vengeance SODIMM memory kit, 2400MHz CL16 1.2V, memungkinkan Anda untuk secara otomatis meningkatkan kinerja konfigurasi ulang BIOS tanpa Anda. Ramping dan desain menarik untuk memastikan kompatibilitas fisik dengan semua notebook Generasi 6 Intel Core i5 dan i7. Setiap modul dibangun dengan menggunakan DRAM yang dipilih dengan cermat untuk memungkinkan stabilitas yang prima dan didukung oleh garansi seumur hidup Corsair.'
),(
    'Anyone who designs anything to be used by humans -- from physical objects to computer programs to conceptual tools -- must read this book, and it is an equally tremendous read for anyone who has to use anything created by another human. It could forever change how you experience and interact with your physical surroundings, open your eyes to the perversity of bad design and the desirability of good design, and raise your expectations about how things should be designed.'
),(
    'This book introduces the fundamental concepts necessary for designing, using, and implementing database systems and database applications. Our presentation stresses the fundamentals of database modeling and design, the languages and models provided by the database management systems, and database system implementation techniques.'
),(
    'In The Unicorn Project, we follow Maxine, a senior lead developer and architect, as she is exiled to the Phoenix Project, to the horror of her friends and colleagues, as punishment for contributing to a payroll outage. She tries to survive in what feels like a heartless and uncaring bureaucracy and to work within a system where no one can get anything done without endless committees, paperwork, and approvals.'
),(
    'Kaos Model Tie Dye aman dan tidak luntur. Bahan Kain 100% Premium Cotton-Combed 30s Ultrasoft. Bahan sejuk dan nyaman dipakai, gak gampang kusut, gak gampang berbulu dan mudah dicuci.'
),(
    'Kaos Model Polos dengan Gambar aman dan tidak luntur. Bahan Kain 100% Premium Cotton-Combed 30s Ultrasoft. Bahan sejuk dan nyaman dipakai, gak gampang kusut, gak gampang berbulu dan mudah dicuci.'
),(
    'Kaos model polos dengan stiker gambar tulisan di bagian tengah. Kaos aman dan tidak luntur. Bahan Kain 100% Premium Cotton-Combed 30s Ultrasoft. Bahan sejuk dan nyaman dipakai, gak gampang kusut, gak gampang berbulu dan mudah dicuci.'
);


/* 1g */
ALTER TABLE
    `payment_methods` AUTO_INCREMENT = 1;
INSERT INTO `payment_methods`(`name`, `status`)
VALUES('GoPay', 1),('OVO', 1),('DANA', 1);

/* 1h */
ALTER TABLE
    `users` AUTO_INCREMENT = 1;
INSERT INTO `users`(`name`, `status`, `dob`, `gender`)
VALUES('James', 1, '1995-02-12', 'M'),('Mark', 1, '1995-02-12', 'M'),('Katarina', 1, '1999-04-23', 'F'),('Gomez', 1, '1998-02-12', 'F'),('Denise', 1, '1998-05-28', 'F');

/* 1i dan 1j*/
ALTER TABLE
    transactions AUTO_INCREMENT = 1;
INSERT INTO `transactions`(
    `user_id`,
    `payment_method_id`,
    `status`
)
VALUES(1, 3, 'success'),(1, 3, 'success'),(1, 3, 'success'),(2, 1, 'success'),(2, 1, 'success'),(2, 1, 'success'),(3, 1, 'success'),(3, 1, 'success'),(3, 1, 'success'),(4, 2, 'success'),(4, 2, 'success'),(4, 2, 'success'),(5, 1, 'success'),(5, 1, 'success'),(5, 1, 'success');

INSERT INTO `transaction_details`(
    `transaction_id`,
    `product_id`,
    `status`,
    `qty`,
    `price`
)
VALUES(1, 3, 'success', 1, 30000),(1, 5, 'success', 1, 50000),(1, 4, 'success', 3, 40000),(2, 4, 'success', 2, 40000),(2, 5, 'success', 1, 50000),(2, 3, 'success', 2, 30000),(3, 3, 'success', 1, 30000),(3, 6, 'success', 1, 60000),(3, 8, 'success', 1, 80000),(4, 8, 'success', 2, 80000),(4, 7, 'success', 3, 70000),(4, 4, 'success', 2, 40000),(5, 3, 'success', 2, 30000),(5, 6, 'success', 1, 60000),(5, 7, 'success', 1, 70000),(6, 4, 'success', 2, 40000),(6, 7, 'success', 1, 70000),(6, 8, 'success', 2, 80000),(7, 5, 'success', 1, 50000),(7, 4, 'success', 1, 40000),(7, 6, 'success', 2, 60000),(8, 6, 'success', 2, 60000),(8, 7, 'success', 1, 70000),(8, 4, 'success', 1, 40000),(9, 6, 'success', 1, 60000),(9, 3, 'success', 1, 30000),(9, 4, 'success', 2, 40000),(10, 5, 'success', 3, 50000),(10, 6, 'success', 2, 60000),(10, 7, 'success', 1, 70000),(11, 4, 'success', 2, 40000),(11, 6, 'success', 2, 60000),(11, 8, 'success', 1, 80000),(12, 7, 'success', 2, 70000),(12, 6, 'success', 1, 60000),(12, 8, 'success', 1, 80000),(13, 3, 'success', 1, 30000),(13, 4, 'success', 2, 40000),(13, 5, 'success', 2, 50000),(14, 3, 'success', 3, 30000),(14, 5, 'success', 1, 50000),(14, 7, 'success', 1, 70000),(15, 3, 'success', 1, 30000),(15, 6, 'success', 2, 60000),(15, 8, 'success', 1, 80000);

UPDATE
    `transactions` T
SET
    T.total_qty =(
    SELECT
        SUM(qty)
    FROM
        `transaction_details`
    WHERE
        id = transaction_id
),
T.total_price =(
    SELECT
        SUM(qty * price)
    FROM
        `transaction_details`
    WHERE
        id = transaction_id
),
T.updated_at =(CURRENT_TIMESTAMP());

/* 2a */
SELECT `name` FROM `users` WHERE `gender` = 'M';

/* 2b */
SELECT `id`, `name` FROM `products` WHERE `id` = 3;

/* 2c */
SELECT * FROM `users` WHERE `created_at` > NOW() - INTERVAL 7 day;

/* 2d */
SELECT COUNT(*) AS TotalFemaleUser FROM `users` WHERE `gender` = 'F';

/* 2e */
SELECT * FROM `users` ORDER BY `name` ASC;

/* 2f */
SELECT * FROM `products` WHERE id < 6;

/* 3a */
UPDATE `products` SET `name` = 'product dummy' WHERE id = 1;

/* 3b */
UPDATE `transaction_details` SET `qty`= 3 WHERE `product_id` = 1;

/* 4a */
DELETE FROM `products` WHERE `id` = 1;

/* 4b */
DELETE FROM `products` WHERE `product_type_id` = 1;