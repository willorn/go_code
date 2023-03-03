package main

import (
	"fmt"
	"time"
)

// timeDemo 时间对象的年月日时分秒
func main() {
	//now := time.Now() // 获取当前时间
	//fmt.Printf("current time:%v\n", now)

	now := time.Now()

	// Go 语言中时间格式化的布局不是常见的Y-m-d H:M:S，而是使用 2006-01-02 15:04:05.000（记忆口诀为2006 1 2 3 4 5）
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	fmt.Printf("当前时间：%v\n", now.Format("2006/01/02 15:04:05.000"))

	// 编写程序统计一段代码的执行耗时时间，单位精确到微秒。
	for i := 0; i < 10; i++ {
		startTime := time.Now()
		time.Sleep(300 * time.Millisecond)
		//println(time.Now().Sub(startTime).Microseconds())
		fmt.Printf("运行时间：%v 微秒\n", time.Now().Sub(startTime).Microseconds())
	}
}

// 定时器
func tickDemo() {
	gap := time.Second
	actTime := time.Tick(gap) //定义一个1秒间隔的定时器
	for i := range actTime {
		fmt.Println(i) //每秒都会执行的任务
	}
}
