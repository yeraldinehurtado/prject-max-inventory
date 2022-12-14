create database max_inventory;
use max_inventory;

create table USERS (
    id int not null auto_increment,
    email varchar(255) not null,
    name varchar(255) not null,
    password varchar(255) not null,
    primary key (id)
);

create table PRODUCTS (
    id int not null auto_increment,
    name varchar(255) not null,
    description varchar(255) not null,
    price float not null,
    created_by int not null,
    primary key (id),
    foreign key (created_by) references USERS(id)
);

create table ROLES (
    id int not null auto_increment,
    name varchar(255) not null,
    primary key (id)
);

create table USER_ROLES (
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key (id),
    foreign key (user_id) references USERS(id),
    foreign key (role_id) references ROLES(id)
);

INSERT into ROLES (id, name) values (1, 'admin'); -- administradores
INSERT INTO ROLES (id, name) VALUES (2, 'seller'); -- vendedores
INSERT into ROLES (id, name) values (3, 'customer'); --usuarios comunes
-- docker run -d --name maria -p 3306:3306 --env MARIADB_ROOT_PASSWORD=rootroot mariadb:latest
-- el primer puerto es el de mi pc 
-- los dos puntos significa 'hacia'
-- el segundo puerto es el puerto del contenedor
-- con env pasamos variable de entorno
-- Como salió error por el puerto, coloqué 8080:80