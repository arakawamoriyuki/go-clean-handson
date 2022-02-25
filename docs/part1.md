# go 入門

参考

- [はじめての Go](http://gihyo.jp/dev/feature/01/go_4beginners)
- [Go Web プログラミング](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/01.0.html)

```
$ go version
go version go1.17.7 darwin/amd64
```

## 概要

Go は、2009 年に Google により発表されたオープンソースのプログラミング言語です。

- シンプルな言語仕様
- Web、Windows、Mac、Linux、Android、iOS など環境に合わせた実行ファイルを生成できるクロスコンパイル
- 標準のコーディング規約や静的解析系ツール、テストツールなどがある
- 巨大なコードでも高速にコンパイルできて大規模開発にも適してる
- 並行処理のサポート
- 型やメソッドがあるが、継承はない。変わりに埋め込み型などがある
- 例外の扱いが特殊

### シンプル

繰り返し構文は for 文のみです。

また、三項演算子もありません。

表現のばらつきを抑えられ、同じような書き方になります。

### 危険の回避

意図せぬエラーの原因になる暗黙の型変換などは排除され、

使われていない変数、使われていないインポートはコンパイルエラーになります。

### 例外の排除

`try/catch` がなく、発生した例外を戻り値として呼び出し側に返すことでハンドリングします。

また、パニックとリカバという例外を制御する方法もあります。

### 他値を返す

複数の値を return できます。

### 平行処理

ゴルーチンという軽量スレッドを用いて並行処理ができます。

同時に実行されているゴルーチンの間ではチャネルという機能でデータをやりとりできます。

### 標準パッケージ

https://golang.org/pkg/

- `encoding/json` JSON から構造体、構造体から JSON に変換
- `os`, `io` ファイルの作成、書き込み、読み込みなど
- `io/ioutil` ファイルの操作のラッパー
- `net/http` HTTP サーバ作成 API
- `text/template` web に必要な html のテンプレートエンジン

### コマンド

https://golang.org/doc/cmd

|  コマンド  |                説明                 |
| :--------: | :---------------------------------: |
|  go build  |         プログラムのビルド          |
|   go run   |      プログラムのビルドと実行       |
|   go mod   |           パッケージ管理            |
| go install |  プログラムのビルドとインストール   |
|   go get   |        外部パッケージの取得         |
|   go fmt   | Go の規約に合わせてプログラムを整形 |
|   go fix   |       古いコードを書き換える        |
|   go vet   |              静的解析               |
|  go test   |     テストやベンチマークの実行      |
|   go doc   |    ソースからドキュメントの生成     |

---

## 動かしてみる

### インストール

`brew install go` や `goenv` 、もしくは[インストーラー](https://golang.org/doc/install)でインストールできます。

### Hello world

```main.go
package main

import "fmt"

func main(){
    fmt.Printf("Hello, world")
}
```

```
$ go run main.go
hello world
```

### ビルドしてバイナリを実行してみる

```
$ go build main.go
$ ls
main main.go
$ ./main
hello world
```

各環境で動くバイナリができます。

他の環境で build したい場合は `GOOS=linux` や `GOARCH=amd64` など環境変数で指定してビルドできます。

### コードフォーマット

標準のコーディング規約とフォーマッターが付属されていて、以下コマンドでフォーマットすることができます。

```
$ go fmt main.go
```

vscode では以下設定で保存時に自動フォーマットできます。

```
"[go]": {
    "editor.formatOnSave": true,
}
```

---

## プロジェクト構成

GOPATH モード と モジュール対応モード があります。

GOPATH モード は GOPATH 環境変数からの相対パスで構成します。

モジュール対応モード は任意のディレクトリで開発・ビルドが可能なモードです。

1.17 から標準なモジュール対応モードで説明します。

### サブモジュールと外部モジュール

まずはモジュール名を決めて `go mod init` します。

モジュール名は `{Repository FQDN}/{Repository Path}/{Module Path}` のような命名規則(例えば `github.com/arakawamoriyuki/go-clean-handson/test` など)にすると、誰でも利用できるようになります。

公開しない場合は `main` などでも問題ありません。

```sh
$ go mod init main
or
$ go mod init github.com/arakawamoriyuki/go-clean-handson/test
```

自作したサブモジュール、外部のモジュールを読み込んで使ってみましょう。

必要なパッケージがあれば以下のように追加することができます。

```sh
$ go get github.com/arakawamoriyuki/gosample
```

以下のような構成でサブパッケージと外部パッケージの利用例をみてみましょう。

```sh
$ tree
├── go.mod
├── go.sum
├── main.go
└── subpkg
    ├── other.go
    └── subpkg.go
```

```subpkg/subpkg.go
package subpkg

var Message string = "subpkg/subpkg.go message"
```

```subpkg/other.go
package subpkg

var OtherMessage string = "subpkg/other.go message"
```

```main.go
package main

import (
	"fmt"

	"main/subpkg"

	"github.com/arakawamoriyuki/gosample"
)

func main() {
	fmt.Println(subpkg.Message)
	fmt.Println(subpkg.OtherMessage)
	fmt.Println(gosample.Message)
}
```

早速実行してみます。

```sh
$ go run main.go
```

同一パッケージの関数の場合はファイル名関係なく、関数名や変数名のみで呼べ出せます。

ファイル名は関係なく、 `subpkg/other.go` は `subpkg` パッケージなので呼び出せていることにも注目しましょう。

参考: https://dev.classmethod.jp/articles/go-setup-and-sample/

また、以下でビルドすることができ、そのディレクトリに実行ファイルを作成できます。

```
$ go build
$ ./main
```

### 命名規則

- ファイル名
  - スネークケース `snake_case.go`
- ディレクトリ名
  - 小文字英字のみ `sample/`
  - 必要ならケバブケース `sample-module/`
- メソッド
  - エクスポートする関数やメソッド(public)はアッパーキャメルケース `UpperCamelCase`
  - ファイル内でしか利用しない関数やメソッド(private)はローワーキャメルケース `lowerCamelCase`
- レシーバー名
  - 省略した英語 1 文字か 2 文字など短くつける多い HttpClient は `c` など
- 変数、引数名
  - 英語 1 文字か 2 文字でも良いが、可読性を考えて長くても良い `httpClient` など
- 参考
  - https://zenn.dev/keitakn/articles/go-naming-rules
  - https://github.com/kubernetes/kubernetes/tree/master/api

---

## 言語仕様

### パッケージ宣言

```sample.go
package main
```

`go run` では main パッケージの main 関数が実行されます。

サブモジュールを作成する場合はディレクトリ名に合わせてつけてください。

### インポート

外部モジュールはリポジトリ含めた絶対パス、自作のサブモジュールなどはモジュール名に合わせて変更します。

区切りにカンマは必要ありません。

```sample.go
import (
    # 組み込み関数
    "fmt"

    # サブモジュール (モジュール名がmainの場合)
	"main/subpkg"

    # サブモジュール (モジュール名がリポジトリ含めた絶対パスの場合)
	"github.com/arakawamoriyuki/go-clean-handson/test/subpkg"

    # 外部モジュール
    "github.com/arakawamoriyuki/gosample"
)
```

インポートにはいくつかのオプションがあります。

- オプションなしの場合はモジュール名を指定して使える `fmt.Println()`
- `文字列` はエイリアスになる( `fmt.Println()` が `f.Println()` )
- `.` は中の関数が展開される ( `strings.ToUpper()` が `ToUpper()` )
- `_` は使用していないパッケージだと明示する (使用してなくてもコンパイルエラーにならない)

```sample.go
import (
    "fmt"
    f "fmt"
    . "strings"
    _ "github.com/wdpress/gosample"
)
```

### 型

|     型     |             説明              |
| :--------: | :---------------------------: |
|   uint8    |     8 ビット符号なし整数      |
|   uint16   |     16 ビット符号なし整数     |
|   uint32   |     32 ビット符号なし整数     |
|   uint64   |     64 ビット符号なし整数     |
|    int8    |     8 ビット符号あり整数      |
|   int16    |     16 ビット符号あり整数     |
|   int32    |     32 ビット符号あり整数     |
|   int64    |     64 ビット符号あり整数     |
|  float32   |       32 ビット浮動小数       |
|  float64   |       64 ビット浮動小数       |
| complex64  |        64 ビット複素数        |
| complex128 |       128 ビット複素数        |
|    byte    |      uint8 のエイリアス       |
|    rune    |   Unicode のコードポイント    |
|    uint    | 32 か 64 ビットの符号なし整数 |
|    int     | 32 か 64 ビットの符号あり整数 |
|  uintptr   |   ポインタ値用符号なし整数    |
|   error    | エラーを表わすインタフェース  |

文字列はダブルクォートで書くことができます。

```sample.go
var Message string = "hello world"
```

ヒアドキュメントはバッククォートで書くことができます。

```sample.go
var Message string = `first line
second line
third line`
```

また、コンパイラが明らかにわかる型を推論して型宣言を省略することもできます。

```sample.go
message := "hello world"
```

### 変数

変数は `var 変数名 型 = 値` のような形式で定義できます。

また、複数定義することもでき、同じ型なら省略可能です。

```sample.go
var message string = "hello world"

var foo, bar, buz string = "foo", "bar", "buz"
var (
    a string = "aaa"
    b = "bbb"
    c = "ccc"
    d = "ddd"
    e = "eee"
)
```

### 定数

定数は const で定義でき、再代入不可になります。

```sample.go
const Hello string = "hello"
Hello = "bye" // cannot assign to Hello
```

### ゼロ値

代入しない場合、変数はゼロ値で初期化されます。

```sample.go
var i int // 整数型のゼロ値 0 になる
```

|      型      |            ゼロ値            |
| :----------: | :--------------------------: |
|    整数型    |              0               |
| 浮動小数点型 |             0.0              |
|     bool     |            false             |
|    string    |              ""              |
|     配列     |     各要素がゼロ値の配列     |
|    構造体    | 各フィールドがゼロ値の構造体 |
| そのほかの型 |             nil              |

### if

条件に丸括弧は必要ありません。また、三項演算子はありません。

```sample.go
a, b := 10, 100
if a > b {
    fmt.Println("a is larger than b")
} else if a < b {
    fmt.Println("a is smaller than b")
} else {
    fmt.Println("a equals b")
}
```

### for

for を利用してループを回すことができます。

```sample.go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

while のような書き方も可能です。

```sample.go
n := 0
for n < 10 {
    fmt.Printf("n = %d\n", n)
    n++
}
```

また、以下のように無限ループを作ることもできます。

```sample.go
for {
    doSomething()
}
```

ループを終了する `break` 、スキップする `continue` なども利用可能です。

配列は range を使ってループすることができます。

```sample.go
for i, v := range []string{"a", "b", "c"} {
    fmt.Println(i, v)
}
```

### switch

以下のようにカンマで区切った複数の値も指定可能です。

条件に一致した処理を走らせることができます。

```sample.go
n := 10
switch n {
case 15:
    fmt.Println("FizzBuzz")
case 5, 10:
    fmt.Println("Buzz")
case 3, 6, 9:
    fmt.Println("Fizz")
default:
    fmt.Println(n)
}
```

言語によっては `break` がない場合次の case も評価される言語もありますが、

golang では 1 つの case 実行されると次の case に移ることはありません。

`fallthrough` キーワードで次にうつる事もできます。

```sample.go
n := 3
switch n {
case 3:
    n = n - 1
    fallthrough
case 2:
    n = n - 1
    fallthrough
case 1:
    n = n - 1
    fmt.Println(n) // 0
}
```

### 関数

関数は `func` で作ります。引数には型を指定、複数同じ型なら一つにまとめることもできます。

```sample.go
func sum(i, j int) {
    fmt.Println(i + j)
}
```

また、戻り値は関数定義のあとに型を指定します。

```sample.go
func sum(i, j int) int {
    return i + j
}
```

複数値を返す場合、複数の型を指定します。

```sample.go
func swap(i, j int) (int, int) {
    return j, i
}
```

名前付き戻り値で `return` の後の値を省略することもできます。(代入された値を返す。代入されていなければ結果的にゼロ値を返す)

```sample.go
func div(i, j int) (result int, err error) {
    if j == 0 {
        err = errors.New("divied by zero")
        return // return 0, errと同じ
    }
    result = i / j
    return // return result, nilと同じ
}
```

以下のように無名関数を作ることもできます。

```sample.go
func(i, j int) {
    fmt.Println(i + j)
}(2, 4)
```

関数を変数に代入できます。

```sample.go
var sum func(i, j int) = func(i, j int) {
    fmt.Println(i + j)
}
```

可変長引数も利用できます。

```sample.go
func sum(nums ...int) (result int) {
    // numsは[]int型
    for _, n := range nums {
        result += n
    }
    return
}

func main() {
    fmt.Println(sum(1, 2, 3, 4))  // 10
}
```

---

### エラー

go は `try/catch` や `throw` がありません。

変わりにエラーを戻り値で返すことでハンドリングします。

また、エラーの作成は errors パッケージを使います。

```sample.go
package main

import (
    "errors"
    "fmt"
    "log"
)

func div(i, j int) (int, error) {
    if j == 0 {
        // 自作のエラーを返す
        return 0, errors.New("divied by zero")
    }
    return i / j, nil
}

func main() {
    n, err := div(10, 0)
    if err != nil {
        // エラーを出力しプログラムを終了する
        log.Fatal(err)
    }
    fmt.Println(n)
}
```

複数の値を返す場合はエラーを最後にする慣習があります。

## 配列

配列は固定長で長さを指定します。

```sample.go
var arr1 [4]string
```

`[...]` で暗黙的に長さの指定ができます。

```sample.go
arr := [...]string{"a", "b", "c", "d"}
```

引数で受け取る場合にも型と長さの指定をする必要があります。

```sample.go
func fn(arr [4]string) {
    fmt.Println(arr)
}

func main() {
    var arr1 [4]string
    var arr2 [5]string

    fn(arr1) // ok
    fn(arr2) // コンパイルエラー
}
```

スライスという可変長配列も定義できます。

```sample.go
var s []string
s := []string{"a", "b", "c", "d"}
```

値を部分的に切り出す事ができます。

```sample.go
s := []int{0, 1, 2, 3, 4, 5}
fmt.Println(s[2:4])      // [2 3]
fmt.Println(s[0:len(s)]) // [0 1 2 3 4 5]
fmt.Println(s[:3])       // [0 1 2]
fmt.Println(s[3:])       // [3 4 5]
fmt.Println(s[:])        // [0 1 2 3 4 5]
```

### append

`append` はスライスの末尾に値を追加し、その結果を返す組込み関数です。

```sample.go
s1 := []string{"a", "b"}
s1 = append(s1, "c") // s1にs2を追加
fmt.Println(s1)      // [a b c]
```

また、可変長の値を受け取ることもできます。

```sample.go
s1 := []string{"a", "b"}
s2 := []string{"c", "d"}
s1 = append(s1, s2...) // s1にs2を追加
fmt.Println(s1)        // [a b c d]
```

### range

添字によるアクセスの代わりに `range` を使用できます。

```sample.go
s1 := []string{"a", "b", "c", "d"}

for index, value := range s1 {
    // index = 添字, value = 値
    fmt.Println(index, value)
    // 0 a
    // 1 b
    // 2 c
    // 3 d
}
```

また、map 型もループで回せます。

```sample.go
months := map[string]int{
    "January": 1,
    "February": 2,
}

for key, value := range months {
    fmt.Println(key, value)
    // January 1
    // February 2
}
```

## マップ

`string` のキーに `int` の値を格納するマップ

```sample.go
months := map[string]int{
    "January": 1,
    "February": 2,
}
```

キーの存在確認は以下のように判定できます。

```sample.go
_, ok := months["January"]
if ok {
    // データがあった場合
}
```

マップからデータを消す場合は delete を使います。

```sample.go
delete(months, "January")
```

## ポインタ

Go はポインタを扱うことができます。

アンパサンド(&)はアドレス演算子。値からポインタへの変換を行う。

アスタリスク(\*)は間接参照演算子。ポインタから値への変換を行う。

int などでも参照渡しできるし配列を値渡しもできます。

```sample.go
func callByValue(i int) {
    i = 20 // 代入しても呼び出し側へ影響しない
}

func callByRef(i *int) {
    *i = 20 // 代入すると呼び出し側の変数も変わる
}

func main() {
    var i int = 10
    callByValue(i) // 値を渡す
    fmt.Println(i) // 10
    callByRef(&i) // アドレスを渡す
    fmt.Println(i) // 20
}
```

引数受ける側も関数利用側もポインタ渡す/受けるなら `*` や `&` つけないとコンパイルエラーになります。

## defer

ファイル開いたり、データベースにコネクション貼った場合など、エラーが起きても起きなくても実行して欲しい処理に使います。

`defer` は延期という意味。他の言語でいう `finaly` のような使い方をします。

```sample.go
func main() {
    file, err := os.Open("./error.go")
    if err != nil {
        // エラー処理
    }
    // 関数を抜ける前に必ず実行される
    defer file.Close()
    // 正常処理
}
```

## パニック

エラーは戻り値によって表現するのが基本ですが、

配列やスライスの範囲外にアクセスした場合や、ゼロ除算をしてしまった場合などはエラーを返せません。

この状態をパニックという。

パニックで発生したエラーは `recover` で拾えるので、defer で処理する事でエラー処理ができます。

```sample.go
func main() {
    defer func() {
        err := recover()
        if err != nil {
            // runtime error: index out of range
            log.Fatal(err)
        }
    }()

    a := []int{1, 2, 3}
    fmt.Println(a[10]) // パニックが発生
}
```

パニックは自分で起こす事もできます。

```sample.go
a := []int{1, 2, 3}
for i := 0; i < 10; i++ {
    if i >= len(a) {
        panic(errors.New("index out of range"))
    }
    fmt.Println(a[i])
}
```

---

## type

`type` キーワードを使って独自の型を作ることができます。

一例を見てみましょう。

以下の関数は int 型の `id` と `priority` を受け取ります。

```sample.go
func ProcessTask(id, priority int) {
}
```

同じ `int` 型なので

引数の順番間違えてもコンパイルが通り、間違えやすいインターフェースになっています。

```sample.go
var id int = 3
var priority int = 5
ProcessTask(id, priority)
ProcessTask(priority, id) // 順番間違えてもコンパイル通る
```

場合に応じて独自の型を定義すると間違いが減り、安全になります。

```sample.go
type ID int
type Priority int

func ProcessTask(id ID, priority Priority) {
}

var id ID = 3
var priority Priority = 5
ProcessTask(priority, id) // コンパイルエラー
```

## 構造体（struct）

構造体を独自の型として宣言できます。

また、構造体のプロパティはドットでアクセス可能です。

```sample.go
type Task struct {
    ID int
    Detail string
    Done bool
}

var task Task = Task{
    ID: 1,
    Detail: "buy the milk",
    Done: true,
}
fmt.Println(task.ID) // 1
fmt.Println(task.Detail) // "buy the milk"
fmt.Println(task.Done) // true
```

## ポインタ型

```sample.go
var task Task = Task{} // Task型
var task *Task = &Task{} // Taskのポインタ型
```

ポインタ型ではない型は値渡しされます。

```sample.go
type Task struct {
    ID int
    Detail string
    Done bool
}

func Finish(task Task) {
    task.Done = true
}

func main() {
    task := Task{Done: false}
    Finish(task)
    fmt.Println(task.Done) // falseのまま
}
```

ポインタ型は参照渡しされます。

```sample.go
func Finish(task *Task) {
    task.Done = true
}

func main() {
    task := &Task{Done: false}
    Finish(task)
    fmt.Println(task.Done) // true
}
```

## new

構造体は組み込み関数 `new` を使い、ゼロ値で初期化できます。

```sample.go
type Task struct {
    ID int
    Detail string
    Done bool
}

task := new(Task)
fmt.Println(task.ID == 0) // true
fmt.Println(task.Detail == "") // true
fmt.Println(task.Done == false) // true
```

## コンストラクタ

Go にはコンストラクタにあたる構文がありません。

代わりに New で始まる関数を定義し、その内部で構造体を生成するのが通例です。

```sample.go
func NewTask(id int, detail string) *Task {
    task := &Task{
        ID: id,
        Detail: detail,
        Done: false,
    }
    return task
}

func main() {
    task := NewTask(1, "buy the milk")
    // &{ID:1 Detail:buy the milk Done:false}
    fmt.Printf("%+v", task)
}
```

## メソッド

メソッドはメソッド名の前に定義したい型を指定します。

```sample.go
func (変数名 メソッドを定義したい型) メソッド名() 戻り値の型 {
}
```

```sample.go
package main

import (
	"fmt"
)

type Task struct {
    ID int
    Detail string
    Done bool
}

func NewTask(id int, detail string) *Task {
    task := &Task{
        ID: id,
        Detail: detail,
        Done: false,
    }
    return task
}

// taskの文字列表現を返す
func (task Task) String() string {
    str := fmt.Sprintf("%d) %s %t", task.ID, task.Detail, task.Done)
    return str
}

// taskを完了する
func (task *Task) Finish() {
    task.Done = true
}

func main() {
    task := NewTask(1, "buy the milk")
	fmt.Println(task.String()) // 1) buy the milk false
	task.Finish()
	fmt.Println(task.String()) // 1) buy the milk true
}
```

## インターフェース

`type インターフェース名 interface {}` でメソッドが定義されている事を強制する用途に利用します。

```sample.go
package main

import (
	"fmt"
)

type Task struct {
	ID     int
	Detail string
	Done   bool
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		Done:   false,
	}
	return task
}

func (task Task) String() string {
	str := fmt.Sprintf("%d) %s %t", task.ID, task.Detail, task.Done)
	return str
}

// Stringerインターフェースを定義
type Stringer interface {
	String() string
}

// この関数には.String()メソッドを実装しているオブジェクトを渡せる。
func print(stringer Stringer) {
	fmt.Println(stringer.String())
}

func main() {
	task := NewTask(1, "buy the milk")

	// TaskはString()実装しているので渡せる
	print(task)

	// Stringer型の変数iを定義できる
	var s Stringer
	s = task
	fmt.Println(s)

	// Stringer型の配列に入れることもできる
	stringers := []Stringer{
		task,
		NewTask(2, "buy the banana"),
		NewTask(3, "buy the pan"),
	}
	fmt.Println(stringers)
}
```

実際は別の型であっても、インターフェースの指定の通りの実装されていれば置き換えることができます。

### interface{}

すべての引数を受け付ける Any 型が作れる

```sample.go
type Any interface {
}

func Do(any Any) {
  // do something
}
```

書き方自体以下と同じです。

```sample.go
func Do(any interface{}) {
  // do something
}
```

Tips: v1.18 からジェネリクスが取り入れられる関係で、any も標準の識別子として利用可能になります。

## 型の埋め込み

Go では，継承はサポートされていません。

代わりにほかの型を「埋め込む」(Embed) という方式で構造体やインタフェースの振る舞いを拡張できます。

以下は Task に Reminder を埋め込むを埋め込み、Task が Reminder のプロパティやメソッドが利用可能になる例です。

```sample.go
package main

import (
	"fmt"
	"time"
)

type Reminder struct {
	ExpiredAt time.Time
}

func (r *Reminder) RemainingTime() time.Duration {
	now := time.Now()
	return r.ExpiredAt.Sub(now)
}

func NewReminder(minute int) *Reminder {
	return &Reminder{
		ExpiredAt: time.Now().Add(time.Duration(minute) * time.Minute),
	}
}

type Task struct {
	ID        int
	Detail    string
	Done      bool
	*Reminder // Reminderを埋め込む(Embed)
}

func NewTask(id int, detail, firstName, lastName string) *Task {
	task := &Task{
		ID:       id,
		Detail:   detail,
		Done:     false,
		Reminder: NewReminder(10),
	}
	return task
}

func main() {
	task := NewTask(1, "buy the milk", "Jxck", "Daniel")

	// TaskにReminderのプロパティやメソッドが埋め込まれている
	fmt.Println(task.ExpiredAt)
	fmt.Println(task.RemainingTime())
}
```

## インタフェースの埋め込み

構造体( `struct` )だけではなくインターフェース( `interface` )も「埋め込む」(Embed) 事ができます

```sample.go
// 読み込みメソッドがある事をインターフェースで宣言
type Reader interface {
    Read(p []byte) (n int, err error)
}

// 書き込みメソッドがある事をインターフェースで宣言
type Writer interface {
    Write(p []byte) (n int, err error)
}

// 読み込みメソッドも書き込みメソッドもある事をインターフェースで宣言
type ReadWriter interface {
    Reader
    Writer
}
```

## 型変換 (キャスト)

明示的に型変換(キャスト)ができます。

```sample.go
var s string = "abc"
var b []byte = []byte(s) // string -> []byte
fmt.Println(b)           // [97 98 99]
```

キャストに失敗した場合はパニックが発生します。

```sample.go
// cannot convert "a" (type string) to type int
a := int("a")
```

## 型の検査

Type Assertion で型を調べる事ができます。

第一戻り値が元の値、第二戻り値が調べた結果です。

```sample.go
s, ok := value.(string) // Type Assertion
if ok {
    fmt.Printf("value is string: %s\n", s)
} else {
    fmt.Printf("value is not string\n")
}
```

Type Switch で型で分岐処理ができます。

```sample.go
switch v := value.(type) {
case string:
    fmt.Printf("value is string: %s\n", v)
case int:
    fmt.Printf("value is int: %d\n", v)
case Stringer:
    fmt.Printf("value is Stringer: %s\n", v)
}
```

---

## ゴルーチン

軽量スレッド、ゴルーチンを利用して非同期処理できる

以下はリクエストを 3 回送るコードです。

同期処理の場合は、1 度目のリクエスト完了後に 2 度目,3 度目...と直列で続きます。

```sample.go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    urls := []string{
        "https://google.com",
        "https://google.com",
        "https://google.com",
    }
    for _, url := range urls {
        res, err := http.Get(url)
        if err != nil {
            log.Fatal(err)
        }
        defer res.Body.Close()
        fmt.Println(url, res.Status)
    }
}
```

`go` キーワードで非同期処理で実装した場合は、並行してリクエストを行うことができます。

```sample.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

