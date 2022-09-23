# go-example-rest-api
Example of rest API written in [Golang](https://go.dev/), used for learning purposes.
## Commands
Run doc regeneration (openAPI)
```
swag init -d .\cmd\exaple-rest-api\ --parseDependency=true
```
Run API (hardcoded http://localhost:8080)
```
go run ./cmd/exaple-rest-api/main.go
```
