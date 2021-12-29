package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Worker struct {
	Mutex    sync.Mutex
	id       int
	timeout  time.Duration
	url      string
	openPort []int64
	done     int
	chanData chan int64
}

func NewWorker(id int, timeout time.Duration, url string) *Worker {
	return &Worker{
		Mutex:    sync.Mutex{},
		id:       id,
		timeout:  timeout,
		url:      url,
		openPort: make([]int64, 0),
		done:     0,
		chanData: make(chan int64, 1024),
	}
}

func (w *Worker) ScanUrl(ch chan int64) {
	for port := range w.chanData {
		//fmt.Println("start ", w.id, port)
		//time.Sleep(10 * time.Second)
		opened, _ := isOpen(fmt.Sprintf("%s:%d", w.url, port), w.timeout)
		if opened {
			fmt.Println(w.id, port)
			w.openPort = append(w.openPort, port)
		}
		w.done++
	}
	wg.Done()
}

func isOpen(url string, timeout time.Duration) (bool, error) {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.Dial("tcp", url)
	if err == nil {
		_ = conn.Close()
		return true, nil
	}

	return false, err
}

var url = "172.16.101.107"

var workerNum = 250
var wg sync.WaitGroup

func main() {
	timeout := time.Millisecond * 200

	production := make(chan int64)

	workers := make([]*Worker, 0)
	for i := 0; i < workerNum; i++ {
		workers = append(workers, NewWorker(i, timeout, url))
	}

	startTime := time.Now()
	for _, worker := range workers {
		wg.Add(1)
		go worker.ScanUrl(production)
	}

	go func() {
		for port := 0; port < 30000; port++ {
			workers[port%workerNum].chanData <- int64(port)
		}

		for _, worker := range workers {
			close(worker.chanData)
		}
		fmt.Println("done")
	}()

	wg.Wait()
	cost := time.Since(startTime)
	openPorts := make([]int64, 0)
	for _, worker := range workers {
		openPorts = append(openPorts, worker.openPort...)
	}
	fmt.Println(cost)
	fmt.Println(openPorts)
	fmt.Println(len(openPorts))
}

// without timeout
// 10 		27.011663986s	44
// 100 		7.007735893s 	44
// 1000 	10.011529755s	44
// 10000 	32.159210139s	44

// 	v1
// 	10 		29.387848324s 	44
//	100		5.753825316s	44
//	250		7.471289272s	44
// 	500		6.836116825s	44
//	1000	8.747139076s	44
