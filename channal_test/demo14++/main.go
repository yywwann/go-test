package main

import (
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Print("顾客#", c, "进入酒吧")
	seat := <-bar // 需要一个位子来喝酒
	log.Print("++ customer#", c, " drinks at seat#", seat)
	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
	bar <- seat // 释放座位，离开酒吧
}

//将通道用做计数信号量（counting semaphore）
//缓冲通道可以被用做计数信号量。 计数信号量可以被视为多主锁。
//如果一个缓冲通道的容量为N，那么它可以被看作是一个在任何时刻最多可有N个主人的锁。
//上面提到的二元信号量是特殊的计数信号量，每个二元信号量在任一时刻最多只能有一个主人。
//
//计数信号量经常被使用于限制最大并发数。
//
//和将通道用做互斥锁一样，也有两种方式用来获取一个用做计数信号量的通道的一份所有权。
//通过发送操作来获取所有权，通过接收操作来释放所有权；
//通过接收操作来获取所有权，通过发送操作来释放所有权。
//下面是一个通过接收操作来获取所有权的例子：

func main() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10) // 此酒吧有10个座位
	// 摆放10个座位。
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId) // 均不会阻塞
	}

	// demo1 start
	//for customerId := 0; ; customerId++ {
	//	time.Sleep(time.Second)
	//	go bar24x7.ServeCustomer(customerId)
	//}
	// demo1 end

	// 在上例中，尽管在任一时刻同时在喝酒的顾客数不会超过座位数10，但是在某一时刻可能有多于10个顾客进入了酒吧，
	//  因为某些顾客在排队等位子。
	// 在上例中，每个顾客对应着一个协程。虽然协程的开销比系统线程小得多，但是如果协程的数量很多，则它们的总体开销还是不能忽略不计的。
	// 所以，最好当有空位的时候才创建顾客协程。
	// demo2 start
	//// 这个for循环和上例不一样。
	//for customerId := 0; ; customerId++ {
	//	time.Sleep(time.Second)
	//	seat := <- bar24x7 // 需要一个空位招待顾客
	//	go bar24x7.ServeCustomerAtSeat(customerId, seat)
	//}
	// demo2 end

	// 在上面这个修改后的例子中，在任一时刻最多只有10个顾客协程在运行（但是在程序的生命期内，仍旧会有大量的顾客协程不断被创建和销毁）。
	//
	// 在下面这个更加高效的实现中，在程序的生命期内最多只会有10个顾客协程被创建出来。
	// demo3 start
	consumers := make(chan int)
	for i := 0; i < cap(bar24x7); i++ {
		go bar24x7.ServeCustomerAtSeat2(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
	// demo3 end

	//} // 睡眠不属于阻塞状态
	//select {
	//
	//}
}

func (bar Bar) ServeCustomerAtSeat(c int, seat Seat) {
	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
	bar <- seat // 释放座位，离开酒吧
}

func (bar Bar) ServeCustomerAtSeat2(consumers chan int) {
	for c := range consumers {
		seatId := <-bar
		log.Print("++ 顾客#", c, "在第", seatId, "个座位开始饮酒")
		time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
		log.Print("-- 顾客#", c, "离开了第", seatId, "个座位")
		bar <- seatId // 释放座位，离开酒吧
	}
}
