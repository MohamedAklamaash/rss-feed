## **Spin up a database container**

```
docker run -d \
  --name rss-blogs \
  -e POSTGRES_USER=aklamaash \
  -e POSTGRES_PASSWORD=akla123 \
  -e POSTGRES_DB=blogs \
  -p 5432:5432 \
  postgres
```

## PGAdmin

```
docker run -d \
  --name pgadmin \
  -e PGADMIN_DEFAULT_EMAIL=admin@local.dev \
  -e PGADMIN_DEFAULT_PASSWORD=admin123 \
  -p 5050:80 \
  dpage/pgadmin4:latest

```

## Install db tools

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

```
go get github.com/sqlc-dev/sqlc/cmd/sqlc@lates
```

## *Goose is for tracking migrations*

## *SQLC is interesting*

we can set the sqlc.yaml file in the root of the project, set folder for 
queries and schema and when we run sqlc generate we get all base boiler-plate code 