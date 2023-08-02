# docker-go-postgres
dockerとgoとpostgresqlのアプリケーション

# postgresql
dbはコンテナで立ち上げる
goアプリケーションがコンテナで稼働するpostgresへ接続しに行く

```
	dsn := "host=db user=postgres pass=postgres dbname=postgres	port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

docker-compose側で[db]serviceを定義しているので、go側で[host=db]を指定する