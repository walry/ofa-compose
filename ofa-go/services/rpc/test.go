package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "ofa/protocol"
	"time"
)

//import (
//	"context"
//	test "ofa/protocol"
//)
//
//type TestService struct {
//
//}
//
//func (ts *TestService)Call(ctx context.Context,req *test.Request)(*test.Response,error){
//	return &test.Response{
//		Data: "ok",
//	},nil
//}

const (
	ADDRESS = "127.0.0.1:9999"
)

func Client(){
	cc,_ := grpc.Dial(ADDRESS,grpc.WithInsecure())
	defer cc.Close()
	client := pb.NewTestServiceClient(cc)

	ctx,e := context.WithTimeout(context.Background(),time.Second)
	fmt.Printf("ctx err:%+v\n",e)
	req := &pb.Request{
		Msg: "hello",
	}

	res,err := client.Call(ctx,req)
	fmt.Printf("recived:%+v\n",res)
	fmt.Printf("recived err:%+v\n",err)
}