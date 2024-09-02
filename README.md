# Go Chevere

`go run main.go`

`go mod init go-chevere`

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

## Migrate and deploy prima in the database

```bash
go run github.com/steebchen/prisma-client-go migrate dev --preview-feature --create-only
```

## Database

for docker test

```bash
DATABASE_URL=postgres://developer:developer@postgres:5432/go_prisma?schema=public
```

for local test

```bash
DATABASE_URL=postgres://developer:developer@localhost:5432/go_prisma?schema=public
```
