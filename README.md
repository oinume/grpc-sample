# How to run

### Installation
```
$ make install-commands
$ make setup
```

### Generate go code from proto

```
$ make proto/go
```

### Run server

```
$ go run ./cmd/main.go
```

### Call API with curl

```
$ curl -v -X POST -d '{"name": "oinume", "real_name": "kazuhiro oinuma"}' http://localhost:5000/v1/users | jq .

...

{
  "id": "12345",
  "name": "hoge",
  "real_name": "kazuhiro oinuma"
}
```
