package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"time"
)

//
type dataChan struct {
	Data interface{}
	Err  error
}

type DataFn func() (interface{}, error)

func GetDataChansWithCtx(ctx context.Context, fns []DataFn) chan *dataChan {
	result := make(chan *dataChan, len(fns))
	for _, fn := range fns {
		go func(fn DataFn) {
			var resp interface{}
			var err error
			done := make(chan struct{})
			// create channel with buffer size 1 to avoid goroutine leak
			panicChan := make(chan interface{}, 1)

			go func() {
				defer func() {
					if p := recover(); p != nil {
						panicChan <- p
					}
				}()

				resp, err = fn()
				close(done)
			}()

			select {
			case p := <-panicChan:
				panic(p)
			case <-ctx.Done():
				result <- &dataChan{
					Data: nil,
					Err:  ctx.Err(),
				}
			case <-done:
				result <- &dataChan{
					Data: resp,
					Err:  err,
				}
			}
		}(fn)
	}
	return result
}

func GetDataChans(fns []DataFn) chan *dataChan {
	result := make(chan *dataChan, len(fns))
	for _, fn := range fns {
		go func(fn DataFn) {
			data, err := fn()
			result <- &dataChan{
				Data: data,
				Err:  err,
			}
		}(fn)
	}
	return result
}

type A string

func getDataA(n int) (A, error) {
	if n%30 == 0 {
		return "", errors.New("get A error")
	}
	time.Sleep(time.Duration(n) * time.Millisecond)
	return A("-A" + strconv.Itoa(n)), nil
}

type B string

func getDataB(n int) (B, error) {
	if n%30 == 0 {
		return "", errors.New("get B error")
	}
	time.Sleep(time.Duration(n) * time.Millisecond)
	return B("-B" + strconv.Itoa(n)), nil
}

type C string

func getDataC(n int) (C, error) {
	if n%30 == 0 {
		return "", errors.New("get C error")
	}
	time.Sleep(time.Duration(n) * time.Millisecond)
	return C("-C" + strconv.Itoa(n)), nil
}

func run(ans string, datas chan *dataChan, n int) (string, error) {
	for i := 0; i < n; i++ {
		data, ok := <-datas
		if !ok {
			return "", errors.New("chan close")
		}

		if err := data.Err; err != nil {
			return "", err
		}

		switch data.Data.(type) {
		case A:
			ans += string(data.Data.(A))
		case B:
			ans += string(data.Data.(B))
		case C:
			ans += string(data.Data.(C))
		}
	}

	return ans, nil
}

func main() {
	rand.Seed(time.Now().Unix())
	ctxTimeout := rand.Intn(700)
	ATimeout := rand.Intn(500)
	BTimeout := rand.Intn(500)
	CTimeout := rand.Intn(500)
	fmt.Println("ctx timeout := ", ctxTimeout)
	fmt.Println("A timeout := ", ATimeout)
	fmt.Println("B timeout := ", BTimeout)
	fmt.Println("C timeout := ", CTimeout)

	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ctxTimeout)*time.Millisecond)
	defer cancel()

	fns := []DataFn{
		func() (interface{}, error) {
			return getDataA(ATimeout)
		},
		func() (interface{}, error) {
			return getDataB(BTimeout)
		},
		func() (interface{}, error) {
			return getDataC(CTimeout)
		},
	}

	datas := GetDataChansWithCtx(ctx, fns)

	ans := ""
	ans, err := run(ans, datas, len(fns))
	latency := time.Since(start)
	fmt.Println("latency := ", latency)
	if err != nil {
		fmt.Println("err := ", err)
		return
	}

	fmt.Println("ans := ", ans)
}
