CREATE TABLE IF NOT EXISTS test.products (
	id bigint auto_increment NOT NULL,
	sku varchar(50) NOT NULL,
	`name` varchar(250) NOT NULL,
	price_unit decimal(18,2) NOT NULL,
	price_final decimal(18,2) NOT NULL,
	last_stock_id bigint NULL,
	created_at datetime NOT NULL,
	updated_at datetime NULL,
	trash int(1) DEFAULT 0 NOT NULL,
	CONSTRAINT products_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;