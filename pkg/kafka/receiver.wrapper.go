// Copyright © 2019 Tricent Technology.

package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"net"
	"sync"
	"time"
)

type (
	Receiver struct {
		// Topic    string
		base
		consumer sarama.Consumer
		wg       sync.WaitGroup
		fn       func(msg *sarama.ConsumerMessage)
	}
)

func NewReceiver(keyPathPrefix string, fn func(msg *sarama.ConsumerMessage)) (receiver *Receiver) {
	receiver = &Receiver{
		// Topic: topic,
		fn: fn,
	}

	var err error
	if receiver.config, err = receiver.preEntry(keyPathPrefix); err != nil {
		logrus.Error(err)
		return
	}

	receiver.consumer, err = sarama.NewConsumer(receiver.kafkaConfig.Brokers, nil)
	if err != nil {
		logrus.Errorf("NewConsumer failed: %v | receiver.kafkaConfig.Brokers = %v", err, receiver.kafkaConfig.Brokers)
		if len(receiver.kafkaConfig.Brokers) > 0 {
			addrs, _ := net.LookupIP(receiver.kafkaConfig.Brokers[0])
			for _, addr := range addrs {
				if ipv4 := addr.To4(); ipv4 != nil {
					fmt.Println("  IPv4: ", ipv4)
				}
			}
		}
	}

	// go receiver.run()

	return
}

func (s *Receiver) Listen(topic string, doneCh, exitCh chan struct{}) {
	go s.run(topic, doneCh, exitCh)
}

func (s *Receiver) run(topic string, doneCh, exitCh chan struct{}) {
	if s.consumer == nil {
		logrus.Warn("kafka consumer not available")
		return
	}

	partitionList, err := s.consumer.Partitions(topic)
	if err != nil {
		logrus.Errorf("Partitions failed: %v", err)
	}

	for partition := range partitionList {
		logrus.Debugf("    partition %v", partition)
	RETRY_CONSUME_PARTITION:
		pc, err := s.consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			logrus.Warnf("consumer.ConsumePartition error: %v", err)
			time.Sleep(100 * time.Millisecond)
			goto RETRY_CONSUME_PARTITION
		}

		defer pc.AsyncClose()

		s.wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer s.wg.Done()
			for msg := range pc.Messages() {
				if s.fn != nil {
					s.fn(msg)
				} else {
					fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				}
			}
			select {
			case <-doneCh:
				exitCh <- struct{}{}
				return
			}
		}(pc)

		s.wg.Wait()
	}

}

func (s *Receiver) Stop() {
	_ = s.Close()
}

func (s *Receiver) Close() (err error) {
	if s.consumer != nil {
		if err = s.consumer.Close(); err != nil {
			logrus.Errorf("[kafka][receiver] Closing consumer failed: %v", err)
		}
		s.consumer = nil
	}
	return
}
