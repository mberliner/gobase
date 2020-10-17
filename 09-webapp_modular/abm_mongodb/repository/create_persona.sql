CREATE TABLE IF NOT EXISTS persona (
id bigint(20) NOT NULL AUTO_INCREMENT,
nombre varchar(100) NOT NULL,
apellido varchar(100) DEFAULT NULL,
fecha_nacimiento DATE DEFAULT NULL,
PRIMARY KEY (id));
