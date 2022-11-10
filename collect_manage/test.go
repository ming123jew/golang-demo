// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var (
// 	num    int64
// 	wg     sync.WaitGroup
// 	rwlock sync.RWMutex
// )

// func write() {
// 	rwlock.Lock()

// 	num = num + 1
// 	// 模拟真实写数据消耗的时间
// 	time.Sleep(10 * time.Millisecond)

// 	rwlock.Unlock()
// 	wg.Done()
// }

// func read() {
// 	rwlock.RLock()

// 	// 模拟真实读取数据消耗的时间
// 	time.Sleep(time.Millisecond)

// 	rwlock.RUnlock()
// 	wg.Done()
// }

// func main() {
// 	// 用于计算时间 消耗
// 	start := time.Now()

// 	// 开5个协程用作 写
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go write()
// 	}

// 	// 开500 个协程，用作读
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go read()
// 	}

// 	// 等待子协程退出
// 	wg.Wait()
// 	end := time.Now()

// 	// 打印程序消耗的时间
// 	fmt.Println(end.Sub(start))
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// 	"time"
// )

// var num int64
// var l sync.Mutex
// var wg sync.WaitGroup

// // 普通版加函数
// func add() {
// 	num = num + 1
// 	wg.Done()
// }

// // 互斥锁版加函数
// func mutexAdd() {
// 	l.Lock()
// 	num = num + 1
// 	l.Unlock()
// 	wg.Done()
// }

// // 原子操作版加函数
// func atomicAdd() {
// 	atomic.AddInt64(&num, 1)
// 	wg.Done()
// }

// func main() {
// 	// 目的是 记录程序消耗时间
// 	start := time.Now()
// 	for i := 0; i < 20000000; i++ {

// 		wg.Add(1)

// 		// go add() // 无锁的  add函数 不是并发安全的
// 		go mutexAdd() // 互斥锁的 add函数 是并发安全的，因为拿不到互斥锁会阻塞，所以加锁性能开销大

// 		// go atomicAdd() // 原子操作的 add函数 是并发安全，性能优于加锁的
// 	}

// 	// 等待子协程 退出
// 	wg.Wait()

// 	end := time.Now()
// 	fmt.Println(num)
// 	// 打印程序消耗时间
// 	fmt.Println(end)
// 	fmt.Println(end.Sub(start))
// }

// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// )

// func main() {

// 	var Atomicvalue atomic.Value
// 	Atomicvalue.Store([]int{1, 2, 3, 4, 5})
// 	fmt.Println("main before testA: ", Atomicvalue)
// 	testA(Atomicvalue)
// 	fmt.Println("main after testA: ", Atomicvalue)

// 	// 复位
// 	Atomicvalue.Store([]int{1, 2, 3, 4, 5})
// 	fmt.Println("\n")

// 	fmt.Println("main before testB: ", Atomicvalue)
// 	testB(&Atomicvalue)
// 	fmt.Println("main after testB: ", Atomicvalue)
// }

// func testA(Atomicvalue atomic.Value) {
// 	Atomicvalue.Store([]int{6, 7, 8, 9, 10})
// 	fmt.Println("testA: ", Atomicvalue)
// }

// func testB(Atomicvalue *atomic.Value) {
// 	Atomicvalue.Store([]int{6, 7, 8, 9, 10})
// 	fmt.Println("testB: ", Atomicvalue)
// }

package test
