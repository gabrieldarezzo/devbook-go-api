API baseado no curso:  
https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/  


```go
fmt.Print("\n####\n####\n####____\n####\n")
fmt.Print()
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

# Get User (Filtered by Criteria)
```shell
curl  http://localhost:3333/users/user\?darezzo
# ====
curl  http://localhost:3333/users\?user\=darezzo | jq

curl  http://localhost:3333/users\?user\=no+exists | jq
```


# Get User/{id}
```shell
curl  http://localhost:3333/users/1
# ====
curl  http://localhost:3333/users/3 | jq
```



# Get UpdateUser/{id}
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

