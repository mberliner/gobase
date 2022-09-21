DROP TABLE IF EXISTS user;
CREATE TABLE user (
id bigint(20) NOT NULL AUTO_INCREMENT,
usuario varchar(100) NOT NULL,
nombre varchar(100) NOT NULL,
apellido varchar(100) DEFAULT NULL,
edad int(10) DEFAULT NULL,
password varchar(100) DEFAULT NULL,
UNIQUE KEY (usuario),
PRIMARY KEY (id));
