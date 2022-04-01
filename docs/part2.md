# gin 入門

## 概要

- https://gin-gonic.com/ja/docs/
- https://github.com/gin-gonic/gin

> Gin は、Golang で書かれた Web アプリケーションフレームワークです。martini に似たAPIを持ちながら、httprouter のおかげでそれより40倍以上も速いパフォーマンスがあります。良いパフォーマンスと生産性が必要であれば、Gin が好きになれるでしょう。

express sinatra frask に似た軽量フレームワークで、goでWebアプリケーションを作る際によく使われます。

### 特徴

https://gin-gonic.com/ja/docs/introduction/

Ginは以下特徴があります。

- 高速
- ミドルウェアの支援
- クラッシュフリー
- JSON バリデーション
- ルーティングのグループ化
- エラー管理
- 組み込みのレンダリング
- 拡張性

## サンプルプロジェクト

https://gin-gonic.com/ja/docs/quickstart/

公式のサンプルプロジェクトを実行してコードと動作を見てみましょう。

```sh
$ cd /path/to/go-clean-handson/gin-samples/01sample
$ go mod download
$ go run main.go
```

`main.go`

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
```

`GET /ping` ではstatus:200でpongというテキストを返すようになっています。

```sh
$ curl http://localhost:8080/ping
pong
```

`GET /user/:name` ではDBを模したmapオブジェクトからユーザーを引いてデータを返すようになっています。

`:name` などURLを変数化して値を取れています。今はデータがないので `no value` を返しています。

```sh
$ curl http://localhost:8080/user/foo
{"status":"no value","user":"foo"}
```

中盤では `gin.BasicAuth` ミドルウェアが定義されており、`POST /admin` に適用されています。

これにより、`POST /admin` リクエスト時に、事前処理として `gin.BasicAuth` が実行されます。

ミドルウェアは独自の関数も自由に設定できます。

`POST /admin` ではjsonで送られた `value` の値を元にユーザーがDBを模したmapオブジェクトに登録されるようになっています。

`foo` アカウントで `bar` パスワードを使って、Basic認証用トークンを作成し、 `value` に `test` を登録してみましょう。

```sh
$ curl http://localhost:8080/user/foo
{"status":"no value","user":"foo"}

$ echo -n 'foo:bar' | base64
Zm9vOmJhcg==

$ curl -X POST \
    http://localhost:8080/admin \
    -H 'authorization: Basic Zm9vOmJhcg==' \
    -H 'content-type: application/json' \
    -d '{"value":"test"}'
{"status":"ok"}

$ curl http://localhost:8080/user/foo
{"user":"foo","value":"test"}
```

まとめ

- GETやPOSTリクエストの処理方法
- ミドルウェアの作り方
- グローバル変数はリクエスト毎ではなくサーバー起動毎に保持される

### hotload

`go run` に標準ではhotload機能がありません。

goやginに限ったツールではないですが、 `reflex` を利用しましょう。

ファイルの変更に応じてサーバー再起動してくれます。

```
$ go install github.com/cespare/reflex@latest
$ source ~/.zshrc
$ reflex -r '(\.go$|go\.mod)' -s go run ./main.go
```

## APIの使い方のサンプル

https://gin-gonic.com/ja/docs/examples/

公式にサンプルがたくさん揃っています。この中からいくつか見ていきましょう。

- AsciiJSON
- BasicAuth ミドルウェアを使う
- body を異なる構造体にバインドするには
- cookieの設定と取得
- GET,POST,PUT,PATCH,DELETE,OPTIONS メソッドを使う
- Gin を使って複数のサービスを稼働させる
- graceful restart と stop
- HTML をレンダリングする
- HTMLチェックボックスをバインドする
- HTTP/2 サーバープッシュ
- io.Reader からのデータを返す
- JSONP をレンダリングする
- Let's Encrypt のサポート
- Multipart/Urlencoded されたデータをバインドする
- Multipart/Urlencoded フォーム
- PureJSON
- SecureJSON
- URLをバインドする
- XML, JSON, YAML, ProtoBuf をレンダリングする
- カスタム HTTP 設定
- カスタムバリデーション
- カスタムミドルウェア
- カスタムログファイル
- クエリ文字列あるいはポストされたデータをバインドする
- クエリ文字列のみバインドする
- クエリ文字列のパラメータ
- クエリ文字列やフォーム投稿によるパラメータをマッピングする
- テンプレートを含めた1つのバイナリをビルドする
- デフォルトで設定されるミドルウェアがない空の Gin を作成する
- パスに含まれるパラメータ
- ファイルアップロード
- フォーム投稿されたリクエストを構造体にバインドする
- フォーム投稿によるクエリ文字列
- ミドルウェアを利用する
- ミドルウェア内の Goroutine
- モデルへのバインディングとバリデーションする
- リダイレクト
- ルーティングをグループ化する
- ルーティングログのフォーマットを定義する
- ログファイルへ書き込むには
- ログ出力の色付けを制御する
- 複数のテンプレート
- 静的ファイルを返す

### ミドルウェアとルーティング

ミドルウェアとルーティングの定義方法についてもう少しふかぼってみてみましょう。

```sh
$ cd /path/to/go-clean-handson/gin-samples/02middleware
$ go mod download
$ go run main.go
```

`main.go`

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyBenchLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        // request 処理の前
		t := time.Now()

		// サンプル変数を設定
		c.Set("example", "12345")

		c.Next()

		// request 処理の後
		// レイテンシ表示
		latency := time.Since(t)
		log.Print(latency)

		// 送信予定のステータスコードを表示
		status := c.Writer.Status()
		log.Println(status)
	}
}

func AuthRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// // 認証失敗させる
		// c.JSON(http.StatusUnauthorized, gin.H{"error": true, "from": "AuthRequiredMiddleware"})
		// c.Abort()
	}
}

func benchmarkEndpoint(c *gin.Context) {
	// MyBenchLoggerMiddlewareで設定された変数を表示
	example := c.MustGet("example").(string)
	log.Println(example)

	c.JSON(http.StatusOK, gin.H{"error": false, "from": "benchmarkEndpoint"})
}

func meEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"error": false, "from": "meEndpoint"})
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/benchmark", MyBenchLoggerMiddleware(), benchmarkEndpoint)

	authorized := r.Group("/auth")
	authorized.Use(AuthRequiredMiddleware())
	{
		authorized.GET("/me", meEndpoint)
	}

	r.Run(":8080")
}
```

