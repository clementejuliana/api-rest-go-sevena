
CREATE TABLE notificacaos(
    NotificacaoID serial primary key,
    Usuario INT,
    Conteudo VARCHAR(255)
);

INSERT INTO Notificacaos(NotificacaoID, Usuario, Conteudo) VALUES
(1, 'sobre terminar esse curso e esse tcc e ganhar no minino 5k inicial depois 23 seria o suficiente'),
(2, 'sobre tentar ganhar algo nessa vida caralho to na segunda graduação, e minha vida não mudou.');


