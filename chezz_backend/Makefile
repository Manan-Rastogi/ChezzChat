migrateup:
	migrate -path models/migration -database "mysql://<your user>:<your password>@tcp(localhost:3306)/chezz" -verbose up

rmlog:
	rm chezz.log

test:
	go test -v -cover ./...

.PHONY: migrateup rmlog test