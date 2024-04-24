API baseado no curso:  
https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/learn/lecture/22117390#questions/12582678


```go
fmt.Print("\n####\n####\n####____\n####\n")
fmt.Print()


## Ou caso queira, é possivel usar esse 'Dumper': 
spew.Dump(x)
```




### GET User
```shell
curl http://localhost:3333/users
```


### POST User
```shell
curl -X POST \
  http://localhost:3333/users \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Gabriel Sousa Darezzo",
    "nick": "gabrieldarezzo",
    "email": "darezzo.gabriel@example.com",
    "password": "pass123"
}'


curl -X POST \
  http://localhost:3333/users \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Gabriel Sousa Darezzo",
    "nick": "gabriel.darezzo",
    "email": "darezzo.gabriel@gmail.com",
    "password": "pass123"
}'


curl -X POST \
  http://localhost:3333/users \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Gabriel Sousa Darezzo",
    "nick": "darezzo-date",
    "email": "darezzo@teste.com",
    "password": "pass123"
}'


curl -X POST \
  http://localhost:3333/users \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Junior",
    "nick": "junior-date",
    "email": "junior@teste.com",
    "password": "pass123"
}'
```

## Get User (Filtered by Criteria)
```shell
curl  http://localhost:3333/users/user\?darezzo
## ====
curl  http://localhost:3333/users\?user\=darezzo | jq

curl  http://localhost:3333/users\?user\=no+exists | jq
```


## Get User/{id}
```shell
curl  http://localhost:3333/users/1
## ====
curl  http://localhost:3333/users/3 | jq
```



## Get UpdateUser/{id}
```shell
curl --request PUT \
  --url http://localhost:3333/users/1 \
  --header 'Content-Type: application/json' \
  --data '{
  "name": "Atualizadinho!",
  "nick": "gabrieldarezzo",
  "email": "darezzo.gabriel@example.com",
  "created_at": "2024-04-21T01:16:27-03:00"
}'
```

## Get DeleteUser/{id}
```shell
curl --request DELETE \
  --url http://localhost:3333/users/7 \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM3NjU5MjQsInVzZXJJZCI6N30.bR_d--YqCxtY6HehGM4cQcwnSiZDpPQVtw5JXIAVNrA' \
  --header 'Content-Type: application/json'
```




## Post InvalidEmail
```shell
curl -X POST \
  http://localhost:3333/users \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Junior",
    "nick": "junior-date",
    "email": "EmailInvalido@test.com",
    "password": "pass123"
}'
```


## Post login
```shell
curl -X POST \
  http://localhost:3333/login \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "darezzo.gabriel@gmail.com",
    "password": "pass1234"
}'
```

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM3Njk2NjYsInVzZXJJZCI6M30.uofsj5V82zyxHQR37NRud9i-zfqmXejIqxhhWArIW4A
```



## GET users/user=darezzo (withToken)
```shell
curl --request GET \
  --url 'http://localhost:3333/users?user=darezzo' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM5NDExMDIsInVzZXJJZCI6M30.jclhR6Vzq80zU2BihLSyb5LA-kbqlNXmODy0UyGWnT4' \
  --header 'Content-Type: application/json'
```



## Get UpdateUser/{id} (withToken)
```shell
curl --request PUT \
  --url http://localhost:3333/users/1 \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM3NjU5MjQsInVzZXJJZCI6N30.bR_d--YqCxtY6HehGM4cQcwnSiZDpPQVtw5JXIAVNrA' \
  --data '{
  "name": "Atualizadinho!",
  "nick": "gabrieldarezzo",
  "email": "darezzo.gabriel@example.com",
  "created_at": "2024-04-21T01:16:27-03:00"
}'
```

## Get FallowUser/{id} (withToken)
```shell
curl --request GET \
  --url 'http://localhost:3333/users/1/followers' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM5NDExMDIsInVzZXJJZCI6M30.jclhR6Vzq80zU2BihLSyb5LA-kbqlNXmODy0UyGWnT4' \
  --header 'Content-Type: application/json'
```


## Get FollowingUsers/{id} (withToken)
```shell
curl --request GET \
  --url 'http://localhost:3333/users/1/following' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM5NDExMDIsInVzZXJJZCI6M30.jclhR6Vzq80zU2BihLSyb5LA-kbqlNXmODy0UyGWnT4' \
  --header 'Content-Type: application/json'
```



## Get user/{id}/UpdatePassword (withToken)
```shell
curl --request POST \
  --url http://localhost:3333/users/1/update-password \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM5NDQ3MjUsInVzZXJJZCI6MX0.p-AuCd8LmP41sasjeAUY1U3Dn6wif5Llu3V3lhm5quY' \
  --data '{
  "password": "pass12345",
  "new_password": "new_password"
}'
```




# Articles  

// id          int auto_increment primary key,
// title       varchar(100) not null,
// content     varchar(300) not null,
// likes       INT default 0,
// author_id   INT NOT NULL,
// created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
// FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE

## [POST] /articles
```shell
curl --request POST \
  --url http://localhost:3333/articles \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM5NDY3MjIsInVzZXJJZCI6MX0.o8y80wemrKD2uWD7zaX-qlZpCsqdD7GVnDyWQKNyj9s' \
  --data '{
  "title": "Construindo um framework em GO, chamado Horse (Go-HORSE)",
  "content": "TEXT____Construindo um framework em GO, chamado Horse (Go-HORSE)___TEXT"
}'
```