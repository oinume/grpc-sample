# How to run

## Installation
```
$ make install-commands
$ make setup
```

## Generate go code from proto

```
$ make proto/go
```

## Run server

```
$ go run ./cmd/main.go
```

## Call API with curl

### Create a new user
 
```
$ curl -v -X POST -d '{"name": "oinume", "real_name": "kazuhiro oinuma"}' http://localhost:5000/v1/users | jq .

{
  "id": "12345",
  "name": "oinume",
  "real_name": "kazuhiro oinuma"
}
```

### Get a user

```
$ curl -v -X GET http://localhost:5000/v1/users/1 | jq .

{
  "id": "1",
  "name": "oinume",
  "real_name": "Kazuhiro Oinuma"
}
```

### List users

```
$ curl -v -X GET http://localhost:5000/v1/users | jq .

{
  "users": [
    {
      "id": "1",
      "name": "oinume",
      "real_name": "kazuhiro oinuma"
    },
    {
      "id": "2",
      "name": "akuwano",
      "real_name": "Akihiro Kuwano"
    },
    {
      "id": "3",
      "name": "oranie",
      "real_name": "Takashi Narita"
    }
  ]
}
```