`GET /benchmark` では `MyBenchLoggerMiddleware` ミドルウェアが設定されています。

`MyBenchLoggerMiddleware` では `c.Next()` を利用して事前事後処理を定義している例です。

事前処理として現在時刻を保持しておき、かつ `c.Set` を使って変数を設定しています。

`c.Set` を使って定義された変数は後続のハンドラで利用することができます。

また、事後処理として処理前の時間との現在時刻を比較してレイテンシを計算表示しています。

さらに、返す予定のステータスコードなども確認することができます。

```sh
$ curl http://localhost:8080/benchmark
{"error":false,"from":"benchmarkEndpoint"}
```

サーバー側のログは以下のようになります。

```
2022/03/25 13:06:03 12345
2022/03/25 13:06:03 553.765µs
2022/03/25 13:06:03 200
[GIN] 2022/03/25 - 13:06:03 | 200 |     566.854µs |             ::1 | GET      "/benchmark"
```

`GET /auth/me` もみてみましょう。

以下コードで `/auth` をグループ化し、 `AuthRequiredMiddleware` を使うように設定されています。

これにより、内部のルーティングはすべて `/auth` が着くようになり、`AuthRequiredMiddleware` が適用されます。

```go
authorized := r.Group("/auth")
authorized.Use(AuthRequiredMiddleware())
```

実際に叩いて確認してみましょう。

```sh
$ curl http://localhost:8080/auth/me
{"error":false,"from":"meEndpoint"}
```

`AuthRequiredMiddleware` が特に何もしていないのでそのまま通ります。

`AuthRequiredMiddleware` のコメントアウトを外し、認証失敗をシミュレーションしてみましょう。

```go
func AuthRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 認証失敗させる
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "from": "AuthRequiredMiddleware"})
		c.Abort()
	}
}
```

```sh
$ curl http://localhost:8080/auth/me
{"error":true,"from":"AuthRequiredMiddleware"}
```

まとめ

- カスタムミドルウェアの定義方法と情報伝達、事前事後処理、処理の中断方法
- ルーティングのグループ化方法とミドルウェアの適用方法

参考

- ミドルウェアを利用する https://gin-gonic.com/ja/docs/examples/using-middleware/
- カスタムミドルウェア https://gin-gonic.com/ja/docs/examples/custom-middleware/
- ルーティングをグループ化 https://gin-gonic.com/ja/docs/examples/grouping-routes/

### バインドやバリデーション

リクエストパス、クエリ、ボディパラメータをgoの構造体へのバインドする方法についてみていきます。

```sh
$ cd /path/to/go-clean-handson/gin-samples/03bind
$ go mod download
$ go run main.go
```

`main.go`

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	ID    int    `uri:"id" json:"id" binding:"required"`
	Title string `form:"title" json:"title"`
	Score int    `form:"score" json:"score"`
}

func sampleHandler(c *gin.Context) {
	request := Request{}
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	r.POST("/sample/:id", sampleHandler)

	r.Run(":8080")
}
```

`POST /sample/:id` では

パスパラメータとしてURLに `:id` をint型で

フォーム(jsonボディやformdataやクエリ)パラメータとして `title` をstring型、`score` をint型で受け取るように設定しています。

また、リクエストパラメータをそのまま返すようになっています。

それぞれURLやクエリパラメータ、Bodyパラメータをbindするために以下関数などが使われています。

- ShouldBind https://github.com/gin-gonic/gin#bind-query-string-or-post-data
- ShouldBindUri https://github.com/gin-gonic/gin#bind-uri
- ShouldBindHeader https://github.com/gin-gonic/gin#bind-header

実際に試してみましょう。

```sh
$ curl -X POST 'http://localhost:8080/sample/1?score=2' -H 'content-type: application/json' -d '{"title":"test"}'
{"id":1,"title":"test","score":2}
```

また、構造や型にそぐわない形式の場合エラーを返すのでためしてみましょう。

```sh
$ curl -X POST 'http://localhost:8080/sample/aaa?score=2' -H 'content-type: application/json' -d '{"title":"test"}'
{"error":"strconv.ParseInt: parsing \"aaa\": invalid syntax"}
```

パスパラメータを `uri` フォームパラメータを `form` バリデーションなどの定義に `binding` 返す値の構造として `json` などを構造体に定義することができます。

このようなstrictのメタ情報をタグと言い、さまざまな処理に利用されます。

まとめ

- パラメータの変数へのバインド方法
- バリデーションも含めた構造体のタグ設定方法

参考

- ShouldBind https://github.com/gin-gonic/gin#bind-query-string-or-post-data
- ShouldBindUri https://github.com/gin-gonic/gin#bind-uri
- ShouldBindHeader https://github.com/gin-gonic/gin#bind-header
- body を異なる構造体にバインドするには(jsonのバインド) https://gin-gonic.com/ja/docs/examples/bind-body-into-dirrerent-structs/
- クエリ文字列あるいはポストされたデータをバインドする https://gin-gonic.com/ja/docs/examples/bind-query-or-post/
- パスに含まれるパラメータ https://gin-gonic.com/ja/docs/examples/param-in-path/
- モデルへのバインディングとバリデーション https://gin-gonic.com/ja/docs/examples/binding-and-validation/
- Ginフレームワークでのカスタムバリデーション https://qiita.com/emonuh/items/3b531fded9b5ede9d93a
- リクエストパラメータ向けValidationパターンまとめ https://qiita.com/RunEagler/items/ad79fc860c3689797ccc

### テスト

今回はテスト方法についてみていきます。

```sh
$ cd /path/to/go-clean-handson/gin-samples/04test
$ go mod download
```

まずはginのテスト方法についてみてみましょう。

*本来はテストとテスト対象の関数は別々に定義するものですが、今回はわかりやすいよう1ファイルに定義しています。

`request_test.go`

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
```

