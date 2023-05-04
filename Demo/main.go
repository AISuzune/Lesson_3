//package main

//import (
//	"fmt"
//	"runtime"
//	"sync"
//)
//
//var (
//	count int32
//	wg    sync.WaitGroup //这个后面讲锁时讲
//)
//
//func main() {
//	wg.Add(2) //这个后面讲锁时讲
//	go incCount()
//	go incCount()
//	wg.Wait()
//	fmt.Println(count) //这个后面讲锁时讲
//}
//func incCount() {
//	defer wg.Done() //这个后面讲锁时讲
//	for i := 0; i < 2; i++ {
//		value := count
//		// runtime.Gosched() 是让当前 goroutine 暂停的意思，退回执行队列，让其他等待的 goroutine 运行，目的是为了使资源竞争的结果更明显。
//		runtime.Gosched()
//		value++
//		count = value
//	}
//}

//import (
//	"fmt"
//	"sync"
//	"sync/atomic"
//	"time"
//)
//
//var (
//	shutdown int64
//	wg       sync.WaitGroup
//)
//
//// main 函数使用 StoreInt64 函数来安全地修改 shutdown 变量的值。如果哪个 doWork goroutine
//// 试图在 main 函数调用 StoreInt64 的同时调用 LoadInt64 函数，那么原子函数会将这些调用互相同
//// 步，保证这些操作都是安全的，不会进入竞争状态。
//func main() {
//	wg.Add(2)
//	go doWork("A")
//	go doWork("B")
//	time.Sleep(1 * time.Second)
//	fmt.Println("Shutdown Now")
//	atomic.StoreInt64(&shutdown, 1)
//	wg.Wait()
//}
//func doWork(name string) {
//	defer wg.Done()
//	for {
//		fmt.Printf("Doing %s Work\n", name)
//		time.Sleep(250 * time.Millisecond)
//		if atomic.LoadInt64(&shutdown) == 1 {
//			fmt.Printf("Shutting %s Down\n", name)
//			break
//		}
//	}
//}

//

//var (
//	// 逻辑中使用的某个变量
//	count int
//	// 与变量对应的使用互斥锁
//	countGuard sync.Mutex
//)
//
//func GetCount() int {
//	// 锁定
//	countGuard.Lock()
//	// 在函数退出时解除锁定
//	defer countGuard.Unlock()
//	return count
//}

//var (
//	// 逻辑中使用的某个变量
//	count int
//	// 与变量对应的使用互斥锁
//	countGuard sync.RWMutex
//)
//
//func GetCount() int {
//	// 锁定
//	countGuard.RLock()
//	// 在函数退出时解除锁定
//	defer countGuard.RUnlock()
//	return count
//}
//func SetCount(c int) {
//	countGuard.Lock()
//	count = c
//	countGuard.Unlock()
//}
//func main() {
//	// 可以进行并发安全的设置
//	SetCount(1)
//	// 可以进行并发安全的获取
//	fmt.Println(GetCount())
//}

//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	// 构建一个通道
//	ch := make(chan int)
//	// 开启一个并发匿名函数
//	go func() {
//		// 从3循环到0
//		for i := 3; i >= 0; i-- {
//			// 发送3到0之间的数值
//			ch <- i
//			// 每次发送完时等待
//			time.Sleep(time.Second)
//		}
//	}()
//	// 遍历接收通道数据
//	for data := range ch {
//		// 打印通道数据
//		fmt.Println(data)
//		// 当遇到数据0时, 退出接收循环
//		if data == 0 {
//			break
//		}
//	}
//}

// 这个示例程序展示如何用无缓冲的通道来模拟
// 2 个goroutine 间的网球比赛
//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//// wg 用来等待程序结束
//var wg sync.WaitGroup
//
//func init() {
//	rand.Seed(time.Now().UnixNano())
//}
//
//// main 是所有Go 程序的入口
//func main() {
//	// 创建一个无缓冲的通道
//	court := make(chan int)
//	// 计数加 2，表示要等待两个goroutine
//	wg.Add(2)
//	// 启动两个选手
//	go player("Nadal", court)
//	go player("Djokovic", court)
//	// 发球
//	court <- 1
//	// 等待游戏结束
//	wg.Wait()
//}
//
//// player 模拟一个选手在打网球
//func player(name string, court chan int) {
//	// 在函数退出时调用Done 来通知main 函数工作已经完成
//	defer wg.Done()
//	for {
//		// 等待球被击打过来
//		ball, ok := <-court
//		if !ok { // 如果通道被关闭，我们就赢了
//			fmt.Printf("Player %s Won\n", name)
//			return
//		}
//		// 选随机数，然后用这个数来判断我们是否丢球
//		n := rand.Intn(100)
//		if n%13 == 0 {
//			fmt.Printf("Player %s Missed\n", name)
//			// 关闭通道，表示我们输了
//			close(court)
//			return
//		}
//		// 显示击球数，并将击球数加1
//		fmt.Printf("Player %s Hit %d\n", name, ball)
//		ball++
//		// 将球打向对手
//		court <- ball
//	}
//}

// 这个示例程序展示如何用无缓冲的通道来模拟
// 4 个goroutine 间的接力比赛
//

// 这个示例程序展示如何使用
// 有缓存的通道和固定数目的
// goroutine来处理工作
//package main
//
//import "fmt"
//
//func fibonacci(ch, quit chan int) {
//	x, y := 0, 1
//	for {
//		select {
//		case ch <- x:
//			x, y = y, x+y
//		case <-quit:
//			fmt.Println("quit")
//			return
//		}
//	}
//}
//func main() {
//	c := make(chan int)
//	quit := make(chan int)
//	go func() {
//		for i := 0; i < 10; i++ {
//			// 接受通道c传来的值，并打印到控制台
//			fmt.Println(<-c)
//		}
//		// 当协程执行完上述操作后，向quit发送数据
//		quit <- 0
//	}()
//	fibonacci(c, quit)
//}

package main

import "fmt"

func main() {
	for i := range random(100) {
		fmt.Println(i)
	}
}
func random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 0:
			case c <- 1:
			}
		}
	}()
	return c
}
