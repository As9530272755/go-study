package main

import (
	"fmt"
	// 引入 time 包
	"time"
)

func main()  {
	// 看看日期和时间相关的函数和使用方法

	// 1.获取当前时间
	now := time.Now()
	fmt.Printf("now = %v now = %T\n",now,now)

	// 2.通过这个 now 能够获取到年月日时分秒
	fmt.Printf("年 = %v\n",now.Year())
	fmt.Printf("月 = %v\n",int(now.Month()))
	fmt.Printf("日 = %v\n",now.Day())
	fmt.Printf("时 = %v\n",now.Hour())
	fmt.Printf("分 = %v\n",now.Minute())
	fmt.Printf("秒 = %v\n",now.Second())

	// 格式化日期时间
	fmt.Printf("当前时间 %d-%d-%d %d:%d:%d\n",
	now.Year(),int(now.Month()),now.Day(),now.Hour(),now.Minute(),now.Second())

	dateStr := fmt.Sprintf("当前时间 %d-%d-%d %d:%d:%d\n",now.Year(),
	now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())

	fmt.Printf("dateStr = %v \n",dateStr)

	// 格式化日期的第二种方式
	fmt.Printf(now.Format("当前时间：2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("当前年份：2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("当前时分秒：15:04:05"))
	fmt.Println()

	// 假如就像输出当前月份
	fmt.Printf(now.Format("当前月份：01\n"))

	// 需求，每隔 1 秒钟打印一个数字，打印到 100 时就退出。
	i := 0
	for {
		i++
		fmt.Printf("当前毫秒数 = %v\n",i)
		// 打印完了一次之后就休眠 0.1 秒
		time.Sleep(time.Microsecond * 100)
		// 加个判断如果 i 等于 100 我们就 break 退出当前循环
		if i == 100 {
			break
		}
	}

	// unix 和 unixnano 的使用
	fmt.Printf("unix 时间戳 = %v unixnano 时间戳 = %v\n",now.Unix(),now.UnixNano())
}