サーバーを起動し、 `GET /ping` へリクエストし、ステータスコードとボディの内容を検証しています。

さっそく実行してみましょう。

```sh
$ go test -v request_test.go
=== RUN   TestPingRoute
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> command-line-arguments.setupRouter.func1 (3 handlers)
[GIN] 2022/03/25 - 13:53:15 | 200 |       7.444µs |                 | GET      "/ping"
--- PASS: TestPingRoute (0.00s)
PASS
ok  	command-line-arguments	0.224s
```

返すステータスコードを201などに変更して失敗例もみてみましょう。

```go
assert.Equal(t, 201, w.Code)
```

```sh
$ go test -v request_test.go
=== RUN   TestPingRoute
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> command-line-arguments.setupRouter.func1 (3 handlers)
[GIN] 2022/03/25 - 13:53:46 | 200 |      11.612µs |                 | GET      "/ping"
    request_test.go:29:
        	Error Trace:	request_test.go:29
        	Error:      	Not equal:
        	            	expected: 201
        	            	actual  : 200
        	Test:       	TestPingRoute
--- FAIL: TestPingRoute (0.00s)
FAIL
FAIL	command-line-arguments	0.201s
FAIL
```

ここまででginのテスト方法を見てきましたが、goでは `TableDrivenTests` というテスト方法が推奨されています。

以下の正方形の面積を求める `Square` 関数をTableDrivenTestsでテストしてみます。

`tabledriven_test.go`

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Square(a, b int) int {
	return a * b
}

func TestSquere(t *testing.T) {
	asserts := assert.New(t)
	tests := []struct {
		title  string
		input  []int
		output int
	}{
		{
			title:  "2x3の面積は6になる",
			input:  []int{2, 3},
			output: 6,
		},
		{
			title:  "0x1の面積は0になる",
			input:  []int{0, 1},
			output: 0,
		},
	}

	for _, td := range tests {
		td := td
		t.Run("Square:"+td.title, func(t *testing.T) {
			result := Square(td.input[0], td.input[1])
			asserts.Equal(td.output, result)
		})
	}
}
```

入力と出力のセットの配列をループさせ、入力を渡して出力を検証しています。

```sh
$ go test -v tabledriven_test.go
=== RUN   TestSquere
=== RUN   TestSquere/Square:2x3の面積は6になる
=== RUN   TestSquere/Square:0x1の面積は0になる
--- PASS: TestSquere (0.00s)
    --- PASS: TestSquere/Square:2x3の面積は6になる (0.00s)
    --- PASS: TestSquere/Square:0x1の面積は0になる (0.00s)
PASS
ok  	command-line-arguments	0.513s
```

ここまででginのテストとgoの推奨する `TableDrivenTests` の方法をみてきました。

本来は `*_test.go` というファイルに `Test*` という関数の命名規則でテストを定義します。

また、以下コマンドで全て下位ディレクトリに対してテストを実行することができます。

```sh
$ go test -cover -v ./...
```

まとめ

- Ginのテスト方法
- TableDrivenテストの方法

参考

- テスト https://gin-gonic.com/ja/docs/testing/
- TableDrivenTests https://github.com/golang/go/wiki/TableDrivenTests

### HTMLレンダリング

今回の対象外ですが参考程度にHTMLレンダリング方法についてのサンプルを貼っておきます。

- HTML をレンダリングする https://gin-gonic.com/ja/docs/examples/html-rendering/
- HTMLチェックボックスをバインドする https://gin-gonic.com/ja/docs/examples/bind-html-checkbox/

## GinでTODOリストAPIの作成

ここからはGinでTODOリストAPIを作成する全てを学んでいきましょう。

まずはデータベースについて深ぼっていきます。

### データベース

データベース関連(goのマイグレーションとO/RMapper)を利用した例を見ていきます。

以下コマンドでdbを立てておきましょう。

```sh
$ docker compose up
```

#### マイグレーション

いくつかgoプロジェクトで使われるマイグレーションツールを比較してみます。

|ライブラリ名|star|urlや参考情報|概要|
|:-:|:-:|:-:|:-:|
|gorm (auto migration)|27.3k|https://gorm.io/ja_JP/docs/migration.html|一般的に利用されているORMであるgormのマイグレーション機能。基本的にオートマイグレーションで利用しないカラムの削除はされないなど注意が必要。CLIないのでタイミングをコントロールできない|
|golang-migrate|8k|https://github.com/golang-migrate/migrate<br/>https://dev.classmethod.jp/articles/db-migrate-with-golang-migrate/|SQLで定義|
|goose|2k|https://github.com/pressly/goose|SQLで定義|
|sql-migrate|2k|https://github.com/rubenv/sql-migrate<br/>https://ryotarch.com/go-lang/implement-migrate-and-seed-with-sql-migrate-and-gorm/|SQLで定義|
|Gormigrate|746|https://github.com/go-gormigrate/gormigrate|structで定義|
|migu|86|https://github.com/naoina/migu|structで定義|


スタンダードなツールは定まっていない印象です。

gormはオートマイグレーションで注意点が多いので除外し、個人的にはDBの差を吸収して欲しいのでstructやDSLで書く方が好きだが、star数が少ないので除外します。

今回は [golang-migrate](https://github.com/golang-migrate/migrate) を利用した例を紹介します。

参考

- Go製マイグレーションツールまとめ https://qiita.com/nownabe/items/1acce9f6b9f14f74c965

##### golang-migrate install

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

brewやgo getなどでインストールできます。

```sh
$ brew install golang-migrate
$ source ~/.zshrc
$ migrate --version
v4.15.1
```

##### migration作成

適当なTODOアプリを作るディレクトリを作成し、移動してください。

```sh
$ mkdir gin-sample
$ cd gin-sample
```

以下コマンドでマイグレーションファイルを作成します。

```sh
$ migrate create -ext sql -dir migrations -seq create_todos
```

作成されたファイルにTODOテーブルを作る以下SQLを記入します。

`migrations/000001_create_todos.up.sql`

```sql
CREATE TABLE IF NOT EXISTS `todos` (
    `id` BIGINT(20) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(256) NOT NULL,
    `done` BOOLEAN NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME
) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
```

`migrations/000001_create_todos.down.sql`

```sql
DROP TABLE IF EXISTS `todos`;
```

##### migration実行

さっそくテーブルを作ってみましょう。

指定するDATABASE_URLの `localhost:3306` は `tcp()` で囲わないといけないことに注意しましょう。

```sh
$ export DATABASE_URL='mysql://root:pass@tcp(localhost:3306)/gin-sample'

