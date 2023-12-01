package main

import (
	"context"
	"fmt"
	"time"
)

func getcontext() {
	//创建一个上下文context
	ctx := context.Background()

	//在这个上下文上附加值
	ctx = context.WithValue(ctx, "userId", "123")

	ctx = context.WithValue(ctx, "userName", "Tom")

	//在这个context上附加取消信号
	ctx, cancel := context.WithCancel(ctx)

	//模拟一个长任务
	go func() {
		for {
			select {
			case <-ctx.Done():
				//接受到取消context的取消信号，就停止任务，执行下面的代码
				//fmt.Println("Task cancelled")
				//获取上下文结束的原因
				fmt.Println(ctx.Err())
				return
			default:
				fmt.Println("Task running")
				//睡一秒
				time.Sleep(time.Second)
			}
		}
	}()

	//程序可以决定何时出发取消信号
	time.Sleep(time.Second * 5)
	cancel()

	//等待任务退出
	fmt.Println("Main function finished")
	time.Sleep(time.Second * 5)
}

func main() {
	//getcontext()
	ctx := context.WithValue(context.Background(), "userId", "123")
	value := ctx.Value("userId")
	fmt.Println(value)
	fmt.Printf("type: %T", value)

}
