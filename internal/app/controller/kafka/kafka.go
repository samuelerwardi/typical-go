package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/typical-go/typical-rest-server/pkg/logruskit"
	"go.uber.org/dig"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type handler func(context.Context, *kafka.Message) error

type KafkaCtrl struct {
	dig.In
	KafkaConsumer *kafka.Consumer
	Handlers      map[string]handler `optional:"true"`
}

func (ox *KafkaCtrl) KafkaRoute() {
	ox.register("topic_test2", ox.TriggerDPD)

	_ = ox.register
	if err := ox.KafkaConsumer.SubscribeTopics(ox.topics(), nil); err != nil {
		logrus.Errorf("while SubscribeTopics detail: %s", err)
		return
	}

	go func() {
		for {
			msg, err := ox.KafkaConsumer.ReadMessage(-1)
			if err != nil {
				logrus.Errorf("error occurred on consumer %s, detail: %v", os.Getenv("APP_ADDRESS"), err)
				continue
			}

			logrus.Infof("kafka message : %s", msg.String())

			ox.handleMessage(msg)
		}
	}()
}

func (ox *KafkaCtrl) register(topic string, handlr handler) (err error) {
	if ox.Handlers == nil {
		ox.Handlers = make(map[string]handler)
	}

	if _, ok := ox.Handlers[topic]; ok {
		err = fmt.Errorf("topic %s already exists", topic)
		return
	}

	ox.Handlers[topic] = handlr
	return
}

func (ox KafkaCtrl) topics() (topics []string) {
	for k := range ox.Handlers {
		topics = append(topics, k)
	}
	return topics
}

func (ox KafkaCtrl) handleMessage(msg *kafka.Message) {
	defer func() {
		if errs := recover(); errs != nil {
			logrus.Errorf("%s:%v", errs, errs)
		}
	}()

	topic := ""
	if msg.TopicPartition.Topic != nil {
		topic = *msg.TopicPartition.Topic
	}

	logrus.Infof("[kafka][handleMessage] Incomming message from topic %s", msg.TopicPartition)

	if handlr, ok := ox.Handlers[topic]; ok {
		var (
			ctx = context.Background()
		)

		host, err := os.Hostname()
		if err != nil {
			host = "?"
		}

		logruskit.PutField(&ctx, "hostname", host)

		err = handlr(ctx, msg)
		if err != nil {

			logrus.Error(err)
		}

		return
	}
}

func (ox *KafkaCtrl) TriggerDPD(ctx context.Context, msg *kafka.Message) (err error) {
	type Human struct {
		Name string `json:"name"`
	}
	obj := Human{}
	err = json.Unmarshal(msg.Value, &obj)
	if err != nil {
		return err
	}
	time.Sleep(8 * time.Second)
	logrus.Infof("[samuel-debug][end] Message on %s: %s %s \n", msg.TopicPartition, string(msg.Value),
		time.Now().Format("2006-01-02 15-04-05"))
	return
}
