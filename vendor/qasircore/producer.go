package qasircore

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

type ProducerKafka struct {
	env       Environtment
	hostKafka string
	portKafka string
	topic     string
	data      string
	producer  sarama.AsyncProducer
}

func (pk *ProducerKafka) SetTopic(topic string) {
	pk.topic = topic
}

func (pk *ProducerKafka) SetData(data string) {
	pk.data = data
}

func (pk *ProducerKafka) setupDefaultConnection() {
	producer, err := sarama.NewAsyncProducer([]string{pk.hostKafka + ":" + pk.portKafka}, nil)

	if err != nil {
		panic(err)
	}

	pk.producer = producer
}

func (pk *ProducerKafka) Send() {
	// defer func() {
	// 	if err := pk.producer.Close(); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, errors int

	select {
	case pk.producer.Input() <- &sarama.ProducerMessage{Topic: pk.topic, Key: nil, Value: sarama.StringEncoder(pk.data)}:
		enqueued++
	case err := <-pk.producer.Errors():
		log.Println("Failed to produce message", err)
		errors++
	case <-signals:
		break
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}

func NewProducerKafka(env Environtment) *ProducerKafka {
	var producerKafka ProducerKafka

	producerKafka.hostKafka = env.Get("KAFKA_HOST")
	producerKafka.portKafka = env.Get("KAFKA_PORT")
	producerKafka.env = env

	producerKafka.setupDefaultConnection()

	return &producerKafka
}
