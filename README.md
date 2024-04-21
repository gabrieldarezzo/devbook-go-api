API baseado no curso:  
https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/  


```go
fmt.Print("\n####\n####\n####____\n####\n")
fmt.Print()


# Ou caso queira, é possivel usar esse 'Dumper': 
spew.Dump(x)
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

# Get DeleteUser/{id}
```shell
curl --request DELETE \
  --url http://localhost:3333/users/1 \
  --header 'Content-Type: application/json'
```


# Post InvalidEmail
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


# Post InvalidEmail
```shell
curl -X POST \
  http://localhost:3333/users/login \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "darezzo.gabriel@gmail.com",
    "password": "pass123"
}'
```



