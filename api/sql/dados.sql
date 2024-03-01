insert into usuarios (nome, nick, email, senha)
values
("Andre da Silva", "andre", "andre_silva@gmail.com", "$2a$10$YBFyp8fTosdIlTYHJ5GAiehKPaNo7Noio9D8KmqB52L9.fS.hSB3C"),
("Bernardo da Silva", "bern123", "bernardo_silva@gmail.com", "$2a$10$YBFyp8fTosdIlTYHJ5GAiehKPaNo7Noio9D8KmqB52L9.fS.hSB3C"),
("Caio da Silva", "caio", "caio_silva@gmail.com", "$2a$10$YBFyp8fTosdIlTYHJ5GAiehKPaNo7Noio9D8KmqB52L9.fS.hSB3C");


insert into seguidores(usuario_id, seguidor_id)
values
(1,2),
(3,2),
(1,3);