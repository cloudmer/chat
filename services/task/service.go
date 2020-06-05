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
	queue := service.readConfig()
	queue.run()
}

func (service *service) readConfig() queue {
	switch config.Option.Queue.Type {
	case "kafka":
		service.readKafkaConfig()
		return new(kafka)
	case "redis":
		service.readRedisConfig()
		return new(redis)
	default:
		tool.Logger.Fatal().Msg("请配置 queue.type")
	}
	return nil
}

func (service *service) readKafkaConfig() {
	if config.Option.Queue.Kafka.Addr == "" {
		tool.Logger.Fatal().Msg("请配置 queue.kafka.addr")
	}
	if config.Option.Queue.Kafka.Port == 0 {
		tool.Logger.Fatal().Msg("请配置 queue.kafka.port")
	}
}

func (service *service) readRedisConfig() {
	if config.Option.Queue.Redis.Addr == "" {
		tool.Logger.Fatal().Msg("请配置 queue.redis.addr")
	}
	if config.Option.Queue.Redis.Port == 0 {
		tool.Logger.Fatal().Msg("请配置 queue.redis.port")
	}
}