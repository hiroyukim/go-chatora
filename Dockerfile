FROM alpine:latest

MAINTAINER hiroyukim

RUN apk --update add sqlite curl && rm -rf /var/cache/apk/*

RUN curl -L https://github.com/hiroyukim/go-chatora/releases/download/v0.0.1/go-chatora_0.0.1_linux_amd64.tar.gz \
        | tar zx -C /bin \
        && mv /bin/go-chatora_0.0.1_linux_amd64/go-chatora /bin/go-chatora

RUN sqlite3 go-chatora.db "CREATE TABLE key_values ( key TEXT NOT NULL UNIQUE, value TEXT NOT NULL);" \
        && sqlite3 go-chatora.db "INSERT INTO key_values (key,value) VALUES('9s','s')"

RUN echo -e "---\ndriver_name: sqlite3\ndata_source_name: /go/go-chatora.db" > /config.yml

CMD ["/bin/go-chatora", "-c", "/config.yml"]
