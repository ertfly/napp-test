CREATE TABLE test.configs (
	id varchar(30) NOT NULL,
	`value` LONG VARCHAR NULL,
	`description` varchar(100) NOT NULL,
	CONSTRAINT configs_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;