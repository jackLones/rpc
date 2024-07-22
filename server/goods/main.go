package main

import (
	"context"
	"fmt"
	"goods/proto/goodsService"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

//rpc远程调用的接口,需要实现goods.proto中定义的Goods服务接口,以及里面的方法
//1.定义远程调用的结构体和方法,这个结构体需要实现GoodsServer的接口

type Goods struct{}

//GoodsServer方法参考goods.pb.go中的接口
/*
// GoodsServer is the server API for Goods service.
type GoodsServer interface {
    // 通过rpc来指定远程调用的方法:
    // AddGoods方法:增加商品, 这个方法里面实现对传入的参数AddGoodsReq, 以及返回的参数AddGoodsRes进行约束
    AddGoods(context.Context, *AddGoodsReq) (*AddGoodsRes, error)
    // 获取商品列表: GetGoodsReq 参数可为空, 返回参数GetGoodsRes是一个商品相关的切片
    GetGoods(context.Context, *GetGoodsReq) (*GetGoodsRes, error)
}
*/
//增加商品数据
func (this Goods) AddGoods(c context.Context, req *goodsService.AddGoodsReq) (*goodsService.AddGoodsRes, error) {
	fmt.Println(req)
	//模拟返回操作,正式项目在这里进行数据库的操作即可,根据操作结果,返回相关数据
	return &goodsService.AddGoodsRes{
		Message: "增加成功" + req.Goods.Title, //需要获取商品title
		Success: true,
	}, nil
}

//获取商品列表
func (g Goods) GetGoods(c context.Context, req *goodsService.GetGoodsReq) (*goodsService.GetGoodsRes, error) {
	//  GoodsList []*GoodsModel
	var tempList []*goodsService.GoodsModel //定义返回的商品列表切片
	//模拟从数据库中获取商品的请求,循环结果,把商品相关数据放入tempList切片中
	for i := 0; i < 10; i++ {
		tempList = append(tempList, &goodsService.GoodsModel{
			Title:   "商品" + strconv.Itoa(i), // strconv.Itoa(i): 整型转字符串类型
			Price:   float64(i),               //float64(i): 强制转换整型为浮点型
			Content: "测试商品内容" + strconv.Itoa(i),
		})
	}
	return &goodsService.GetGoodsRes{
		GoodsList: tempList,
	}, nil
}

func main() {
	//1. 初始一个 grpc 对象
	grpcServer := grpc.NewServer()
	//2. 注册服务
	//helloService.RegisterGoodsServer(grpcServer, &Goods{})
	// &Hello{}和 new(Hello)相同
	goodsService.RegisterGoodsServer(grpcServer, new(Goods))
	//3. 设置监听， 指定 IP、port
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	// 4退出关闭监听
	defer listener.Close()
	//5、启动服务
	grpcServer.Serve(listener)
}