$ migrate -database ${DATABASE_URL} -path migrations up
1/u create_todos (86.991731ms)
```

作成できたらDBクライアントで接続し、テーブルが作成されているか確認しましょう。

```
$ mysql --host=127.0.0.1 --port=3306 --user=root --password=pass

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| gin-sample         |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+

mysql> use gin-sample;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A
Database changed

mysql> show tables;
+----------------------+
| Tables_in_gin-sample |
+----------------------+
| schema_migrations    |
| todos                |
+----------------------+
2 rows in set (0.01 sec)
```

また、再度 `migrate up` した際に(既に適用済みなので)変更されないか、 `migrate down` した時にテーブルが消えているかなども確認してみましょう。

```
$ migrate -database ${DATABASE_URL} -path migrations up
no change

$ migrate -database ${DATABASE_URL} -path migrations down
Are you sure you want to apply all down migrations? [y/N]
y
Applying all down migrations
1/d create_todos (108.060026ms)

$ migrate -database ${DATABASE_URL} -path migrations up
1/u create_todos (86.991731ms)
```

参考

- https://github.com/golang-migrate/migrate/tree/master/database/mysql

#### データベースライブラリ、O/RMapperの選定

|ライブラリ名|star|urlや参考情報|概要|
|:-:|:-:|:-:|:-:|
|database/sql|-|-|標準ライブラリ。リッチではなくSQLを利用したクエリ。構造体にマッピングができない|
|gorm|27k|https://github.com/go-gorm/gorm|高機能。構造体マッピング可能。ORM|
|sqlx|11k|https://github.com/jmoiron/sqlx|軽量。構造体マッピング。基本的にSQL|
|gorp|3k|https://github.com/go-gorp/gorp|構造体マッピング。デフォルトでSQL|

今回はO/RMapperとして利用できる `gorm` を使った例を見ていきます。

参考

- https://blog.p1ass.com/posts/go-database-sql-wrapper/

##### gorm

gormを利用したデータベースの操作についてみていきましょう。

```sh
$ cd /path/to/go-clean-handson/gin-samples/05orm
$ go mod download
```

`main.go`

```go
package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func main() {
	dsn := "root:pass@tcp(localhost:3306)/gin-sample"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create
	if result := db.Create(&Todo{Name: "test", Done: false}); result.Error != nil {
		log.Fatal(result.Error)
		return
	}
}
```

まずはレコードの作成例です。 `db.Create` 対象のレコードオブジェクトを渡して作成します。

結果のErrorプロパティにエラーが入っているのでそれを検証してハンドリングできます。

また、Todo構造体に `gorm.Model` をembed(埋め込み)していますが、これを定義することにより、以下プロパティをもつことになります。

```go
// gorm.Modelの定義
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

実行してデータが作成されるか確認しましょう。

```sh
$ go run main.go
```

次は取得の例です。

`db.First` に変数を渡してバインドさせます。また、条件を指定することも可能です。

これも実行してログを確認してみましょう。

```go
// Read
var todo Todo
if result := db.First(&todo, 1); result.Error != nil {
    log.Fatal(result.Error)
    return
}
fmt.Println(todo.Name)

if result := db.First(&todo, "name = ?", "test"); result.Error != nil {
    log.Fatal(result.Error)
    return
}
fmt.Println(todo.Name)
```

次は更新の例です。

`Update` を利用して変更します。また、`Updates` を利用して複数のカラムを変更できます。

```go
// Update
if result := db.Model(&todo).Update("name", "changed"); result.Error != nil {
    log.Fatal(result.Error)
    return
}
fmt.Println(todo.Name)

if result := db.Model(&todo).Updates(map[string]interface{}{"name": "changed 3", "done": false}); result.Error != nil {
    log.Fatal(result.Error)
    return
}
fmt.Println(todo.Name)
fmt.Println(todo.Done)
```

次は削除の例です。

`Delete` を利用して削除します。また、`gorm.Model` を埋め込んだ構造体の場合は `deleted_at` に日付が挿入される論理削除になります。

```go
// Delete (soft delete)
if result := db.Delete(&todo, 1); result.Error != nil {
    log.Fatal(result.Error)
    return
}
```

参考

- https://gorm.io/ja_JP/docs/index.html
- https://gorm.io/ja_JP/docs/models.html#gorm-Model

### 最低限のルーティング定義

DB操作を学んだのでGinで本格的にAPIを作成していきます。

先ほど作った(migrationsディレクトリのある)TODOリストプロジェクトに移動してください。

```sh
$ cd /path/to/gin-sample
```

必要なパッケージを定義してパッケージをダウンロードしておきましょう。

`go.mod`

```go.mod
module main

go 1.17

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.4
	github.com/stretchr/testify v1.7.1
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70 // indirect
	golang.org/x/sys v0.0.0-20220307203707-22a9840ba4d7 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
```

`go.sum`

