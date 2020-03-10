package main

import (
	"fmt"
	"github.com/weibreeze/breeze-go"
)

func main() {
	var f float64
	b := breeze.NewBuffer(32)
	breeze.WriteStringType(b)
	b.WriteByte(byte(30))

	var data []byte
	for i:=0; i< 30 ;i++  {
		data = append(data, 'a')
	}

	b.Write(data)
	bErr := breeze.ReadFloat64(b, &f)
	fmt.Println("err:", bErr)
	fmt.Println("float:", f)
}
