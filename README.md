# lmwn-assignment-go

## How to Run
```shell
go mod tidy
```

```shell
go run cmd/main.go
```

## How to test

#### Use Curl commands
```shell
# health check
curl http://localhost:8080
# get summary of all provinces
curl http://localhost:8080/summary
# get summary of Nakhon Ratchasima and Nong Khai
curl http://localhost:8080/summary?filter=Nakhon+Ratchasima,Nong+Khai
```

#### Use HTTP file
open `test.http`, send a HTTP request using Goland or VS Code Rest Client

## Dependencies
```shell
go mod init lmwn-assignment-go

go get -u github.com/gin-gonic/gin

go get -u github.com/golang/mock/mockgen/model
go install github.com/golang/mock/mockgen@v1.6.0
```
