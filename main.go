package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func main() {
	// to produce messages
	brokers := []string{"localhost:29092"} // bootstrap.servers
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	topic := "topic name"

	dialer := &kafka.Dialer{
		DualStack: true,
		SASLMechanism: plain.Mechanism{
			Username: "username", // access key
			Password: "password", // secret
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            rootCAs,
		},
	}

	w := kafka.NewWriter(kafka.WriterConfig{
		Dialer:    dialer,
		Brokers:   brokers,
		Topic:     topic,
		Balancer:  &kafka.Hash{},
		BatchSize: 1,
	})

	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-asdasdsa"),
			Value: []byte("Hello World patrick aaa!"),
		},
	)

}
