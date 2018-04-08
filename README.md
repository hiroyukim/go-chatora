# go-chatora

It is a server that transparently caches rdbms.


It is structured to transparently retrieve data from backend's rdbms. It is suitable for cases where management is troublesome when it is desired to perpetuate data with KVS or the like. Access by HTTP and get data with GET. With the `type` parameter, it returns data in the form plain, json, messagepack. There is an Onmemory `arc` and a Storage` leveldb` for caching structures.

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
    + driver_name
        + sqlite3
        + mysql
    + data_source_name
        + sqlite3: `/path/to/your.db'
        + mysql: `[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]`
    + table_name
        + string
    + lru_max_size
        + int
    + port
        + int
    + debug_mode
        + bool

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

### request

```
curl http://0.0.0.0:8080/?key=yourkey&type=json
```

### response type

+ plain(default)
+ json
+ msgp


## Author

hiroyukim