func main() {
    wait := new(sync.WaitGroup)
    urls := []string{
        "https://google.com",
        "https://google.com",
        "https://google.com",
    }
    for _, url := range urls {
        // waitGroupに追加
        wait.Add(1)
        // 取得処理をゴルーチンで実行する
        go func(url string) {
            res, err := http.Get(url)
            if err != nil {
                log.Fatal(err)
            }
            defer res.Body.Close()
            fmt.Println(url, res.Status)
            // waitGroupから削除
            wait.Done()
        }(url)
    }
    // 待ち合わせ
    wait.Wait()
}
```

## チャネル

複数のゴルーチンのデータをやりとりしたい場合チャネルを利用することができます。。

まずは組み込み関数 `make` を使い、以下のようにチャネルを作成、書き込み、読み込みができます。

```sample.go
// stringを扱うチャネルを生成
ch := make(chan string)

// チャネルにstringを書き込む
ch <- "a"

// チャネルからstringを読み出す
message := <- ch
```

チャネルを利用して並行して HTTP リクエストし、早く取得されたステータスから順に受け取って処理しておく事ができます。

```sample.go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    urls := []string{
        "https://google.com",
        "https://google.com",
        "https://google.com",
    }
    statusChan := make(chan string)
    for _, url := range urls {
        // 取得処理をゴルーチンで実行する
        go func(url string) {
            res, err := http.Get(url)
            if err != nil {
                log.Fatal(err)
            }
            defer res.Body.Close()
            statusChan <- res.Status
        }(url)
    }
    for i := 0; i < len(urls); i++ {
        // waitする必要がなく、先に受け取ったチャネルから処理できる
        fmt.Println(<-statusChan)
    }
}
```

## select 文を用いたイベント制御

読み出しや書き込みイベント制御のための select 文もあります。

主な用途は for/select 文と break を用いて実装するタイムアウト処理などに利用されます。

```sample.go
ch1 := make(chan string)
ch2 := make(chan string)
for {
    select {
    case c1 := <-ch1:
        // ch1からデータを読み出したときに実行される
    case c2 := <-ch2:
        // ch2からデータを読み出したときに実行される
    case ch2 <- "c":
        // ch2にデータを書き込んだときに実行される
    }
}
```

また、`make` 関数には同時に 3 つまで読み出されないかぎりチャネルに書き込まない制御をするバッファ付きチャネルや、

同時起動数制御などメッセージキューのような動作をしてくれるバッファ指定引数があります。

---

## サーバーを立ててみよう

`httprouter` を使って簡単なサーバーを立ててみましょう。

```
$ go mod init main
$ go get github.com/julienschmidt/httprouter@latest
```

`main.go` を作成します。

```main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello golang")
}

