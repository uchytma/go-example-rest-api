# go-example-rest-api
Example of rest API written in [Golang](https://go.dev/), used for learning purposes.

API consists of two endpoints:
- **/math/greatest-common-divisor** - compute greatest common divisor of given set of numbers
- **/math/least-common-multiple** - compute least common multiple of given set of numbers

Used dependencies:
- **go-chi** (API router)
- **swaggo** (Swagger UI and OpenAPI doc generation lib)

## Commands
Run doc regeneration (openAPI)
```
swag init -d .\cmd\exaple-rest-api\ --parseDependency=true
```

Run API (hardcoded http://localhost:8080)
```
go run ./cmd/exaple-rest-api/main.go
```

Run tests
```
go test -v ./testing
```

## Api documentation
Api doc (swagger) is avaliable at `<url>/swagger/index.html`

![image](https://user-images.githubusercontent.com/28567403/192160129-f7eabf3b-f3d3-4a94-87d8-fadc5d44297f.png)
