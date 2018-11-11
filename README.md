# gofmt-api-server

## Launch
```
go run cmd/fmtapi/main.go
```

## Test
```
curl localhost:8080/format -X POST --data-binary @sample_file.go
```
