// Copyright © 2019 Tricent Technology.

package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"time"
)

type (
	Msg struct {
		Data string
	}

	Sender interface {
		Send(msg string)
		SendMsg(msg *Msg)
		SendDirectly(msg *sarama.ProducerMessage)
		Stop()
	}
)

// var (
// 	akSender *appKafkaSender
// )

// keyPathPrefix = "server.deps.mq"
func NewSender(keyPathPrefix string) (sender Sender) {
	// if akSender != nil {
	// 	return
	// }

	akSender := &appKafkaSender{
		exitCh: make(chan struct{}),
		inData: make(chan *Msg),
	}

	var err error
	if akSender.config, err = akSender.preEntry(keyPathPrefix); err != nil {
		logrus.Error(err)
		return
	}

	go akSender.run()

	sender = akSender
	return
}

// func (s *appKafkaSender) Send(msg *Msg) {
// 	s.Send(msg)
// }
//
// func SendDirectly(msg *sarama.ProducerMessage) {
// 	if akSender != nil {
// 		akSender.SendDirectly(msg)
// 	}
// }

// func Stop() {
// 	if akSender != nil {
// 		akSender.Stop()
// 	}
// }

type appKafkaSender struct {
	base
	exitCh      chan struct{}
	inKafkaData chan *sarama.ProducerMessage
	inData      chan *Msg
	producer    sarama.SyncProducer
}

func (s *appKafkaSender) run() {
	defer func() {
		// akSender = nil
	}()

	for {
		select {
		case msg := <-s.inData:
			s.send(msg)
		case msg := <-s.inKafkaData:
			s.sendDirectly(msg)
		case <-s.exitCh:
			return
		}
	}
}

func (s *appKafkaSender) send(msg *Msg) {
	mmsg := &sarama.ProducerMessage{
		Topic:     s.kafkaConfig.Topics[0],
		Partition: int32(-1),
		Key:       sarama.StringEncoder(s.kafkaConfig.Key),
	}

	mmsg.Value = sarama.ByteEncoder(msg.Data)
	s.sendDirectly(mmsg)
}

func (s *appKafkaSender) sendDirectly(msg *sarama.ProducerMessage) {
	if s.producer == nil {
		var err error
		s.producer, err = sarama.NewSyncProducer(s.kafkaConfig.Brokers, s.config)
		if err != nil {
			logrus.Errorf("[kafka] NewSyncProducer failed. %v", err)
			return
		}
		// defer s.producer.Close()
	}

RETRY_SEND:
	partition, offset, err := s.producer.SendMessage(msg)
	if err != nil {
		logrus.Warnf("[kafka] Send Message Fail: %v", err)
		time.Sleep(100 * time.Millisecond)
		goto RETRY_SEND
	}

	logrus.Tracef("Partion = %d, offset = %d\n", partition, offset)
}

func (s *appKafkaSender) Send(msg string) {
	s.inData <- &Msg{Data: msg}
}

func (s *appKafkaSender) SendMsg(msg *Msg) {
	s.inData <- msg
}

func (s *appKafkaSender) SendDirectly(msg *sarama.ProducerMessage) {
	s.inKafkaData <- msg
}

func (s *appKafkaSender) Stop() {
	s.exitCh <- struct{}{}
}
