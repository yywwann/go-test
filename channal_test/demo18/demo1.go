package main

import (
	"log"
	"math/rand"
	"time"
)

type Customer struct{ id int }
type Bar chan Customer

func (bar Bar) ServeCustomer(c Customer) {
	log.Print("++ 顾客#", c.id, "开始饮酒")
	time.Sleep(time.Second * time.Duration(3+rand.Intn(16)))
	log.Print("-- 顾客#", c.id, "离开酒吧")
	<-bar // 离开酒吧，腾出位子
}

func Demo1() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10) // 最对同时服务10位顾客
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumer := Customer{customerId}
		select {
		case bar24x7 <- consumer: // 试图进入此酒吧
			go bar24x7.ServeCustomer(consumer)
		default:
			log.Print("顾客#", customerId, "不愿等待而离去")
		}
	}
	select {}
}

//峰值限制（peak/burst limiting）
//将通道用做计数信号量用例和通道尝试（发送或者接收）操作结合起来可用实现峰值限制。 峰值限制的目的是防止过大的并发请求数。
//
//下面是对将通道用做计数信号量一节中的最后一个例子的简单修改，从而使得顾客不再等待而是离去或者寻找其它酒吧。
