# jackal
実験用の以下の最低限の機能のみを実装したメッセージアプリ  
Experimental message app that implements only the following functions

- create user
- send message
- fetch message
---

### require Environment variable
```dotenv
# pulsar(message broker)のaddr
PULSAR_URL=
# redis(message保存用のdb?)のaddr
REDIS_URL=
# sqlite(userデータ保存用)のpath
SQLITE_PATH=
```

example
```dotenv
PULSAR_URL=pulsar://host.docker.internal:6650
REDIS_URL=host.docker.internal:6379
SQLITE_PATH=/data/sqlite/jackal.sqlite
```

### run
pulsarの起動に時間がかかるので先に起動しておくと良いかもしれない
```shell
# .env.localに環境変数が書いてあったとすると
# If the environment variables are written in .env.local...
docker compose --env-file .env.local up
```


---

### todo
- tui client
- e2ee
- replace pulsar to other mq(pulsar is too heavy for this project)