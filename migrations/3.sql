CREATE TABLE IF NOT EXISTS test.stock (
	id bigint auto_increment NOT NULL,
	product_id bigint NOT NULL,
	stock_total decimal(18,4) DEFAULT 0 NOT NULL,
	stock_cut decimal(18,4) DEFAULT 0 NOT NULL,
	stock_available decimal(18,4) DEFAULT 0 NOT NULL,
	created_at datetime NOT NULL,
	CONSTRAINT stock_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;