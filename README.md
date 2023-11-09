# fib_api

課題の提出リポジトリです。

## デプロイ先

https://fib-api-kqqc.onrender.com/

## API の仕様

仕様書：https://valiant-shingle-e5c.notion.site/fib-api-6a9d4ef264164479a0b8e9cd4162c29e?pvs=4

### リクエスト

`/fib?n={値}`

### レスポンス　（正常）

ステータス：200  
レスポンス：`{"result": {値}}`

### レスポンス （エラー）

状況：パラメータが渡されなかった場合  
ステータス：400  
レスポンス： `{"code": 400, "error": "Parameter 'n' is required."}`

状況：渡されたパラメータ n が 1 未満の場合  
ステータス：400  
レスポンス：`{"code": 400, "error": "Parameter 'n' is too small."}`

## 環境構築方法

1. リポジトリをクローン
2. 以下のコマンドを実行

- `cd fib-api`
- `go mod tidy`
- `go run main.go`

3. `localhost:8080/`にアクセス

## テストの実行方法

以下のコマンドを実行

- `go test -v .`
