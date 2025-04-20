package start_create_topics

import (
	"github.com/IBM/sarama"
	"github.com/ether-echo/migrations/pkg/logger"
)

var (
	log = logger.Logger().Named("create topics").Sugar()
)

func CreateNewTopics(cnf *sarama.Config) {

	admin, err := sarama.NewClusterAdmin([]string{"kafka:9092"}, cnf)
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
		"taro": {
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
		"numerology": {
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
			log.Infof("Failed to create topic %s: %v", topic, err)
		} else {
			log.Infof("Topic %s created successfully", topic)
		}
	}
}
