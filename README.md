# lmwn-assignment-go

## How to Run
The web server will listen on port `8080`

```shell
go mod tidy
go run cmd/main.go
```
or
Open a executable file located on `./bin/...`

## How to test

#### Use `curl` commands
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

## Project Structure
```
.
├── cmd
│        └── main.go
└──  src
         ├── biz.go
         ├── biz_test.go
         ├── custom_type.go
         ├── domain.go
         ├── handler.go
         ├── handler_test.go
         ├── http.go
         ├── http_test.go
         ├── mock_reporter.go
         ├── set.go
         └── set_test.go
```

## Unit Test
```
➜  lmwn-assignment-go git:(main) ✗ go test ./...
?       lmwn-assignment-go/cmd  [no test files]
ok      lmwn-assignment-go/src  1.440s
➜  lmwn-assignment-go git:(main) ✗ go test -cover ./...
?       lmwn-assignment-go/cmd  [no test files]
ok      lmwn-assignment-go/src  1.355s  coverage: 75.5% of statements
```

## Results
see more at `./result/...`

![get_summary_filter.png](result%2Fget_summary_filter.png)

## Dependencies
```shell
go mod init lmwn-assignment-go

go get -u github.com/gin-gonic/gin

go get -u github.com/golang/mock/mockgen/model
go install github.com/golang/mock/mockgen@v1.6.0
```

## Build

```shell
GOOS=linux GOARCH=amd64 go build -o ./bin/linux cmd/main.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/mac_intel cmd/main.go
GOOS=darwin GOARCH=arm64 go build -o ./bin/mac_m1 cmd/main.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows.exe cmd/main.go
```
