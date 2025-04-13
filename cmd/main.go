package main

import (
	"github.com/IBM/sarama"
	"log"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)
	config := sarama.NewConfig()
	admin, err := sarama.NewClusterAdmin([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatal("Failed to create admin:", err)
	}
	defer admin.Close()

	// Определяем список топиков
	topics := map[string]*sarama.TopicDetail{
		"start": {
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
		"support": {
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
		"message": {
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	// Создаём каждый топик
	for topic, details := range topics {
		err = admin.CreateTopic(topic, details, false)
		if err != nil {
			log.Printf("Failed to create topic %s: %v", topic, err)
		} else {
			log.Printf("Topic %s created successfully", topic)
		}
	}
}
