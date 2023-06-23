package infra

import (
	"fmt"
	"github.com/typical-go/typical-rest-server/internal/app/controller/kafka"
	"os"

	"github.com/sirupsen/logrus"
	kf "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type (
	// @envconfig (prefix:"BROKER_KAFKA")
	KafkaCfg struct {
		AivenAddress string `envconfig:"ADDRS" default:":9092"`
		Address      string `envconfig:"ADDRESS" default:":9092"`
		IsAiven      string `envconfig:"IS_AIVEN" default:"false"`
	}
	KafkaSvc struct {
		KafkaCtrl *kafka.KafkaCtrl
		notify    chan error
	}
)

// @ctor
//func NewKafkaSvc(kafkaCtrl *kafka.KafkaCtrl, notify chan error) (kafkasvc *KafkaSvc) {
//	kafkasvc = &KafkaSvc{}
//	//return &KafkaSvc{KafkaCtrl: kafkaCtrl, notify: notify}
//	return
//}

func InitKafkaConfig(cfgApp *App, cfgKafka *KafkaCfg) *kf.ConfigMap {
	host, err := os.Hostname()
	if err != nil {
		host = "?"
	}

	brokerAddress := cfgKafka.Address
	if cfgKafka.IsAiven == "true" {
		brokerAddress = cfgKafka.AivenAddress
	}

	return &kf.ConfigMap{
		"bootstrap.servers":             brokerAddress,
		"group.id":                      fmt.Sprintf("%s-%s", cfgApp.Key, cfgApp.Env),
		"client.id":                     host,
		"max.poll.interval.ms":          10 * 1000 * 60,
		"partition.assignment.strategy": "roundrobin",
	}
}

// @ctor
func NewProducer(cfgApp *App, cfgKafka *KafkaCfg) *kf.Producer {
	producer, err := kf.NewProducer(InitKafkaConfig(cfgApp, cfgKafka))
	if err != nil {
		logrus.Fatalf("kafka producer: %s", err.Error())
	}
	logrus.Info("Connected to kafka producer")

	return producer
}

// @ctor
func NewConsumer(cfgApp *App, cfgKafka *KafkaCfg) *kf.Consumer {
	kafkaConsumer, err := kf.NewConsumer(InitKafkaConfig(cfgApp, cfgKafka))
	if err != nil {
		logrus.Fatalf("kafka consumer: %s", err.Error())
	}

	logrus.Info("Connected to kafka consumer")
	return kafkaConsumer
}
