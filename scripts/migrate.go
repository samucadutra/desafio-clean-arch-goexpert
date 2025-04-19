// File: scripts/migrate.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samucadutra/desafio-clean-arch-goexpert/configs"
	"github.com/samucadutra/desafio-clean-arch-goexpert/internal/infra/database"
	"log"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(
		configs.DBDriver,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.DBUser,
			configs.DBPassword,
			configs.DBHost,
			configs.DBPort,
			configs.DBName))

	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	migrationManager := database.NewMigrationManager(db)
	if err := migrationManager.Up(); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}
