package main

import (
	"github.com/IBM/sarama"
	"github.com/ether-echo/migrations/internal/start_create_topics"
	"github.com/ether-echo/migrations/internal/start_migrations"
	"github.com/ether-echo/migrations/pkg/config"
	_ "github.com/lib/pq"
	"time"

	"github.com/ether-echo/migrations/pkg/logger"
)

var (
	log = logger.Logger().Named("main").Sugar()
)

func main() {

	cnf, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	confSarama := sarama.NewConfig()

	time.Sleep(6 * time.Second)

	// запуск миграции
	start_migrations.OpenDB(cnf)

	// запуск топиков
	start_create_topics.CreateNewTopics(confSarama)

}
