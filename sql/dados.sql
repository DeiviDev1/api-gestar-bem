INSERT into usuarios (nome,nick, email,senha)
values
    ('Usuario 1', 'usuario_1', 'usuario01@gmail.com', '$2a$10$3a7Wog4LW8THr3XsqviNvOWWC9kfPADCZvfm4axRQ4f3ydgTqvj9q'),
    ('Usuario 2', 'usuario_2', 'usuario02@gmail.com', '$2a$10$3a7Wog4LW8THr3XsqviNvOWWC9kfPADCZvfm4axRQ4f3ydgTqvj9q'),
    ('Usuario 3', 'usuario_3', 'usuario03@gmail.com', '$2a$10$3a7Wog4LW8THr3XsqviNvOWWC9kfPADCZvfm4axRQ4f3ydgTqvj9q'),
    ('Usuario 4', 'usuario_4', 'usuario04@gmail.com', '$2a$10$3a7Wog4LW8THr3XsqviNvOWWC9kfPADCZvfm4axRQ4f3ydgTqvj9q');

INSERT into seguidores (usuario_id, seguidor_id)
values
    (1,2),
    (3,1),
    (1,3),
    (2,1),
    (2,3),
    (2,4),
    (3,2),
    (3,4),
    (4,1),
    (4,2),
    (4,3);

INSERT into publicacao (titulo, conteudo, autor_id)
values
    ('Publicação do User 1', 'Conteúdo da publicação do User 1', 1),
    ('Publicação do User 2', 'Conteúdo da publicação do User 2', 2),
    ('Publicação do User 3', 'Conteúdo da publicação do User 3', 3),
    ('Publicaçaõ do User 4', 'Conteúdo da publicação do User 4', 4);

