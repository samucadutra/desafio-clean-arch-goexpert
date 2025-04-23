# Makefile for managing database migrations using golang-migrate

# This command will create a new migration file in the sql/migrations directory.
createmigration:
	migrate create -ext=sql -dir=internal/infra/database/migrations -seq init

# This command will apply all up migrations to the database.
migrateup:
	migrate -path=internal/infra/database/migrations -database "mysql://root:root@tcp(mysql:3306)/orders" -verbose up

# This command will apply all down migrations to the database.
migratedown:
	migrate -path=internal/infra/database/migrations -database "mysql://root:root@tcp(mysql:3306)/orders" -verbose down

# PHONY is used to declare that the target is not a file.
.PHONY: migrateup migratedown createmigration