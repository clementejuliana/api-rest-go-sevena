-- init.sql

-- Criação da tabela

CREATE TABLE tipo_usuario
(
    id serial PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
    tipo_usuario VARCHAR(255) UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE estado
(
    id serial PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
    nome VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE cidade
(
    id serial PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
    nome VARCHAR(255) NOT NULL,
    estado_id int,
    FOREIGN KEY (estado_id) REFERENCES estado (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE instituicao
(
    id serial PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
    nome VARCHAR(255) NOT NULL,
    sigla VARCHAR(255) NOT NULL,
    cnpj VARCHAR(14) NOT NULL,
    endereco VARCHAR(255) NOT NULL,
    telefone VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    cidade_id int NOT NULL,
    FOREIGN KEY (cidade_id) REFERENCES cidade (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);




CREATE TABLE usuario
(
    id serial PRIMARY KEY NOT NULL ,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
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
    FOREIGN KEY (cidade_id) REFERENCES cidade (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE local
(
    id SERIAL PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('Disponivel', 'inativo')),
    sala VARCHAR(50) NOT NULL,
    setor VARCHAR(50) NOT NULL,
    DataHoraFim DATETIME,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE evento
(
    id SERIAL PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
    nome VARCHAR(350) NOT NULL,
    descricao VARCHAR(500) NOT NULL,
    data_inicio DATE NOT NULL,
    data_final DATE NOT NULL,
    DataHoraInicio DATETIME,
    DataHoraFim DATETIME,
    local_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE inscricaoEmEvento
(
    id SERIAL PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
    data DATE NOT NULL,
    hora TIME NOT NULL,
    evento_id INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(evento_id) REFERENCES evento(id),
    FOREIGN KEY(usuario_id) REFERENCES usuario(id)
);

CREATE TABLE tipoAtividade
(
    id SERIAL PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
    nome VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE atividade
(
    id SERIAL PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
    tipo_atividade_id INTEGER NOT NULL,
    titulo VARCHAR(255) NOT NULL,
    resumo VARCHAR(255),
    data DATE NOT NULL,
    hora_inicio TIME NOT NULL,
    hora_termino TIME NOT NULL,
    valor_inscricao FLOAT4,
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
    FOREIGN KEY (local_id) REFERENCES local (id)
);


CREATE TABLE controlePresenca
(
    id SERIAL PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
    hora_entrada TIME NOT NULL,
    hora_saida TIME NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE inscricaoEmAtividade
(
    id SERIAL PRIMARY KEY NOT NULL,
    atividade_id INTEGER NOT NULL,
    evento_id INTEGER NOT NULL,
    status VARCHAR(255) CHECK (status IN ('pendente', 'confirmada', 'cancelada')),
    data TIME NOT NULL,
    hora TIME NOT NULL,
    controle_presenca_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (atividade_id) REFERENCES atividade (id),
    FOREIGN KEY (evento_id) REFERENCES evento (id),
    FOREIGN KEY (controle_presenca_id) REFERENCES controlePresenca(id)
);

CREATE TABLE administrador
(
    id SERIAL PRIMARY KEY NOT NULL,
    status VARCHAR(255) CHECK (status IN ('ativo', 'inativo')),
    usuario_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuario (id)
);

CREATE TABLE auth
(
    id SERIAL PRIMARY KEY NOT NULL,
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
    senha VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

);


CREATE TABLE recuperacaoSenha
(
    id SERIAL PRIMARY KEY NOT NULL,
    usuario_id INTEGER NOT NULL,
    token VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    data_expiracao TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuario (id)
);


CREATE TABLE notificacao
(
    id SERIAL PRIMARY KEY NOT NULL,
    usuario_id INTEGER NOT NULL,
    notificacao_id INTEGER NOT NULL,
    conteudo VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuario (id)

);




