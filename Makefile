migrateup:
	migrate -path db/migration -database "postgres://myuser:mypassword@localhost/hello_go_todo?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://myuser:mypassword@localhost/hello_go_todo?sslmode=disable" -verbose down

.PHONY: migrateup migratedown