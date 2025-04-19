package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrationManager struct {
	db *sql.DB
}

func NewMigrationManager(db *sql.DB) *MigrationManager {
	return &MigrationManager{db: db}
}

func (m *MigrationManager) Up() error {
	driver, err := mysql.WithInstance(m.db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("could not create migration driver: %w", err)
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"./migrations",
		"orders",
		driver,
	)
	if err != nil {
		return fmt.Errorf("could not create migrator: %w", err)
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not run migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}
