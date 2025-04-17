package start_migrations

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"github.com/ether-echo/migrations/pkg/config"
	"github.com/ether-echo/migrations/pkg/logger"
	"github.com/pressly/goose/v3"
)

var (
	log = logger.Logger().Named("database and goose").Sugar()
)

func OpenDB(cnf *config.Config) {

	db, err := sql.Open(cnf.DBName, fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", cnf.DBUser, cnf.DBPassword, cnf.DBName, cnf.DBHost))
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	err = goose.Up(db, "./goose_migrations")
	if err != nil {
		log.Fatalf("ошибка миграции %v", err)
	}
}
