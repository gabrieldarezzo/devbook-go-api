API baseado no curso:  
https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/  


```go
fmt.Print("\n####\n####\n####Like a Hurricane!\n####")
```




## GET User
```shell
curl http://localhost:3333/users
```


## POST User
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
    "nick": "gabrieldarezzo2",
    "email": "darezzo.gabriel2@example.com",
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

# Get User (Filtered)
```shell
curl  http://localhost:3333/users/user\?darezzo
# ====
curl  http://localhost:3333/users\?user\=darezzo | jq
```
