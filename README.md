# How to run

## Requirements

- Go 1.11 or later
- protoc (protobuf)

## Installation
```
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

You can specify listening port like this
```
$ PORT=5000 GRPC_PORT=5001 go run ./cmd/main.go
```

## Call API with curl

### Create a new user
 
```
$ curl -X POST -d '{"name": "oinume", "real_name": "kazuhiro oinuma"}' http://localhost:5000/v1/users | jq .

{
  "id": "12345",
  "name": "oinume",
  "real_name": "kazuhiro oinuma"
}
```

### Get a user

```
$ curl -X GET http://localhost:5000/v1/users/1 | jq .

{
  "id": "1",
  "name": "oinume",
  "real_name": "Kazuhiro Oinuma"
}
```

### List users

```
$ curl -X GET http://localhost:5000/v1/users | jq .

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

### Delete a user

```
curl -X DELETE http://localhost:5000/v1/users/12345 | jq .

{}
```