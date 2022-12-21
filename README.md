
# line-dynamo-lambda
下記リポジトリで nextjs と amplifyの検証時に作成したapi
https://github.com/tokatu4561/next-apmlify-sample-todo


## start
```bash
# localのDynamoDB接続用
$ aws configure set aws_access_key_id dummy     --profile local
$ aws configure set aws_secret_access_key dummy --profile local
$ aws configure set region ap-northeast-1       --profile local

# ネットワーク作成 samとdynamodbを同じネットワーク内に配置する
$ docker network create lambda-local
# dynamodbコンテナ作成
$ docker compose up -d --build
```

#### 実行をテスト
```bash
$ make build
$ sam local start-api --docker-network lambda-local
```
