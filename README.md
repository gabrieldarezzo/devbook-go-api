# Curso GO | GoLang desenvolvendo uma aplicação do zero  

API baseado no curso:  
https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/


No curso o @OtavioGallego aborda:

* Desenvolver uma rede social utilizando a linguagem Go
* Utilizar Go para desenvolvimento front-end e back-end
* Trabalhar de forma prática e eficiente com concorrência
* Construir uma aplicação de linha de comando
* Construir uma API altamente robusta e escalável utlizando boas práticas
* Todos os fundamentos da linguagem de forma aprofundada


## Instalação
Para subir em ambiente local, é interessante ter:
- Docker https://docs.docker.com/engine/install/
- Go https://go.dev/doc/install
- make*
- jq* (Opcional), Ótima extensão para 'santizar' e melhorar a visualização.

### *make
Execute o comando abaixo:  
```shell
make -version
```  
Se o `make` estiver instalado, você verá sua versão. Caso contrário, será necessário instalá-lo antes de prosseguir.

### *go
Execute o comando abaixo:  
```shell
go version
```  
Se o `go` estiver instalado, você verá sua versão. Caso contrário, será necessário instalá-lo antes de prosseguir.


### *docker
Execute o comando abaixo:  
```shell
docker --version
```  
Se o `docker` estiver instalado, você verá sua versão. Caso contrário, será necessário instalá-lo antes de prosseguir.

## Subindo local

### Make Magic!

```shell
# Subir o docker
make up
# Rodar a aplicação em go
make run
```

Rotas criadas: 
## Get User (Filtered by Criteria)
```shell
curl  http://localhost:3333/users/user\?darezzo
curl  http://localhost:3333/users\?user\=darezzo | jq
curl  http://localhost:3333/users\?user\=no+exists | jq
```

## Get User/{id}
```shell
curl  http://localhost:3333/users/1
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

## [GET] /articles
```shell
curl --request GET \
  --url http://localhost:3333/articles/1 \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM5NDY3MjIsInVzZXJJZCI6MX0.o8y80wemrKD2uWD7zaX-qlZpCsqdD7GVnDyWQKNyj9s'
```  

## [GET] GetAll /articles
```shell
curl --request GET \
  --url http://localhost:3333/articles \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTQwNjMzMzAsInVzZXJJZCI6MX0.T5CNHGJMGWVMmhpyNql-XSIsDHI2te3-zWacXggfwn0'
```  

## [UPDATE] update a article /articles/{id}
```shell
curl --request PUT \
  --url http://localhost:3333/articles/1 \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTQwNjMzMzAsInVzZXJJZCI6MX0.T5CNHGJMGWVMmhpyNql-XSIsDHI2te3-zWacXggfwn0' \
  --data '{
  "title": "Construindo um framework em GO, chamado Horse (Go-HORSE)",
  "content": "TEXT____Construindo um framework em GO, chamado Horse (Go-HORSE)___TEXT"
}'
```  

## [DELETE] delete a article /articles/{id}
```shell
curl --request DELETE \
  --url http://localhost:3333/articles/1 \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTQwNjMzMzAsInVzZXJJZCI6MX0.T5CNHGJMGWVMmhpyNql-XSIsDHI2te3-zWacXggfwn0'
```  

## [GET] GetAll /articles
```shell
curl --request GET \
  --url http://localhost:3333/users/2/articles \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTQwNjMzMzAsInVzZXJJZCI6MX0.T5CNHGJMGWVMmhpyNql-XSIsDHI2te3-zWacXggfwn0'
```  

## [POST] GetAll /articles
```shell
curl --request POST \
  --url http://localhost:3333/increase-like-articles/1 \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTQwNjMzMzAsInVzZXJJZCI6MX0.T5CNHGJMGWVMmhpyNql-XSIsDHI2te3-zWacXggfwn0'
```

## Cada tem o console.log | var_dump que merece....
```go
spew.Dump(x) // BUG
spew.Dump(x) // BUG
```
