-- init.sql

-- Criação da tabela

CREATE TABLE cidade(
id serial PRIMARY key,
status VARCHAR(5) not NULL,
nome VARCHAR(255) NOT NULL,
sigla CHAR(2) NOT NULL,
FOREIGN KEY (estado_id) REFERENCES estado (id)
);

--Inserindo dados na tabela cidade
INSERT INTO cidade (nome, sigla, estado_id)
VALUES
('São Paulo', 'SP', 1),
('Campinas', 'SP', 1),
('Ribeirão Preto', 'SP', 1),
('Rio de Janeiro', 'RJ', 2),
('Niterói', 'RJ', 2),
('Belo Horizonte', 'MG', 3);

CREATE TABLE estado(
id serial PRIMARY key,
status VARCHAR(5) not NULL,
nome VARCHAR(255) NOT NULL,
FOREIGN KEY (id) REFERENCES cidade (estado_id)
);

INSERT INTO estado (status, nome)
VALUES
('ativo', 'São Paulo', 1),
('ativo', 'Rio de Janeiro', 2),
('ativo', 'Minas Gerais', 3);


CREATE TABLE tipo_usuario (
id serial PRIMARY KEY,
nome VARCHAR(255) NOT NULL
);

INSERT INTO tipo_usuario (nome)
VALUES
('Professor'),
('Aluno'),
('Funcionário');

CREATE TABLE instituicao (
id serial PRIMARY KEY,
nome VARCHAR(255) NOT NULL,
idade_id INT NOT NULL,
FOREIGN KEY (cidade_id) REFERENCES cidade (id)
);
INSERT INTO instituicao (nome, cidade_id)
VALUES
('Universidade de São Paulo', 1),
('Universidade Federal do Rio de Janeiro', 2),
('Universidade Federal de Minas Gerais', 3);


CREATE TABLE usuario (
id serial PRIMARY KEY,
status VARCHAR(5) NOT NULL,
nome VARCHAR(255) NOT NULL,
cpf VARCHAR(11) NOT NULL,
rg VARCHAR(12) NOT NULL,
genero VARCHAR(255) NOT NULL,
data_nascimento timestamp NOT NULL,
email VARCHAR(255) NOT NULL,
telefone VARCHAR(255) NOT NULL,
escolaridade VARCHAR(255) NOT NULL,
profissao VARCHAR(255) NOT NULL,
foto_perfil VARCHAR(255) NULL,
tipo_usuario_id INT NOT NULL,
instituicao_id INT NOT NULL,
cidade_id INT NOT NULL,
FOREIGN KEY (tipo_usuario_id) REFERENCES tipo_usuario (id),
FOREIGN KEY (instituicao_id) REFERENCES instituicao (id),
FOREIGN KEY (cidade_id) REFERENCES cidade (id)
);

INSERT INTO usuario (status, nome, cpf, rg, genero, data_nascimento, email, telefone, escolaridade, profissao, foto_perfil, tipo_usuario_id, instituicao_id, cidade_id) VALUES ('ativo', 'João da Silva', '123.456.789-00', '000.456.789-00', 'masculino', '1990-01-01', 'joao.da.silva@email.com', '123456789', 'ensino superior', 'engenheiro', NULL, 1, 1, 1);
INSERT INTO usuario (status, nome, cpf, rg, genero, data_nascimento, email, telefone, escolaridade, profissao, foto_perfil, tipo_usuario_id, instituicao_id, cidade_id) VALUES ('ativo', 'Juliana Clemente', '145.456.789-00', '456.456.789-00', 'feminino', '1995-10-01', 'ju@email.com', '12345', 'ensino superior', 'desenvolvedora', NULL, 1, 1, 1);
INSERT INTO usuario (status, nome, cpf, rg, genero, data_nascimento, email, telefone, escolaridade, profissao, foto_perfil, tipo_usuario_id, instituicao_id, cidade_id) VALUES ('ativo', 'Thiago Gouvea', '178.456.789-00', '093.456.789-00', 'masculino', '1998-12-01', 'thiago@email.com', '1234', 'ensino superior', 'dev', NULL, 1, 1, 1);
INSERT INTO usuario (status, nome, cpf, rg, genero, data_nascimento, email, telefone, escolaridade, profissao, foto_perfil, tipo_usuario_id, instituicao_id, cidade_id) VALUES ('ativo', 'Jaqueline Menezes', '143.456.789-00', '094.456.789-00', 'feminino', '1990-08-02', 'jaque@email.com', '1234', 'ensino superior', 'gestora', NULL, 1, 1, 1);





