package producer

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Shopify/sarama"

	"github.com/rs/zerolog/log"
)

type ProgressEventType = int

const (
	Created ProgressEventType = iota
	Updated
	Removed
)

type ProgressEvent struct {
	Type ProgressEventType
	Body map[string]interface{}
}

type LogProducer interface {
	Send(evType ProgressEventType, req, res interface{}, err error) error
}

type logProducer struct {
	kafkaTopic   string
	syncProducer sarama.SyncProducer

	messagesCh chan *sarama.ProducerMessage
}

func New(ctx context.Context, broker []string, kafkaTopic string, capacity int) (LogProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(broker /*[]string{broker}*/, config)
	if err != nil {
		return nil, err
	}

	lp := &logProducer{
		kafkaTopic:   kafkaTopic,
		syncProducer: producer,
		messagesCh:   make(chan *sarama.ProducerMessage, capacity)}

	go lp.sendMessages(ctx)

	return lp, nil
}

func (lp *logProducer) Send(evType ProgressEventType, req, res interface{}, err error) error {

	if lp == nil {
		return errors.New("LogProducer is nil")
	}

	message := ProgressEvent{
		Type: evType,
		Body: map[string]interface{}{
			"request":   req,
			"response":  res,
			"timestamp": time.Now(),
		},
	}

	if err != nil {
		message.Body["error"] = err.Error()
	}

	b, err := json.Marshal(message)
	if err != nil {
		return err
	}

	lp.messagesCh <- &sarama.ProducerMessage{
		Topic:     lp.kafkaTopic,
		Partition: -1,
		Value:     sarama.StringEncoder(b),
	}

	return nil
}

func (lp *logProducer) sendMessages(ctx context.Context) {

	for {
		select {

		case msg := <-lp.messagesCh:
			_, _, err := lp.syncProducer.SendMessage(msg)
			if err != nil {
				log.Error().Err(err).Msg("LogProducer: failed to send message to kafka")
			}

		case <-ctx.Done():
			close(lp.messagesCh)
			lp.syncProducer.Close()
			return
		}
	}
}
