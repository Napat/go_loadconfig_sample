# sample loading config and secret environments

``` cmd
# make run           
go run cmd/main.go
conf before secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"", DBPassword:""}}
conf after secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"", DBPassword:""}}
#
#
# make run_local     
export APPENV=local; go run cmd/main.go
conf before secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"", DBPassword:""}}
2023/01/22 23:30:43 LoadDotEnv env: local, secretPathName: ./config/secret.env
conf after secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"local_db_user", DBPassword:"local_db_password"}}
#
#
# make run_local2
export APPENV=local2; go run cmd/main.go
conf before secret: &config.AppConfig{Server:config.Server{Port:81}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info2..DEV"}, Secrets:config.Secrets{DBUsername:"", DBPassword:""}}
2023/01/22 23:30:45 LoadDotEnv env: local2, secretPathName: ./config/secret.env.local2
conf after secret: &config.AppConfig{Server:config.Server{Port:81}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info2..DEV"}, Secrets:config.Secrets{DBUsername:"local2_db_user", DBPassword:"local2_db_password"}}
#
#
# make run_export_env
export SECRET_DB_USERNAME=exportenv_user SECRET_DB_PASSWORD=exportenv_password; go run cmd/main.go
conf before secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"", DBPassword:""}}
conf after secret: &config.AppConfig{Server:config.Server{Port:80}, Topics:config.Topics{SomeKafkaInfoTopic:"some.kafka.info..DEV"}, Secrets:config.Secrets{DBUsername:"exportenv_user", DBPassword:"exportenv_password"}}
```
