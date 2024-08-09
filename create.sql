CREATE DATABASE cold_storage;
CREATE USER 'cold_storage_adm'@'localhost';
CREATE USER 'cold_storage_viewer'@'0.0.0.0';
GRANT ALL PRIVILEGES ON cold_storage.* TO 'cold_storage_adm'@'localhost';
GRANT SELECT, SHOW_VIEW ON cold_storage.* TO 'cold_storage_viewer'@'0.0.0.0';

CREATE TABLE status_log (
    id CHAR(36) PRIMARY KEY DEFAULT UUID(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    cold_storage_id VARCHAR(50),
    temp0 FLOAT,
    temp1 FLOAT,
    temp2 FLOAT,
    temp3 FLOAT,
    temp4 FLOAT,
    temp5 FLOAT,
    main_board_failure BOOLEAN
);
