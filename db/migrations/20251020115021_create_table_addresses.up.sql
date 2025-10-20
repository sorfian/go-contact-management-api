CREATE TABLE addresses
(
    id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    contact_id  BIGINT       NOT NULL,
    street      VARCHAR(200) NOT NULL,
    city        VARCHAR(100) NOT NULL,
    province    VARCHAR(100) NOT NULL,
    country     VARCHAR(100) NOT NULL,
    postal_code VARCHAR(10)  NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP    NULL,
    FOREIGN KEY (contact_id) REFERENCES contacts (id) ON DELETE CASCADE ON UPDATE CASCADE,
    INDEX idx_contact_id (contact_id),
    INDEX idx_deleted_at (deleted_at)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;