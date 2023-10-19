-- init.sql

-- Criação da tabela

CREATE TABLE tipo_usuario
(
    id serial PRIMARY KEY,
    nome VARCHAR(255) NOT NULL
);

INSERT INTO tipo_usuario
    (nome)
VALUES
    ('Professor'),
    ('Aluno'),
    ('Funcionário');

CREATE TABLE estado
(
    id serial PRIMARY key,
    status VARCHAR(5) not NULL,
    nome VARCHAR(255) NOT NULL
);


INSERT INTO estado
    (status, nome)
VALUES
    ('ativo', 'São Paulo'),
    ('ativo', 'Rio de Janeiro'),
    ('ativo', 'Minas Gerais');


CREATE TABLE cidade
(
    id serial PRIMARY key,
    status VARCHAR(5) not NULL,
    nome VARCHAR(255) NOT NULL,
    estado_id int,
    FOREIGN KEY (estado_id) REFERENCES estado (id)
);

--Inserindo dados na tabela cidade
INSERT INTO cidade
    (status, nome, estado_id)
VALUES
    ('ativo', 'São Paulo', 1),
    ('ativo', 'Campinas', 1),
    ('ativo', 'Ribeirão Preto', 1),
    ('ativo', 'Rio de Janeiro', 2),
    ('ativo', 'Niterói', 2),
    ('ativo', 'Belo Horizonte', 3);


CREATE TABLE instituicao
(
    id serial PRIMARY KEY,
    status VARCHAR(5) NOT NULL,
    nome VARCHAR(255) NOT NULL,
    sigla VARCHAR(255) NOT NULL,
    cnpj VARCHAR(14) NOT NULL,
    endereco VARCHAR(255) NOT NULL,
    telefone VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    cidade_id int NOT NULL,
    FOREIGN KEY (cidade_id) REFERENCES cidade (id)
);



CREATE TABLE usuario
(
    id serial PRIMARY KEY,
    status VARCHAR(5) NOT NULL,
    nome VARCHAR(255) NOT NULL,
    cpf VARCHAR(15) NOT NULL,
    rg VARCHAR(15) NOT NULL,
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

-- INSERT INTO usuario
--     (status, nome, cpf, rg, genero, data_nascimento, email, telefone, escolaridade, profissao, foto_perfil, tipo_usuario_id, instituicao_id, cidade_id)
-- VALUES
--     ('ativo', 'João da Silva', '123.456.789-00', '000.456.789-00', 'masculino', '1990-01-01', 'joao.da.silva@email.com', '123456789', 'ensino superior', 'engenheiro', NULL, 1, 1, 1);
-- INSERT INTO usuario
--     (status, nome, cpf, rg, genero, data_nascimento, email, telefone, escolaridade, profissao, foto_perfil, tipo_usuario_id, instituicao_id, cidade_id)
-- VALUES
--     ('ativo', 'Juliana Clemente', '145.456.789-00', '456.456.789-00', 'feminino', '1995-10-01', 'ju@email.com', '12345', 'ensino superior', 'desenvolvedora', NULL, 1, 1, 1);
-- INSERT INTO usuario
--     (status, nome, cpf, rg, genero, data_nascimento, email, telefone, escolaridade, profissao, foto_perfil, tipo_usuario_id, instituicao_id, cidade_id)
-- VALUES
--     ('ativo', 'Thiago Gouvea', '178.456.789-00', '093.456.789-00', 'masculino', '1998-12-01', 'thiago@email.com', '1234', 'ensino superior', 'dev', NULL, 1, 1, 1);
-- INSERT INTO usuario
--     (status, nome, cpf, rg, genero, data_nascimento, email, telefone, escolaridade, profissao, foto_perfil, tipo_usuario_id, instituicao_id, cidade_id)
-- VALUES
--     ('ativo', 'Jaqueline Menezes', '143.456.789-00', '094.456.789-00', 'feminino', '1990-08-02', 'jaque@email.com', '1234', 'ensino superior', 'gestora', NULL, 1, 1, 1);

CREATE TABLE local
(
    id SERIAL PRIMARY KEY,
    status VARCHAR(50) NOT NULL,
    sala VARCHAR(50) NOT NULL,
    setor VARCHAR(50),
    evento_id int not null,
    atividades_id int not NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE evento
(
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL,
    nome VARCHAR(350) NOT NULL,
    descricao VARCHAR(500) NOT NULL,
    data_inicio TIMESTAMP NOT NULL,
    data_final TIMESTAMP NOT NULL,
    local_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(local_id)REFERENCES local(id)
);

CREATE TABLE inscricaoEmEvento
(
    id SERIAL PRIMARY KEY,
    status VARCHAR(255),
    data TIMESTAMP NOT NULL,
    hora TIMESTAMP NOT NULL,
    evento_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(evento_id) REFERENCES evento(id),
    FOREIGN KEY(usuario_id) REFERENCES usuario(id)
);

CREATE TABLE tipoAtividade
(
    id SERIAL PRIMARY KEY,
    status VARCHAR(255),
    tipo_da_atividade VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE atividade
(
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL,
    tipo_atividade_id INTEGER NOT NULL,
    titulo VARCHAR(255) NOT NULL,
    resumo VARCHAR(255),
    data TIMESTAMP NOT NULL,
    hora_inicio TIMESTAMP NOT NULL,
    hora_termino TIMESTAMP NOT NULL,
    valor_inscricao FLOAT4 NOT NULL,
    observacao VARCHAR(255),
    ministrante VARCHAR(255) NOT NULL,
    quantidade_vagas INTEGER NOT NULL,
    duracao FLOAT4,
    carga_horaria INTEGER,
    quantidade_inscritos INTEGER NOT NULL,
    local_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tipo_atividade_id) REFERENCES tipoAtividade(id),
    FOREIGN KEY (local_id) REFERENCES local(id)
);

CREATE TABLE controlePresenca
(
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL,
    hora_entrada TIMESTAMP NOT NULL,
    hora_saida TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE inscricaoEmAtividade
(
    id SERIAL PRIMARY KEY,
    atividade_id INTEGER NOT NULL,
    evento_id INTEGER NOT NULL,
    status VARCHAR(255),
    data TIMESTAMP NOT NULL,
    hora TIMESTAMP NOT NULL,
    controle_presenca_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (atividade_id) REFERENCES atividade (id),
    FOREIGN KEY (evento_id) REFERENCES evento (id),
    FOREIGN KEY (controle_presenca_id) REFERENCES controlePresenca(id)
);

CREATE TABLE administrador
(
    id SERIAL PRIMARY KEY,
    usuario_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuario (id)
);

CREATE TABLE auth
(
    id SERIAL PRIMARY KEY,
    usuario_id INTEGER NOT NULL,
    token VARCHAR(255) NOT NULL,
    data_expiracao TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuario (id)

);

CREATE TABLE login
(
    email VARCHAR(255) NOT NULL,
    senha VARCHAR(255) NOT NULL

);


CREATE TABLE recuperacaoSenha
(
    id SERIAL PRIMARY KEY,
    usuario_id INTEGER NOT NULL,
    token VARCHAR(255) NOT NULL,
    data_expiracao TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuario (id)
);


CREATE TABLE notificacao
(
    id SERIAL PRIMARY KEY,
    usuario_id INTEGER NOT NULL,
    notificacao_id INTEGER NOT NULL,
    conteudo VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuario (id)

);




