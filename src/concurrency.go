package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("-------------------------- goroutines --------------------------")
	/**
	1、goroutine是Go runtime管理的轻量级线程

	go f(x,y,z)
	开启一个新的goroutine
	f,x,y,z的求值发生在当前的goroutine中，而f的执行发生在新的goroutine中
	goroutine运行在相同的地址空间，因此对共享内存的访问需要同步。sync包中提供了有用同步工具。
	 */
	go say("world")
	say("hello")

	fmt.Println("-------------------------- channels --------------------------")
	/**
	2、channel是典型的可收发值的导管
	ch <- v 向ch发送值
	v := <-ch 从ch接收值
	默认时，收发都会阻塞，直到另一个准备就绪。这使得goroutines之间不需要使用lock或condition变量来进行同步
	 */
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	/**
	可以为channel设置缓冲区
	 */
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	/**
	发送者可以关闭channel
	接收者可以接受第二个返回值来判断channel是否关闭，若第二个返回值返回false则表示所有值已经接收完毕，且channel已经关闭
	for i := range channel 会一直循环下去，直到channel被关闭。
	 */
	cc := make(chan int, 10)
	go fibonacci(cap(cc), cc)
	for i := range cc {
		fmt.Println(i)
	}

	fmt.Println("-------------------------- select --------------------------")
	/**
	3、select语句让一个goroutine等待，直到其中某个case可以执行。
	当多个case可执行时，会随机选择其中之一
	 */
	ccc := make(chan int)
	quit := make(chan int)
	go func() {
		for i:=0;i<10;i++ {
			fmt.Println(<-ccc)
		}
		quit <- 0
	}()
	fibonacci2(ccc,quit)

	/**
	当select的所有case都未准备就绪时，会运行default
	 */
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
Loop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break Loop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	fmt.Println("-------------------------- sync.Mutex --------------------------")
	/**
	4、sync.Mutex相当于锁，它有Lock和Unlock方法
	 */
	sc := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go sc.Inc("somekey")
		if i%100 == 0 {
			go sc.Value("somekey")
		}
	}

	time.Sleep(time.Second)
	fmt.Println("********")
	fmt.Println(sc.Value("somekey"))

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}