## コマンド
### go mod init
- Goモジュールを初期化するためのコマンド
  - モジュール名をパラメータとして受け取り、その名前でモジュールを作成する
- go.modという特別なファイルを作成する
  - モジュール名とモジュールが依存する他のモジュールの情報が含まれる
```bash
$ go mod init example/hello
```
### go mod tidy
- go.sumファイルが作成され、パッケージのインストールをいい感じにするコマンド
  - go.sumには依存関係が記載されている
- 不要なパッケージも削除する
```bash
$ go mod tidy
```

### go mod edit -replace example.com/greetings=../greetings
- example.com/greetingsを./greetingsに置き換えて、依存関係の場所を特定する
- 以下のような行が`go.mod`に追加される
```go
replace example.com/greetings => ../greetings
```
