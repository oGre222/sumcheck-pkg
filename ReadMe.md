# sumcheck-pkg
# sumcheck问题
1. golang 1.12
```
sumcheck hash: 7uALKaS/Nxv2O2vSwrS0+3KybL58Y2rWyYaRKgl/YOw=

go run main.go 
===>
err: strconv.ParseFloat: parsing "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": invalid syntax 
float: 0

实际使用commit：https://github.com/weibreeze/breeze-go/commit/745bbab56a31ed514662fa23c7748999c206bafa (v0.1.0)
```
2. golang 1.13+
```
sumcheck hash: +++pO6wSn6jNsX6iNnSYjiQsSVPWGD4O/iA0yzGw7s8=

go run main.go 
===>
err: <nil> 
float: 0

实际使用commit：https://github.com/weibreeze/breeze-go/commit/16003c2df404b1b8ebfaf619dfc7b0eebcd67b88
```
# go-checksum
```
go get github.com/vikyd/go-checksum

$GOPATH/bin/go-checksum /path/pkg github.com/weibreeze/breeze-go@v0.1.0
```
