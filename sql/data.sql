INSERT INTO users (name, nick, email, password) VALUES ('Gabriel', 'gabrieldarezzo', 'darezzo.gabriel@gmail.com', '$2a$10$f3.WzCu4V2I4JVtWGjZk6e8BD0mGNRNRnxqn1TVQMsRFhsDOnyYSW'); -- password: pass123
INSERT INTO users (name, nick, email, password) VALUES ('Roberto', 'robert0', 'robert0@gmail.com', '$2a$10$f3.WzCu4V2I4JVtWGjZk6e8BD0mGNRNRnxqn1TVQMsRFhsDOnyYSW');
INSERT INTO users (name, nick, email, password) VALUES ('Maze', 'm4z3', 'm4z3@gmail.com', '$2a$10$f3.WzCu4V2I4JVtWGjZk6e8BD0mGNRNRnxqn1TVQMsRFhsDOnyYSW');

-- m4z3@gmail.com pass123


INSERT INTO followers (user_id, follower_id) VALUES (1, 2);
INSERT INTO followers (user_id, follower_id) VALUES (1, 3);
INSERT INTO followers (user_id, follower_id) VALUES (2, 1);
-- // Maze é chatão, n segue ninguem kkk
-- INSERT INTO followers (user_id, follower_id) VALUES (3, ??);



-- select * from users ; 
-- select * from followers ; 



INSERT INTO articles (title, content, author_id) 
VALUES 
("Titulo de publicação do usuario 1", "Conteudo Conteudo Conteudo Conteudo", 1);




INSERT INTO articles (title, content, author_id) 
VALUES 
("Titulo de publicação usuario 2", "Conteudo Conteudo Conteudo Conteudo", 2);


INSERT INTO articles (title, content, author_id) 
VALUES 
("Titulo de publicação usuario 3", "Conteudo Conteudo Conteudo Conteudo", 3);
