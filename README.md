# ntop
nginxのstubモジュールで得られる情報を表示して流すcliツール

```
2021-09-10 03:18:18 +0900 JST
Active connections: 1
server accepts handled requests
 8 8 196
Reading: 0 Writing: 1 Waiting: 0

2021-09-10 03:18:19 +0900 JST
Active connections: 1
server accepts handled requests
 8 8 197
Reading: 0 Writing: 1 Waiting: 0
...
```

## install
`go install github.com/aokabi/ntop@latest`

## usage
`ntop --host localhost --port 8080` 
