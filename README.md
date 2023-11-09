# fib_api

課題の提出リポジトリです。

## デプロイ先

https://fib-api-kqqc.onrender.com/

## 環境構築方法

1. リポジトリをクローン
2. 以下のコマンドを実行

- `cd fib-api`
- `go mod tidy`
- `go run main.go`

3. `localhost:8080/`にアクセス

## API の仕様

### リクエスト

`/fib?n={値}`

### レスポンス　（正常）

ステータス：200  
レスポンス：`{"result": {値}}`

### レスポンス （エラー）

ステータス：400
レスポンス：`{"error": "1未満のnの値を渡そうとしました。"}`

## テストの実行方法

以下のコマンドを実行

- `cd fib-api`
- `go mod tidy`
- `go test -v .`
