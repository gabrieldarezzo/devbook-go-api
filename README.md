API baseado no curso:  
https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/learn/lecture/22117390#questions/12582678


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
  --url http://localhost:3333/users/7 \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM3NjU5MjQsInVzZXJJZCI6N30.bR_d--YqCxtY6HehGM4cQcwnSiZDpPQVtw5JXIAVNrA' \
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


# Post login
```shell
curl -X POST \
  http://localhost:3333/login \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "m4z3@gmail.com",
    "password": "pass123"
}'
```

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM3Njk2NjYsInVzZXJJZCI6M30.uofsj5V82zyxHQR37NRud9i-zfqmXejIqxhhWArIW4A
```



# GET users/user=darezzo (withToken)
```shell
curl --request GET \
  --url 'http://localhost:3333/users?user=darezzo' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM3NjU5MjQsInVzZXJJZCI6N30.bR_d--YqCxtY6HehGM4cQcwnSiZDpPQVtw5JXIAVNrA' \
  --header 'Content-Type: application/json'
```



# Get UpdateUser/{id} (withToken)
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

# Get FallowUser/{id} (withToken)
```shell
curl --request GET \
  --url 'http://localhost:3333/users/1/followers' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM4OTI3NTksInVzZXJJZCI6M30.VEAb-Fvwz9sTPXCnVTmAL_fFWgNN2tFlsQm9hqKX3qc' \
  --header 'Content-Type: application/json'
```


# Get FollowingUsers/{id} (withToken)
```shell
curl --request GET \
  --url 'http://localhost:3333/users/1/following' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTM4OTI3NTksInVzZXJJZCI6M30.VEAb-Fvwz9sTPXCnVTmAL_fFWgNN2tFlsQm9hqKX3qc' \
  --header 'Content-Type: application/json'
```


