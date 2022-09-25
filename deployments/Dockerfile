# syntax=docker/dockerfile:1

## Build
FROM golang:latest as build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/example-rest-api

RUN go build -v -o server


## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/cmd/example-rest-api/server /server
COPY --from=build /app/docs/swagger.json /doc/swagger.json

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]