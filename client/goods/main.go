package main

import (
	"context"
	"fmt"
	"goods/proto/goodsService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1、连接服务器
	/*
	   credentials.NewClientTLSFromFile ：从输入的证书文件中为客户端构造TLS凭证。
	   grpc.WithTransportCredentials ：配置连接级别的安全凭证（例如，TLS/SSL），返回一个
	   DialOption，用于连接服务器。
	*/
	grpcClient, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
	}
	//2、注册客户端
	client := goodsService.NewGoodsClient(grpcClient)
	//增加
	res1, _ := client.AddGoods(context.Background(), &goodsService.AddGoodsReq{
		Goods: &goodsService.GoodsModel{
			Title:   "测试商品",
			Price:   20,
			Content: "测试商品的内容",
		},
	})
	fmt.Println(res1.Message)
	fmt.Println(res1.Success)

	//获取商品数据
	res2, _ := client.GetGoods(context.Background(), &goodsService.GetGoodsReq{})
	//fmt.Printf("%#v", res2.GoodsList)

	for i := 0; i < len(res2.GoodsList); i++ {
		fmt.Printf("%#v\r\n", res2.GoodsList[i])
	}
}
