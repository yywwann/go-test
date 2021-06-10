package main

import (
	"context"
	"example/etcd_test/naming"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func test1() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		getResp *clientv3.GetResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"172.16.101.107:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	// 写入
	//kv.Put(context.TODO(), "name1", "lesroad")
	//kv.Put(context.TODO(), "name2", "haha")

	// 读取name为前缀的所有key
	if getResp, err = kv.Get(context.TODO(), "/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	} else {
		// 获取成功
		fmt.Println(getResp.Kvs)
	}

	// 删除name为前缀的所有key
	//if _, err = kv.Delete(context.TODO(), "name", clientv3.WithPrevKV()); err != nil {
	//	fmt.Println(err)
	//	return
	//}
}

func test2() {
	if err := naming.Register("127.0.0.1:2379", "appid", "1.2.3.4:5", "otc", 15); err != nil {
		panic(err)
	}

}

func main() {
	//test2()
	test1()
	//time.Sleep(100 * time.Second)
}
