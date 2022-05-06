___Project GO App test___

Run app
```go run ./cmd/web```


Create user
```curl -iL -X POST http://127.0.0.1:4000/user/create```

Get user by ID
```curl -iL -X GET http://127.0.0.1:4000//user?id=1```

Get users list
```curl -iL -X GET http://127.0.0.1:4000//users```
