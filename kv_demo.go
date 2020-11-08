package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	// 建立客户端成功
	fmt.Println("connect succ")
	defer cli.Close()

	// 设置1秒超时，访问 etcd 有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// 设置 key
	_, err = cli.Put(ctx, "hello", "world")

	//操作完毕，取消 etcd
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	// 取值，设置超时为 1 秒
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "hello")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
