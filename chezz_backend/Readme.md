# Golang Migration
```
https://github.com/golang-migrate/migrate

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#usage
```

#   Migrate
## Changes from Go to DB and from DB to Go.
```
>   migrate -version
>   migrate --help
>   mkdir db/migration
>   migrate create -ext sql -dir models/migration -seq chezz_schema
>   migrating up : migrate -path models/migration -database "mysql://root:password@localhost:3306/chezz?sslmode=disable" -verbose up

```