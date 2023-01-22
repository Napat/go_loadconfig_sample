package config

type AppConfig struct {
	Server  Server  `mapstructure:"server"`
	Topics  Topics  `mapstructure:"topics"`
	Secrets Secrets `mapstructure:"secrets"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type Topics struct {
	SomeKafkaInfoTopic string `mapstructure:"some_kafka_info_topic"`
}

type Secrets struct {
	DBUsername string `envconfig:"SECRET_DB_USERNAME"`
	DBPassword string `envconfig:"SECRET_DB_PASSWORD"`
}
