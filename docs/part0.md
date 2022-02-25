# 環境構築

- goenv v2.0.0beta11
- go v1.17.7

## goenv

参考: https://github.com/syndbg/goenv/blob/master/INSTALL.md

brewなどでインストールもしくはリポジトリをcloneします。

```
$ brew install goenv
or
$ git clone https://github.com/syndbg/goenv.git ~/.goenv
```

shellに追加して有効にします。

`~/.bash_profile` などに追加
```
export GOENV_ROOT=$HOME/.goenv
export PATH=$GOENV_ROOT/bin:$PATH
eval "$(goenv init -)"
```

インストールリストに目的のバージョンがあるか確認してください。

```
$ goenv --version
goenv 2.0.0beta11
$ goenv install --list
1.17.7
```

1.17.7 がなければアップデートしてください。

```
$ cd ~/.goenv && git pull && cd -
```

go v1.17.7をインストールします。

```
$ goenv install 1.17.7
$ goenv global 1.17.7
$ go version
go version go1.17.7 darwin/amd64
```
