package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/vaibhavchalse99/config"
)

func CreateMigration(fileName string) error {
	if len(fileName) == 0 {
		return errors.New("FileName is not provided")
	}

	timestamp := time.Now().Unix()

	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", config.Migrationpath(), timestamp, fileName)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", config.Migrationpath(), timestamp, fileName)
	if err := createFile(upMigrationFilePath); err != nil {
		return err
	}

	fmt.Printf("created %s\n", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	fmt.Printf("created %s\n", downMigrationFilePath)
	return nil
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = f.Close()
	return err
}

func RunMigration() error {
	dbConfig := config.Database()
	db, err := sql.Open(dbConfig.Driver(), dbConfig.ConnectionUrl())
	if err != nil {
		return err
	}
	driverInstance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(getMigrationFilePath(), dbConfig.Driver(), driverInstance)
	if err != nil {
		return err
	}
	err = m.Up()
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}
	return err
}

func RollbackMigrations(s string) error {
	steps, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	m, err := migrate.New(getMigrationFilePath(), config.Database().ConnectionUrl())
	if err != nil {
		return err
	}

	err = m.Steps(-1 * steps)
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}

	return err
}

func getMigrationFilePath() string {
	return fmt.Sprintf("file://%s", config.Migrationpath())
}
