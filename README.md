# jackal
実験用の以下の最低限の機能のみを実装したメッセージアプリ  
Experimental message app that implements only the following functions

- create user
- send message
- fetch message
---

### require Environment variable
```dotenv
# rabbitmq(message broker)のaddr
RABBIT_URL=
# redis(message保存用のdb?)のaddr
REDIS_URL=
# sqlite(userデータ保存用)のpath
SQLITE_PATH=
```

example
```dotenv
RABBIT_URL=amqp://guest:guest@rabbitmq:5672/
REDIS_URL=host.docker.internal:6379
SQLITE_PATH=/data/sqlite/jackal.sqlite
```

### run
server
```shell
# .env.localに環境変数が書いてあったとすると
# If the environment variables are written in .env.local...
docker compose --env-file .env.local up
```

client
```shell
go run cmd/tui/main.go -userid=<YOUR_USERID> -name=<DISPLAYNAME> -receiver=<RECEIVER_USERID>
```


---

### todo
- [x] tui client
- [ ] e2ee
- [x] replace pulsar to other mq(pulsar is too heavy for this project)
    pulsar -> rabbitmq