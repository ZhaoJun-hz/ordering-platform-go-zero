## api 相关代码
```shell
goctl api go -api ./api/admin/doc/api/admin.api -dir ./api/admin/ -style goZero
```


## rpc 相关代码
```shell
goctl rpc protoc rpc/sys/sys.proto --go_out=./rpc/sys/ --go-grpc_out=./rpc/sys/ --zrpc_out=./rpc/sys/ -m -style goZero
```


## 生成model代码
```shell
goctl model mysql datasource -url="root:QWERtyui1234@tcp(10.166.66.14:3306)/ordering-platform" -table="sys*" -dir=./rpc/model/sysmodel -style goZero
```

## swagger
```shell
goctl api plugin -plugin goctl-swagger="swagger -filename admin.json" -api ./api/admin/doc/api/admin.api -dir ./docs

```