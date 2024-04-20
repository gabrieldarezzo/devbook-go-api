CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;
CREATE TABLE users(
    id          int auto_increment primary key,
    name        varchar(250) not null,    
    nick        varchar(250) not null unique,
    email       varchar(250) not null unique,
    password    varchar(250) not null,    
    CreatedAt   timestamp default current_timestamp()       
)ENGINE=InnoDb;