```go.sum
github.com/creack/pty v1.1.9/go.mod h1:oKZEueFk5CKHvIhNR5MUki03XCEU+Q6VDXinZuGJ33E=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/gin-contrib/sse v0.1.0 h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=
github.com/gin-contrib/sse v0.1.0/go.mod h1:RHrZQHXnP2xjPF+u1gW/2HnVO7nvIa9PG3Gm+fLHvGI=
github.com/gin-gonic/gin v1.7.7 h1:3DoBmSbJbZAWqXJC3SLjAPfutPJJRN1U5pALB7EeTTs=
github.com/gin-gonic/gin v1.7.7/go.mod h1:axIBovoeJpVj8S3BwE0uPMTeReE4+AfFtqpqaZ1qq1U=
github.com/go-ini/ini v1.66.4 h1:dKjMqkcbkzfddhIhyglTPgMoJnkvmG+bSLrU9cTHc5M=
github.com/go-ini/ini v1.66.4/go.mod h1:ByCAeIL28uOIIG0E3PJtZPDL8WnHpFKFOtgjp+3Ies8=
github.com/go-playground/assert/v2 v2.0.1 h1:MsBgLAaY856+nPRTKrp3/OZK38U/wa0CcBYNjji3q3A=
github.com/go-playground/assert/v2 v2.0.1/go.mod h1:VDjEfimB/XKnb+ZQfWdccd7VUvScMdVu0Titje2rxJ4=
github.com/go-playground/locales v0.13.0/go.mod h1:taPMhCMXrRLJO55olJkUXHZBHCxTMfnGwq/HNwmWNS8=
github.com/go-playground/locales v0.14.0 h1:u50s323jtVGugKlcYeyzC0etD1HifMjqmJqb8WugfUU=
github.com/go-playground/locales v0.14.0/go.mod h1:sawfccIbzZTqEDETgFXqTho0QybSa7l++s0DH+LDiLs=
github.com/go-playground/universal-translator v0.17.0/go.mod h1:UkSxE5sNxxRwHyU+Scu5vgOQjsIJAF8j9muTVoKLVtA=
github.com/go-playground/universal-translator v0.18.0 h1:82dyy6p4OuJq4/CByFNOn/jYrnRPArHwAcmLoJZxyho=
github.com/go-playground/universal-translator v0.18.0/go.mod h1:UvRDBj+xPUEGrFYl+lu/H90nyDXpg0fqeB/AQUGNTVA=
github.com/go-playground/validator/v10 v10.4.1/go.mod h1:nlOn6nFhuKACm19sB/8EGNn9GlaMV7XkbRSipzJ0Ii4=
github.com/go-playground/validator/v10 v10.10.1 h1:uA0+amWMiglNZKZ9FJRKUAe9U3RX91eVn1JYXMWt7ig=
github.com/go-playground/validator/v10 v10.10.1/go.mod h1:i+3WkQ1FvaUjjxh1kSvIA4dMGDBiPU55YFDl0WbKdWU=
github.com/go-sql-driver/mysql v1.6.0 h1:BCTh4TKNUYmOmMUcQ3IipzF5prigylS7XXjEkfCHuOE=
github.com/go-sql-driver/mysql v1.6.0/go.mod h1:DCzpHaOWr8IXmIStZouvnhqoel9Qv2LBy8hT2VhHyBg=
github.com/golang/protobuf v1.3.3/go.mod h1:vzj43D7+SQXF/4pzW/hwtAqwc6iTitCiVSaWz5lYuqw=
github.com/golang/protobuf v1.5.0/go.mod h1:FsONVRAS9T7sI+LIUmWTfcYkHO4aIWwzhcaSAoJOfIk=
github.com/golang/protobuf v1.5.2 h1:ROPKBNFfQgOUMifHyP+KYbvpjbdoFNs+aK7DXlji0Tw=
github.com/golang/protobuf v1.5.2/go.mod h1:XVQd3VNwM+JqD3oG2Ue2ip4fOMUkwXdXDdiuN0vRsmY=
github.com/google/go-cmp v0.5.5 h1:Khx7svrCpmxxtHBq5j2mp/xVjsi8hQMfNLvJFAlrGgU=
github.com/google/go-cmp v0.5.5/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/gofuzz v1.0.0/go.mod h1:dBl0BpW6vV/+mYPU4Po3pmUjxk6FQPldtuIdl/M65Eg=
github.com/jinzhu/inflection v1.0.0 h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=
github.com/jinzhu/inflection v1.0.0/go.mod h1:h+uFLlag+Qp1Va5pdKtLDYj+kHp5pxUVkryuEj+Srlc=
github.com/jinzhu/now v1.1.4/go.mod h1:d3SSVoowX0Lcu0IBviAWJpolVfI5UJVZZ7cO71lE/z8=
github.com/jinzhu/now v1.1.5 h1:/o9tlHleP7gOFmsnYNz3RGnqzefHA47wQpKrrdTIwXQ=
github.com/jinzhu/now v1.1.5/go.mod h1:d3SSVoowX0Lcu0IBviAWJpolVfI5UJVZZ7cO71lE/z8=
github.com/json-iterator/go v1.1.9/go.mod h1:KdQUCv79m/52Kvf8AW2vK1V8akMuk1QjK/uOdHXbAo4=
github.com/json-iterator/go v1.1.12 h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=
github.com/json-iterator/go v1.1.12/go.mod h1:e30LSqwooZae/UwlEbR2852Gd8hjQvJoHmT4TnhNGBo=
github.com/kr/pretty v0.1.0/go.mod h1:dAy3ld7l9f0ibDNOQOHHMYYIIbhfbHSm3C4ZsoJORNo=
github.com/kr/pretty v0.2.1/go.mod h1:ipq/a2n7PKx3OHsz4KJII5eveXtPO4qwEXGdVfWzfnI=
github.com/kr/pretty v0.3.0 h1:WgNl7dwNpEZ6jJ9k1snq4pZsg7DOEN8hP9Xw0Tsjwk0=
github.com/kr/pretty v0.3.0/go.mod h1:640gp4NfQd8pI5XOwp5fnNeVWj67G7CFk/SaSQn7NBk=
github.com/kr/pty v1.1.1/go.mod h1:pFQYn66WHrOpPYNljwOMqo10TkYh1fy3cYio2l3bCsQ=
github.com/kr/text v0.1.0/go.mod h1:4Jbv+DJW3UT/LiOwJeYQe1efqtUx/iVham/4vfdArNI=
github.com/kr/text v0.2.0 h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=
github.com/kr/text v0.2.0/go.mod h1:eLer722TekiGuMkidMxC/pM04lWEeraHUUmBw8l2grE=
github.com/leodido/go-urn v1.2.0/go.mod h1:+8+nEpDfqqsY+g338gtMEUOtuK+4dEMhiQEgxpxOKII=
github.com/leodido/go-urn v1.2.1 h1:BqpAaACuzVSgi/VLzGZIobT2z4v53pjosyNd9Yv6n/w=
github.com/leodido/go-urn v1.2.1/go.mod h1:zt4jvISO2HfUBqxjfIshjdMTYS56ZS/qv49ictyFfxY=
github.com/mattn/go-isatty v0.0.12/go.mod h1:cbi8OIDigv2wuxKPP5vlRcQ1OAZbq2CE4Kysco4FUpU=
github.com/mattn/go-isatty v0.0.14 h1:yVuAays6BHfxijgZPzw+3Zlu5yQgKGP2/hcQbHb7S9Y=
github.com/mattn/go-isatty v0.0.14/go.mod h1:7GGIvUiUoEMVVmxf/4nioHXj79iQHKdU27kJ6hsGG94=
github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421/go.mod h1:6dJC0mAP4ikYIbvyc7fijjWJddQyLn8Ig3JB5CqoB9Q=
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd/go.mod h1:6dJC0mAP4ikYIbvyc7fijjWJddQyLn8Ig3JB5CqoB9Q=
github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742/go.mod h1:bx2lNnkwVCuqBIxFjflWJWanXIb3RllmbCylyMrvgv0=
github.com/modern-go/reflect2 v1.0.2 h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=
github.com/modern-go/reflect2 v1.0.2/go.mod h1:yWuevngMOJpCy52FWWMvUC8ws7m/LJsjYzDa0/r8luk=
github.com/pkg/diff v0.0.0-20210226163009-20ebb0f2a09e/go.mod h1:pJLUxLENpZxwdsKMEsNbx1VGcRFpLqf3715MtcvvzbA=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/rogpeppe/go-internal v1.6.1/go.mod h1:xXDCJY+GAPziupqXw64V24skbSoqbTEfhy4qGm1nDQc=
github.com/rogpeppe/go-internal v1.8.0 h1:FCbCCtXNOY3UtUuHUYaghJg4y7Fd14rXifAYUAtL9R8=
github.com/rogpeppe/go-internal v1.8.0/go.mod h1:WmiCO8CzOY8rg0OYDC4/i/2WRWAB6poM+XZ2dLUbcbE=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=
github.com/stretchr/testify v1.4.0/go.mod h1:j7eGeouHqKxXV5pUuKE4zz7dFj8WfuZ+81PSLYec5m4=
github.com/stretchr/testify v1.6.1/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.7.1 h1:5TQK59W5E3v0r2duFAb7P95B6hEeOyEnHRa8MjYSMTY=
github.com/stretchr/testify v1.7.1/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/ugorji/go v1.1.7/go.mod h1:kZn38zHttfInRq0xu/PH0az30d+z6vm202qpg1oXVMw=
github.com/ugorji/go v1.2.7 h1:qYhyWUUd6WbiM+C6JZAUkIJt/1WrjzNHY9+KCIjVqTo=
github.com/ugorji/go v1.2.7/go.mod h1:nF9osbDWLy6bDVv/Rtoh6QgnvNDpmCalQV5urGCCS6M=
github.com/ugorji/go/codec v1.1.7/go.mod h1:Ax+UKWsSmolVDwsd+7N3ZtXu+yMGCf907BLYF3GoBXY=
github.com/ugorji/go/codec v1.2.7 h1:YPXUKf7fYbp/y8xloBqZOw2qaVggbfwMlI8WM3wZUJ0=
github.com/ugorji/go/codec v1.2.7/go.mod h1:WGN1fab3R1fzQlVQTkfxVtIBhWDRqOviHU95kRgeqEY=
golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3/go.mod h1:IxCIyHEi3zRg3s0A5j5BB6A9Jmi73HwBIUl50j+osU4=
golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70 h1:syTAU9FwmvzEoIYMqcPHOcVm4H3U5u90WsvuYgwpETU=
golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70/go.mod h1:IxCIyHEi3zRg3s0A5j5BB6A9Jmi73HwBIUl50j+osU4=
golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2/go.mod h1:9nx3DQGgdP8bBQD5qxJ1jj9UTztislL4KSBs9R2vV5Y=
golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20190412213103-97732733099d/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20200116001909-b77594299b42/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210423082822-04245dca01da/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20220307203707-22a9840ba4d7 h1:8IVLkfbr2cLhv0a/vKq4UFUcJym8RmDoDboxCFWEjYE=
golang.org/x/sys v0.0.0-20220307203707-22a9840ba4d7/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
golang.org/x/text v0.3.2/go.mod h1:bEr9sfX3Q8Zfm5fL9x+3itogRgK3+ptLWKqgva+5dAk=
golang.org/x/text v0.3.6/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.7 h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=
golang.org/x/text v0.3.7/go.mod h1:u+2+/6zg+i71rQMx5EYifcz6MCKuco9NR6JIITiCfzQ=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=
golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
google.golang.org/protobuf v1.26.0-rc.1/go.mod h1:jlhhOSvTdKEhbULTjvd4ARK9grFBp09yW+WbY/TyQbw=
google.golang.org/protobuf v1.26.0/go.mod h1:9q0QmTI4eRPtz6boOQmLYwt+qCgq0jsYwAQnmE0givc=
google.golang.org/protobuf v1.27.1 h1:SnqbnDw1V7RiZcXPx5MEeqPv2s79L9i7BJUlG/+RurQ=
google.golang.org/protobuf v1.27.1/go.mod h1:9q0QmTI4eRPtz6boOQmLYwt+qCgq0jsYwAQnmE0givc=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c/go.mod h1:JHkPIbrfpd72SG/EVd6muEfDQjcINNoR0C8j2r3qZ4Q=
gopkg.in/errgo.v2 v2.1.0/go.mod h1:hNsd1EY+bozCKY1Ytp96fpM3vjJbqLJn88ws8XvfDNI=
gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.2.8/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.4.0 h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=
gopkg.in/yaml.v2 v2.4.0/go.mod h1:RDklbk79AGWmwhnvt/jBztapEOGDOx6ZbXqjP6csGnQ=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b h1:h8qDotaEPuJATrMmW04NCwg7v22aHH28wwpauUhK9Oo=
gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gorm.io/driver/mysql v1.3.2 h1:QJryWiqQ91EvZ0jZL48NOpdlPdMjdip1hQ8bTgo4H7I=
gorm.io/driver/mysql v1.3.2/go.mod h1:ChK6AHbHgDCFZyJp0F+BmVGb06PSIoh9uVYKAlRbb2U=
gorm.io/gorm v1.23.1/go.mod h1:l2lP/RyAtc1ynaTjFksBde/O8v9oOGIApu2/xRitmZk=
gorm.io/gorm v1.23.2 h1:xmq9QRMWL8HTJyhAUBXy8FqIIQCYESeKfJL4DoGKiWQ=
gorm.io/gorm v1.23.2/go.mod h1:l2lP/RyAtc1ynaTjFksBde/O8v9oOGIApu2/xRitmZk=
```

