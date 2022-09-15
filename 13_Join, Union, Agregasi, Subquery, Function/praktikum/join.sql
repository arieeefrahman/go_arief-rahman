/* 1 */
SELECT * FROM `transactions` WHERE `user_id` = 1 OR `user_id` = 2;

/* 2 */
SELECT SUM(total_price) AS JumlahHargaTransaksi FROM `transactions` WHERE `user_id` = 1;

/* 3 */
SELECT
    COUNT(*) AS TotalTransaksiType2
FROM
    (
    SELECT
        p.id AS ProductID,
        pt.id AS ProductTypesID,
        pt.name AS TYPE,
        td.transaction_id
    FROM
        products AS p
    INNER JOIN product_types AS pt
    ON
        p.product_type_id = pt.id
    INNER JOIN transaction_details AS td
    ON
        p.id = td.product_id
    WHERE
        pt.id = 2
) b;

/* 4 */
SELECT
    P.*,
    PT.name AS types
FROM
    `products` P
LEFT JOIN `product_types` PT ON
    P.product_type_id = PT.id;

/* 5 */
SELECT
    transactions.*,
    products.name,
    users.name
FROM
    `transactions`
INNER JOIN users ON users.id = transactions.user_id
INNER JOIN transaction_details ON transactions.id = transaction_details.transaction_id
INNER JOIN products ON transaction_details.product_id = products.id;

/* 6 */
DELIMITER
    //
CREATE PROCEDURE delete_transaction(IN id_inp INT(11))
BEGIN
    DELETE
FROM
    `transactions`
WHERE
    `id` = id_inp ;
DELETE
FROM
    `transaction_details`
WHERE
    `transaction_id` = id_inp ;
END //
DELIMITER
    ;

/* 7 */

DELIMITER
    //
CREATE PROCEDURE reduce_total_qty()
BEGIN
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
    T.updated_at =(CURRENT_TIMESTAMP());

END //
DELIMITER
    ;

/* 8 */
INSERT INTO products(
    product_type_id,
    operator_id,
    code,
    name,
	status
)
VALUES(
    1,
    1,
    "HPRX1",
    "Hyper X Headset Gaming",
    1
);

SELECT
    *
FROM
    `products`
WHERE
    products.id NOT IN(
    SELECT DISTINCT
        product_id
    FROM
        transaction_details
)