func FizzBuzz(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	num := p.ByName("num")

	// TODO: numをint型に変換
	// TODO: 3の倍数の時にFizz
	// TODO: 5の倍数の時にBuzz
	// TODO: 15の倍数の時にFizzBuzz
	// TODO: それ以外の場合に数値をそのまま表示

	fmt.Fprint(w, num)
}

func main() {
	router := httprouter.New()

	router.GET("/", Hello)
	router.GET("/fizzbuzz/:num", FizzBuzz)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

```

API サーバーを実行します。

```
$ go run main.go
```

コードを確認しながら試して見ましょう。

```
$ curl http://localhost:8080/
Hello golang
```

`/fizzbuzz/:num` で FizzBuzz API を作って見ましょう。

- `strconv.Atoi` を使い、文字列を数値型にキャストできます。 https://xn--go-hh0g6u.com/pkg/strconv/
- `fmt.Fprint` に `http.ResponseWriter` を渡すことで body に書き込みが行えます。 https://github.com/julienschmidt/httprouter

```
$ curl http://localhost:8080/fizzbuzz/1
1
$ curl http://localhost:8080/fizzbuzz/2
2
$ curl http://localhost:8080/fizzbuzz/3
Fizz!
$ curl http://localhost:8080/fizzbuzz/4
4
$ curl http://localhost:8080/fizzbuzz/5
5
$ curl http://localhost:8080/fizzbuzz/6
Buzz!
$ curl http://localhost:8080/fizzbuzz/7
7
$ curl http://localhost:8080/fizzbuzz/15
FizzBuzz!
```
