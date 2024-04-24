CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS articles;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id          int auto_increment primary key,
    name        varchar(250) not null,    
    nick        varchar(250) not null unique,
    email       varchar(250) not null unique,
    password    varchar(250) not null,    
    created_at   timestamp default current_timestamp()       
)ENGINE=InnoDb;



CREATE TABLE followers(
    user_id INT NOT NULL,
    follower_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY(user_id, follower_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE articles(
    id          int auto_increment primary key,
    title       varchar(100) not null,   
    content     varchar(300) not null,   
    likes       INT default 0,
    author_id   INT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;
