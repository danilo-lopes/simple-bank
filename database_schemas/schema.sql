CREATE USER 'banking_svc_user'@'%' IDENTIFIED BY 'p@ssword';

CREATE DATABASE IF NOT EXISTS simple_bank;
GRANT ALL PRIVILEGES ON simple_bank.* TO 'banking_svc_user'@'%';
USE simple_bank;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS clients_transactions;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    cpf VARCHAR(11) NOT NULL UNIQUE
) ENGINE=INNODB;

CREATE TABLE accounts (
    id INT PRIMARY KEY,
      FOREIGN KEY(id) REFERENCES users(id) ON DELETE CASCADE,
    active BOOLEAN NOT NULL,
    available_limit INT NOT NULL
) ENGINE=INNODB;

CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    amount INT NOT NULL,
    merchant VARCHAR(50),
    time VARCHAR(50)
) ENGINE=INNODB;

CREATE TABLE clients_transactions (
    transaction_id INT NOT NULL,
      FOREIGN KEY(transaction_id) REFERENCES transactions(id) ON DELETE CASCADE,
    user_id INT NOT NULL,
      FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY(transaction_id, user_id)
) ENGINE=INNODB;

CREATE TABLE violations (
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  user_id INT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
  violation VARCHAR(100) NOT NULL
) ENGINE=INNODB;