CREATE DATABASE IF NOT EXISTS apigestarbem;
USE apigestarbem;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null ,
    criadoEm timesTamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    primary key(usuario_id, seguidor_id)
)ENGINE=INNODB;


DROP table usuarios;
/*select * from usuarios;

DESC usuarios;
DESC seguidores;

DROP DATABASE apigestarbem;


DELETE FROM seguidores WHERE usuario_id = 3;;