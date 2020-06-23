package chan_r_w

import (
	"fmt"
	"time"
)

func UseChan() {
	t1 := time.Now()
	a := make(chan bool)
	close(a)
	aa := make(chan bool, 20)
	go func() {
		for i := 0; i < 100; i++ {
			aa <- true
		}
		close(aa)
	}()
	b := 0
	for i := 0; i < 1; i++ {
		go func() {
			for <-aa {
				b++
				fmt.Println("ok", "====", b)
				//time.Sleep(time.Second)
			}
		}()
	}

	//for <-ch {
	//	go chanRead1(ch)
	//	go chanRead2(ch)
	//}
	fmt.Println(time.Now().Sub(t1))
	<- a
}

func chanWrite(ch chan<- int, n int) {
	ch <- n
	fmt.Println("chanWrite", n)
}

func chanRead1(ch <-chan int) {
	fmt.Println("chanRead1", <-ch)
}

func chanRead2(ch <-chan int) {
	fmt.Println("chanRead2", <-ch)
}
