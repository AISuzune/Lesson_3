package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup //这个后面讲锁时讲
)

func main() {
	wg.Add(2) //这个后面讲锁时讲
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count) //这个后面讲锁时讲
}
func incCount() {
	defer wg.Done() //这个后面讲锁时讲
	for i := 0; i < 2; i++ {
		value := count
		// runtime.Gosched() 是让当前 goroutine 暂停的意思，退回执行队列，让其他等待的 goroutine 运行，目的是为了使资源竞争的结果更明显。
		runtime.Gosched()
		value++
		count = value
	}
}
