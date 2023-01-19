# sample loading config and secret environments

``` cmd
# make run      
go run cmd/main.go
conf before secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"", DBPassword:""}}
conf after secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"", DBPassword:""}}

# make run_local
export APPENV=local; go run cmd/main.go
conf before secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"", DBPassword:""}}
LoadSecretEnvFile: local
conf after secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"root", DBPassword:"my_db_password"}}

# make run_sample_env 
export DB_USERNAME=sample_user DB_PASSWORD=sample_password; go run cmd/main.go
conf before secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"", DBPassword:""}}
conf after secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"sample_user", DBPassword:"sample_password"}}
```
