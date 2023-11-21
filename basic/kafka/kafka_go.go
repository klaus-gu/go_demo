package main

func (server *AgentServer) kafkaListener() {
	if !server.config.IsTraffic() {
		return
	}

	log.Debug("[Kafka] start listening.")
	addr := []string{server.config.Kafka.Addr}

	groupPcap := "group-pcap"
	groupUpgrade := "group-upgrade-" + server.config.Ip
	groupLog := "group-log-" + server.config.Ip
	groupConfig := "group-config-" + server.config.Ip
	groupEngine := "group-engine-" + server.config.Ip
	groupHttps := "group-https-" + server.config.Ip

	if server.config.IsTraffic() {
		go server.clusterConsumer(addr, []string{common.PcapReplayTopic}, groupPcap)
		go server.clusterConsumer(addr, []string{common.UpgradeTopic}, groupUpgrade)
		go server.clusterConsumer(addr, []string{common.DeviceCfgTopic}, groupConfig)
		go server.clusterConsumer(addr, []string{common.EngineCfgTopic}, groupEngine)
		go server.clusterConsumer(addr, []string{common.HttpsConfig}, groupHttps)
	}

	if server.config.IsSandbox() {
		go server.clusterConsumer(addr, []string{common.UpgradeTopic}, groupUpgrade)
		go server.clusterConsumer(addr, []string{common.LogTopic}, groupLog)
	}
}

func (server *AgentServer) clusterConsumer(brokers []string, topics []string, groupId string) {
	config := cluster.NewConfig()
	//config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumer, err := cluster.NewConsumer(brokers, groupId, topics, config)

	if err != nil {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
				"error":   err,
			}).Error("[Kafka Consumer] new failed.")
		return
	}

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				log.WithFields(log.Fields{
					"groupId":   groupId,
					"ka":        msg.Topic,
					"partition": msg.Partition,
					"offset":    msg.Offset,
					"key":       msg.Key,
				}).Info("[Kafka Consumer] receive message.")

				switch msg.Topic {
				case common.PcapReplayTopic:
					server.pcapreplayHandler(msg.Value)
				case common.UpgradeTopic:
					server.deviceUpgradeHandler(msg.Value)
				case common.LogTopic:
					server.logUploadHandler(msg.Value)
				case common.DeviceCfgTopic:
					server.deviceConfigHandler(msg.Value)
				case common.EngineCfgTopic:
					server.engineConfigHandler(msg.Value)
				case common.HttpsConfig:
					server.httpsConfigHandler(msg.Value)
				default:
					log.WithField("message", msg).Error("[Kafka Consumer] invalid message.")
					continue
				}
				consumer.MarkOffset(msg, "")
			}
		}
	}
}
