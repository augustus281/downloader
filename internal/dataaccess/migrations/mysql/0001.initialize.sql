CREATE TABLE IF NOT EXISTS accounts (
    account_id BIGINT UNSIGNED PRIMARY KEY,
    accountname VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS account_passwords (
    of_account_id BIGINT UNSIGNED PRIMARY KEY,
    hash VARCHAR(128) NOT NULL,
    FOREIGN KEY (of_account_id) REFERENCES accounts(account_id)
);

CREATE TABLE IF NOT EXISTS download_tasks (
    of_account_id BIGINT UNSIGNED PRIMARY KEY,
    download_type SMALLINT NOT NULL,
    url TEXT NOT NULL,
    download_status SMALLINT ENUM(PENDING, DOWNLOADING, FAILED, SUCCESSFULLY),
    metadata TEXT NOT NULL,
    FOREIGN KEY (of_account_id) REFERENCES accounts(account_id)
)