```sh
$ go mod download
```

以下 `main.go` を作成し、最低限のルーティング定義をして動作確認してみましょう。

`main.go`

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("api")
	{
		todosGroup := apiGroup.Group("todos")
		{
			todosGroup.GET("", func(c *gin.Context) {
				c.String(http.StatusOK, "GET /api/todos")
			})
			todosGroup.GET(":id", func(c *gin.Context) {
				id := c.Params.ByName("id")
				c.String(http.StatusOK, fmt.Sprintf("GET /api/todos/%s", id))
			})
			todosGroup.POST("", func(c *gin.Context) {
				c.String(http.StatusCreated, "POST /api/todos")
			})
			todosGroup.PATCH(":id", func(c *gin.Context) {
				id := c.Params.ByName("id")
				c.String(http.StatusOK, fmt.Sprintf("PATCH /api/todos/%s", id))
			})
			todosGroup.DELETE(":id", func(c *gin.Context) {
				id := c.Params.ByName("id")
				c.String(http.StatusNoContent, fmt.Sprintf("DELETE /api/todos/%s", id))
			})
		}
	}

	return router
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
```

hotloadを有効にするため以下でサーバー起動します。

```sh
$ reflex -r '(\.go$|go\.mod)' -s go run ./main.go
```

各APIを試してみましょう。

```sh
$ curl -X GET 'http://localhost:8080/api/todos' -H 'content-type: application/json'
$ curl -X GET 'http://localhost:8080/api/todos/1' -H 'content-type: application/json'
$ curl -X POST 'http://localhost:8080/api/todos' -H 'content-type: application/json' -d '{"title":"test","done":false}'
$ curl -X PATCH 'http://localhost:8080/api/todos/1' -H 'content-type: application/json' -d '{"title":"changed","done":true}'
$ curl -X DELETE 'http://localhost:8080/api/todos/1' -H 'content-type: application/json'
```

### ルーティングをroutersに移動

最低限のルーティングを定義できたので、
次に処理を分けるために `routers/api/todos.go` に実際の処理を、 `routers/router.go` にルーティング定義を移動します。

`routers/api/todos.go`

```go
package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	c.String(http.StatusOK, "GET /api/todos")
}

func GetTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	c.String(http.StatusOK, fmt.Sprintf("GET /api/todos/%s", id))
}

func CreateTodo(c *gin.Context) {
	c.String(http.StatusCreated, "POST /api/todos")
}

func UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	c.String(http.StatusOK, fmt.Sprintf("PATCH /api/todos/%s", id))
}

func DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	c.String(http.StatusNoContent, fmt.Sprintf("DELETE /api/todos/%s", id))
}
```

`routers/router.go`

```go
package routers

import (
	"github.com/gin-gonic/gin"

	"main/routers/api"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("api")
	{
		todosGroup := apiGroup.Group("todos")
		{
			todosGroup.GET("", api.GetTodos)
			todosGroup.GET(":id", api.GetTodo)
			todosGroup.POST("", api.CreateTodo)
			todosGroup.PATCH(":id", api.UpdateTodo)
			todosGroup.DELETE(":id", api.DeleteTodo)
		}
	}

	return router
}

```

`main.go`

```go
package main

import (
	"main/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
```

先ほどと同じようにAPIが動作していることを確認しましょう。

### 設定

データベースの接続情報などの設定は開発/ステージング/本番/テストで変えられるようにしましょう。

今回は `.ini` ファイルを読み込む処理を作ります。

`conf/development.ini`

```ini
[database]
Type = mysql
User = root
Password = pass
Host = localhost
Port = 3306
Name = gin-sample
```

`pkg/setting/setting.go`

```go
package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

func Setup(iniPath string) {
	var err error
	cfg, err = ini.Load(iniPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse %s: %v", iniPath, err)
	}

	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
```

`main.go`

```diff
package main

import (
	"main/models"
+ 	"main/pkg/setting"
	"main/routers"
)

+ func init() {
+ 	setting.Setup("conf/development.ini")
+ }

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
```

main packageのinit関数で初期化処理が可能です。

設定ファイルパスを渡して `conf/development.ini` が読み取れているか確認してみましょう。

`main.go`

```diff
package main

import (
+ 	"fmt"
	"main/pkg/setting"
	"main/routers"
)

func init() {
	setting.Setup()
}

func main() {
+ 	fmt.Println(setting.DatabaseSetting.Host)

	r := routers.SetupRouter()
	r.Run(":8080")
}
```

### モデル定義

次にモデルを定義します。

以下 `models/models.go` はデータベースの接続セットアップし、`db` 変数に保持します。

`models/models.go`

```go
package models

import (
	"fmt"
	"log"
	"main/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Setup() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name,
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}
```

`main.go`

```diff
package main

import (
+ 	"main/models"
	"main/pkg/setting"
	"main/routers"
)

func init() {
	setting.Setup()
+ 	models.Setup()
}

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
```

次に実際のDB操作を行うtodoモデルを作っていきましょう。

`models/todos.go`

```go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int            `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	Done      bool           `json:"done"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func GetTodos() ([]*Todo, error) {
	var todos []*Todo
	if result := db.Find(&todos); result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func GetTodo(id int) (*Todo, error) {
	var todo *Todo
	if result := db.First(&todo, id); result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}

func CreateTodo(todo *Todo) (*Todo, error) {
	if result := db.Create(&todo); result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}

func UpdateTodo(todo *Todo) (*Todo, error) {
	if result := db.Model(&todo).Updates(map[string]interface{}{"name": todo.Name, "done": todo.Done}); result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}

func DeleteTodo(id int) error {
	var todo *Todo
	if result := db.Delete(&todo, id); result.Error != nil {
		return result.Error
	}

	return nil
}
```

それぞれの操作は `result.Error` を返すのでそれを利用してハンドリングします。

Updateに関しては

https://gorm.io/docs/update.html#Updates-multiple-columns

> NOTE When update with struct, GORM will only update non-zero fields,
> you might want to use map to update attributes or use Select to specify fields to update

ゼロ値は処理されないが、ゼロ値で更新したい場合はmapを利用する。と書かれてあるのでmap使って更新しています。(done=false の場合は boolのゼロ値なのでTODOリストを未完了に戻せない。)

### モデルを利用

モデルを利用したAPIの動作を定義していきます。

`routers/api/todo.go`

`GET /api/todos`

```go
package api

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos, err := models.GetTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, todos)
}
```

一覧取得ができるか以下で試してみましょう。

```sh
curl --request GET --url http://localhost:8080/api/todos
```

次は詳細取得です。

`GET /api/todos/:id`

```go
type GetTodoRequest struct {
	ID int `uri:"id" json:"id" binding:"required"`
}

func GetTodo(c *gin.Context) {
	req := GetTodoRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := models.GetTodo(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, todo)
}
```

詳細取得ができるか以下で試してみましょう。

```sh
curl --request GET --url http://localhost:8080/api/todos/1
```

次は作成です。

`POST /api/todos`

```go
type CreateTodoRequest struct {
	Name string `form:"name" binding:"required"`
	Done *bool  `form:"done" binding:"required"`
}

func CreateTodo(c *gin.Context) {
	req := CreateTodoRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := models.CreateTodo(&models.Todo{
		Name: req.Name,
		Done: *req.Done,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, todo)
}
```

作成できるか以下で試してみましょう。

```sh
curl --request POST \
  --url http://localhost:8080/api/todos \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "test",
	"done": false
}'
```

次は更新です。

`PATCH /api/todos/:id`

```go
type UpdateTodoRequest struct {
	ID   int     `uri:"id" binding:"required"`
	Name *string `form:"name"`
	Done *bool   `form:"done"`
}

func UpdateTodo(c *gin.Context) {
	req := UpdateTodoRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := models.UpdateTodo(&models.Todo{
		ID:   req.ID,
		Name: *req.Name,
		Done: *req.Done,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, todo)
}
```

更新できるか以下で試してみましょう。

```sh
curl --request PATCH \
  --url http://localhost:8080/api/todos/1 \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "changed",
	"done": true
}'
```

次は削除です。

`DELETE /api/todos/:id`

```go
type DeleteTodoRequest struct {
	ID int `uri:"id" json:"id" binding:"required"`
}

func DeleteTodo(c *gin.Context) {
	req := DeleteTodoRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DeleteTodo(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.Status(http.StatusNoContent)
}
```

削除ができるか以下で試してみましょう。

```sh
curl --request DELETE --url http://localhost:8080/api/todos/1
```

これでTODOリストAPIの完成です。

### テスト

最後に簡単にテストについて解説します。

まずはテスト用データベースを作成しましょう。

```
$ mysql --host=127.0.0.1 --port=3306 --user=root --password=pass
mysql> create database `gin-sample-test` default character set utf8mb4 collate utf8mb4_bin;
```

マイグレーションも実行します。

```
$ export DATABASE_URL='mysql://root:pass@tcp(localhost:3306)/gin-sample-test'
$ migrate -database ${DATABASE_URL} -path migrations up
```

テスト用設定を用意します。

`/conf/test.ini`

```ini
[database]
Type = mysql
User = root
Password = pass
Host = localhost
Port = 3306
Name = gin-sample-test
```

以下簡単な例ですが、DBも含めたリクエストテストをTableDrivenTestsで実装しています。

`/routers/router_test.go`

```go
package routers

import (
	"fmt"
	"main/models"
	"main/pkg/setting"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: テストケース毎のテストデータ挿入や初期化

func TestRoutersSetupRouter(t *testing.T) {
	setting.Setup("../conf/test.ini")
	models.Setup()
	router := SetupRouter()

	tests := []struct {
		Title          string
		RequestMethod  string
		RequestPath    string
		ResponceStatus int
		ResponceBody   string
	}{
		{
			Title:          "データがない場合空配列を返す",
			RequestMethod:  "GET",
			RequestPath:    "/api/todos",
			ResponceStatus: 200,
			ResponceBody:   "[]",
		},
		{
			Title:          "指定IDのデータがない場合status:400でエラーメッセージを返す",
			RequestMethod:  "GET",
			RequestPath:    "/api/todos/1",
			ResponceStatus: 400,
			ResponceBody:   "{\"message\":\"record not found\"}",
		},
	}

	for _, td := range tests {
		td := td
		title := fmt.Sprintf(
			"SetupRouter %s %s %s",
			td.RequestMethod,
			td.RequestPath,
			td.Title,
		)
		t.Run(title, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(td.RequestMethod, td.RequestPath, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, td.ResponceStatus)
			assert.Equal(t, w.Body.String(), td.ResponceBody)
		})
	}
}
```

このテストは以下コマンドでrootディレクトリから実行できます。

```sh
$ go test -v ./...
```

簡単にテストを実装してみました。

テストケース毎のテストデータ挿入や初期化は別途実装する必要があったり[gomock](https://www.asobou.co.jp/blog/web/gomock)を利用したモックと関数のテストなどが必要かもしれません。プロジェクトに応じて調べてみて下さい。

コードの完全な例は `/gin-samples/06todoapp` にあります。

### 参考プロジェクト

- https://github.com/eddycjy/go-gin-example
- https://github.com/gothinkster/golang-gin-realworld-example-app
