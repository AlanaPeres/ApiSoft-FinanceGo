CREATE DATABASE IF NOT EXISTS softFinance; 
USE softFinance;
DROP TABLE IF EXISTS usuarios, ContaBancaria, depositos;

CREATE TABLE usuarios(
    cpf varchar(14) primary key,
    nome varchar(50) not null,
    email varchar(50) not null unique,
    senha varchar(100) not null
)ENGINE=INNODB;  

CREATE TABLE ContaBancaria (
    Agencia INT,
    ContaBancariaId INT AUTO_INCREMENT PRIMARY KEY,
    Nome VARCHAR(255),
    Cpf VARCHAR(11),
    Saldo DECIMAL(10, 2),
    FOREIGN KEY (Cpf) REFERENCES usuarios(cpf)
)ENGINE=INNODB;


CREATE TABLE Transacoes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tipo VARCHAR(20) NOT NULL CHECK (tipo IN ('deposito', 'transferencia')),
    valor DECIMAL(10, 2) NOT NULL,
    data_hora TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    conta_origem_id INT,
    conta_destino_id INT,
    FOREIGN KEY (conta_origem_id) REFERENCES ContaBancaria(ContaBancariaId),
    FOREIGN KEY (conta_destino_id) REFERENCES ContaBancaria(ContaBancariaId)
) ENGINE=InnoDB;



