# redirect-server

Simple Golang server that leads a client through a redirect chain.

This was used to test the behavior of popular web browsers, which follow between 15 and 20 redirects before giving up.

## Build

```shell
go build
```

## Usage

First:
```shell
redirect-server $PORT
```
Then:
```
curl http://localhost:$PORT/redirect/10
```
