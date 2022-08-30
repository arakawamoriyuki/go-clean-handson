
## tree

```
├── domain                           ドメイン層
│   ├── application                  Application Business Rules
│   │   ├── repository               リポジトリのインターフェース
│   │   │   └── todo.go
│   │   └── usecase                  Usecase
│   │       ├── todo_input_port.go   Input PortのインターフェースとInput Dataの構造定義。ドメインが望む入力形式の定義
│   │       ├── todo_interactor.go   Interactor。Input DataとRepositoryを使ってデータを取得、Output Portの実装であるPresenterに渡す
│   │       └── todo_output_port.go  Output PortのインターフェースとOutput Dataの構造定義。。ドメインが望む出力形式の定義
│   └── entity                       Enterprise Business Rules
│       └── todo.go                  データ構造。DDDにおける値オブジェクトやドメインサービスまで定義していいと思う。
├── go.mod
├── go.sum
├── infrastructure                   Frameworks & Drivers
│   └── router                       Ginの実装。Controllerと繋げる
│       ├── api
│       │   └── todo.go
│       └── router.go
├── interface                        Interface Adapters
│   ├── controller                   Controller。データをInteractorが望むInputDataに整形して渡し、Presenterが整形したデータを表示する
│   │   ├── context.go
│   │   ├── error.go
│   │   └── todo.go
│   ├── presenter                    Presenter。Output Portの実装。OutDataを受け取ってレスポンスを整形する
│   │   └── todo.go
│   └── repository
│       └── todo.go                  リポジトリの実装。Databaseを使ったデータの取得など
├── main.go
├── migrations
│   ├── 000001_create_todos.down.sql
│   └── 000001_create_todos.up.sql
└── readme.md
```

## DB setup

```
$ cd /path/to/go-clean-handson
$ docker compose up

$ mysql --host=127.0.0.1 --port=3306 --user=root --password=pass
mysql> create database `ca-sample` default character set utf8mb4 collate utf8mb4_bin;

$ cd /path/to/go-clean-handson/clean-architecture
$ export DATABASE_URL='mysql://root:pass@tcp(localhost:3306)/ca-sample'
$ migrate -database ${DATABASE_URL} -path migrations up
```