package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"jtthinkStudy/primary/grpcTest/pbfiles"
)

func main() {
	prod := &pbfiles.ProdModel{
		Id:   1,
		Name: "hello",
	}
	b, _ := proto.Marshal(prod)

	prod2:=&pbfiles.ProdModel{}
	proto.Unmarshal(b,prod2)
	fmt.Println(b,prod2)
}
