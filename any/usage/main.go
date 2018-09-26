package main

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
	ptypes "github.com/golang/protobuf/ptypes"

	pb "github.com/lizebang/learning-proto3/any"
)

func main() {
	src := pb.Source{
		Data: []string{"000", "111", "222"},
	}

	any, err := ptypes.MarshalAny(&src)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	t := pb.Transport{
		Data: any,
	}

	// if you don't know src type
	fmt.Println(t.GetData().GetValue(), string(t.GetData().GetValue()))
	// but if you know src type, you can unmarshal it
	var dst pb.Destination
	err = proto.Unmarshal(t.GetData().GetValue(), &dst)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Println(dst.GetData())
}
