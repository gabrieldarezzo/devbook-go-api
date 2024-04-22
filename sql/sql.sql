CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;
CREATE TABLE users(
    id          int auto_increment primary key,
    name        varchar(250) not null,    
    nick        varchar(250) not null unique,
    email       varchar(250) not null unique,
    password    varchar(250) not null,    
    created_at   timestamp default current_timestamp()       
)ENGINE=InnoDb;


DROP TABLE IF EXISTS followers;
CREATE TABLE followers(
    user_id INT NOT NULL,
    follower_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY(user_id, follower_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;
