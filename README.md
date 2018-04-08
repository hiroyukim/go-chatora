# go-chatora

It is a server that transparently caches rdbms.

## install

```sh
go get -u github.com/hiroyukim/go-chatora
```

## build

```sh
cd $GOPATH/src/github.com/hiroyukim/go-chatora
go build
```

## config

+ config.yaml
    + cache_type
        + leveldb
        + arc

```yaml
---
driver_name: sqlite3
data_source_name: ./go-chatora.db
debug_mode: false
```

## run

```sh
./go-chatora -c config.yaml
```

### response type

+ plain(default)
+ json
+ messagepack


## Author

hiroyukim
