package task

import (
	"chat/config"
	"chat/tool"
)

type service struct {
	
}

func New() *service {
	return new(service)	
}

func (service *service) Run()  {
	// 读取配置文件
	service.readConfig()
}

func (service *service) readConfig()  {
	switch config.Config.Queue.Type {
	case "kafka":
		service.readKafkaConfig()
	case "redis":
		service.readRedisConfig()
	default:
		tool.Logger.Fatal().Msg("请配置 queue.type")
	}
}

func (service *service) readKafkaConfig() {
	if config.Config.Queue.Kafka.Addr == "" {
		tool.Logger.Fatal().Msg("请配置 queue.kafka.addr")
	}
	if config.Config.Queue.Kafka.Port == 0 {
		tool.Logger.Fatal().Msg("请配置 queue.kafka.port")
	}
}

func (service *service) readRedisConfig() {
	if config.Config.Queue.Redis.Addr == "" {
		tool.Logger.Fatal().Msg("请配置 queue.redis.addr")
	}
	if config.Config.Queue.Redis.Port == 0 {
		tool.Logger.Fatal().Msg("请配置 queue.redis.port")
	}
}