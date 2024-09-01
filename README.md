# Go Prisma Example

`go run main.go`

`go mod init go-prisma-example`

## libs

```bash
go get github.com/steebchen/prisma-client-go
go get github.com/rs/zerolog/log
go get github.com/julienschmidt/httprouter
```

## Generate the prisma lib

We do this step every time we edit the schema

```bash
go run github.com/steebchen/prisma-client-go generate dev
```
