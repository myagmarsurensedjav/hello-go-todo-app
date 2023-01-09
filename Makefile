migrateup:
    migrate -path db/migration -database "mysql://root:secret@(127.0.0.1:3306)/go-todo?parseTime=true" -verbose up

migratedown:
    migrate -path db/migration -database "mysql://root:secret@(127.0.0.1:3306)/go-todo?parseTime=true" -verbose down

.PHONY: migrateup migratedown