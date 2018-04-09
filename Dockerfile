FROM golang:latest

MAINTAINER hiroyukim

RUN apt-get update && apt-get install -y git sqlite3 libsqlite3-dev

RUN go get -u github.com/hiroyukim/go-chatora \
        && cd $GOPATH/src/github.com/hiroyukim/go-chatora \
        && go build

RUN sqlite3 go-chatora.db "CREATE TABLE key_values ( key TEXT NOT NULL UNIQUE, value TEXT NOT NULL);" \
        && sqlite3 go-chatora.db "INSERT INTO key_values (key,value) VALUES('9s','s')"

RUN echo "---\ndriver_name: sqlite3\ndata_source_name: /go/go-chatora.db" > /go/config.yml

CMD ["/go/bin/go-chatora", "-c", "/go/config.yml"]
