-- init.sql

-- Criação da tabela
CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    email VARCHAR (100) UNIQUE NOT NULL,
    password VARCHAR (100) NOT NULL
);

-- Inserção de dados de exemplo
INSERT INTO users (username, email, password)
VALUES
    ('user1', 'user1@example.com', 'password1'),
    ('user2', 'user2@example.com', 'password2');



CREATE TABLE notificacaos(
    NotificacaoID serial primary key,
    Usuario INT,
    Conteudo VARCHAR(255)
);

INSERT INTO Notificacaos(Usuario, Conteudo) VALUES
(1, 'sobre terminar esse curso e esse tcc e ganhar no minino 5k inicial depois 23 seria o suficiente'),
(2, 'sobre tentar ganhar algo nessa vida caralho to na segunda graduação, e minha vida não mudou.');


