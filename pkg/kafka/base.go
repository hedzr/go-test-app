// Copyright © 2020 Hedzr Yeh.

package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/go-test-app/pkg/vxconf"
	"github.com/sirupsen/logrus"
	"log"
)

type (
	appKafkaConfig struct {
		Version string
		Brokers []string
		Group   string
		Topics  []string
		Oldest  bool
		Key     string
	}

	base struct {
		kafkaConfig *appKafkaConfig
		config      *sarama.Config
	}
)

func (s *base) preEntry(keyPathPrefix string) (config *sarama.Config, err error) {
	var (
		version sarama.KafkaVersion
	)

	enabled := cmdr.GetBoolRP(keyPathPrefix, "enabled")
	if !enabled {
		return
	}

	src := cmdr.GetStringRP(keyPathPrefix, "source")
	if src != "kafka" {
		return
	}

	// keyPath := fmt.Sprintf("server.deps.mq.%v.%v", src, vxconf.RunModeExt())
	// keyPath := fmt.Sprintf("server.deps.mq.%v.prod", src)
	keyPath := fmt.Sprintf("%s.%s", keyPathPrefix, src)
	s.kafkaConfig = new(appKafkaConfig)
	if err = vxconf.LoadSectionTo(vxconf.RunMode(), keyPath, &s.kafkaConfig); err != nil {
		return
	}

	s.kafkaConfig.Brokers = cmdr.GetStringSliceR("kafka-test.addr", s.kafkaConfig.Brokers...)

	version, err = sarama.ParseKafkaVersion(s.kafkaConfig.Version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	logrus.Debugf("kafka version: %v. config: %v", s.kafkaConfig.Version, s.kafkaConfig)

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config = sarama.NewConfig()
	config.Version = version // sarama.V2_3_0_0
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	if s.kafkaConfig.Oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	return
}
