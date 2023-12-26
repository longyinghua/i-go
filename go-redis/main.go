package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// initRedis1
//
//	@Description: 初始化redis连接
//	@return client *redis.Client
func initRedis1() (client *redis.Client) {
	opt, err := redis.ParseURL("redis://:123456@10.40.3.29:6379/1?dial_timeout=5s")
	if err != nil {
		panic(err)
	}
	//fmt.Println("addr is", opt.Addr)
	//fmt.Println("db is", opt.DB)
	//fmt.Println("password is", opt.Password)
	//fmt.Println("dial timeout is", opt.DialTimeout)

	newClient := redis.NewClient(opt)

	return newClient
}

// initRedis2
//
//	@Description: 初始化redis连接
//	@return client *redis.Client
func initRedis2() (client *redis.Client) {
	connetionOpt := &redis.Options{
		Addr:     "10.40.3.29:6379",
		Password: "123456",
		DB:       0,
	}
	newClient := redis.NewClient(connetionOpt)

	return newClient
}

var rdb *redis.Client

// getValueFromRedis redis.Nil判断
func getValueFromRedis1(client *redis.Client, key, defaultValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			return defaultValue, nil //  如果是key不存在，返回默认值
		}

		// 出其他错了
		return "", err
	}
	return val, nil
}

func getValueFromRedis2(ctx context.Context, key, defaultValue string) (string, error) {
	cmd := rdb.Do(ctx, "get", "test")

	result, err := cmd.Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("key is not exist\n")
			return defaultValue, err
		}
		return "", err
	}

	fmt.Printf("result is %s\n", result)
	return result.(string), nil

}

func main() {
	//初始化redis连接
	//client := initRedis2()

	rdb = initRedis1() //  初始化redis连接，全局变量redis连接对象

	//zsetDemon()

	//scanKeyDemo1()
	//scanKeysDemo2()
	//scanKeysDemo3()
	delKeysByMatch("user*", 10*time.Second)

	//ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	//defer cancelFunc()

	//result, err := rdb.Get(ctx, "test").Result()
	//if err != nil {
	//	if err == redis.Nil {
	//		fmt.Printf("key is not exist\n")
	//		return
	//	}
	//	fmt.Printf("err is %s\n", err)
	//	return
	//}
	//fmt.Printf("result is %s\n", result)

	//valueFromRedis2, err := getValueFromRedis2(ctx, "test", "defaultValues")
	//if err != nil {
	//	fmt.Printf("err is %s\n", err)
	//	fmt.Printf("value is %s\n", valueFromRedis2)
	//	return
	//}
	//fmt.Printf("value is %s\n", valueFromRedis2)

	//ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	//defer cancelFunc()

	//valueFromRedis, err := getValueFromRedis(client, "test", "defaultValues")
	//if err != nil {
	//	fmt.Printf("err is %s\n", err)
	//	return
	//}
	//fmt.Printf("value is %s\n", valueFromRedis)

	////执行命令获取结果
	//val, _ := client.Get(ctx, "abc").Result()
	//fmt.Printf("result is %s\n", val)
	//
	////  直接执行命令获取到命令对象
	//stringCmd := client.Get(ctx, "abc")
	//fmt.Println(stringCmd.Val()) //  获取值
	//fmt.Println(stringCmd.Err()) //  获取错误
	//
	////直接执行命令获取到错误
	//err := client.Get(ctx, "abc").Err()
	//if err != nil {
	//	fmt.Printf("err is %s\n", err)
	//	return
	//}
	//
	////直接执行命令获取值
	//value := client.Get(ctx, "abc").Val()
	//fmt.Printf("result is %s\n", value)

	////直接执行命令获取错误
	//err := client.Get(ctx, "abc").Err()
	//if err != nil {
	//	fmt.Printf("err is %s\n", err)
	//	return
	//}

	//get, err := client.Get(ctx, "abc").Result()
	//fmt.Printf("result is %s\n", get)
	//if err != nil {
	//	if err == redis.Nil {
	//		fmt.Printf("key does not exists------")
	//		return
	//	}
	//	//fmt.Printf("err is %s\n", err)
	//	panic(err)
	//}

	//result, errs := client.Do(ctx, "get", "key").Result()

	//result, errs := client.Do(ctx, "get", "abc").Result() //  get为redis获取指定健的命令，“abc”为键名，即通过get命令获取键名为abc的值
	//if errs != nil {
	//	if errs == redis.Nil {
	//		fmt.Printf("key does not exists------")
	//		return
	//	}
	//	//fmt.Printf("err is %s\n", err)
	//	panic(errs)
	//}
	//fmt.Printf("the result interface is %s\n", result.(string)) //  result为interface类型，需要强转为string类型进行结果输出，

	////Do方法返回一个Cmd类型，您可以使用它来获取您想要的类型。在这里，我们使用.Text()来获取字符串类型的结果。如果操作失败，会返回一个错误。
	//text, err := client.Do(ctx, "get", "abc").Text()
	//if err != nil {
	//	fmt.Printf("err is %s\n", err)
	//	return
	//}
	//fmt.Printf("the text result is %s\n", text)

	////使用.Bool()来获取布尔类型的结果
	//flag, errss := client.Do(ctx, "get", "cv").Bool()
	//if errss != nil {
	//	return
	//}
	//fmt.Printf("the flag result is %t\n", flag)

	////	执行任意命令获取错误
	//err := client.Do(ctx, "set", "key", 10, "EX", 3600).Err() //  返回执行命令的cmd对象，再调用err方法获取错误
	//if err != nil {
	//	fmt.Printf("err is %s\n", err)
	//	return
	//}
	//
	////	执行任意命令获取结果
	//result, err := client.Do(ctx, "set", "test", "true", "EX", 3600).Result()
	//if err != nil {
	//	fmt.Printf("err is %s\n", err)
	//	return
	//}
	//fmt.Printf("result is %v\n", result.(string))
	//
	//value, err := client.Do(ctx, "get", "test").Result()
	//if err != nil {
	//	fmt.Printf("err is %s\n", err)
	//	return
	//}
	//fmt.Printf("value is %v\n", value.(string))
	//
	//errs := client.Do(ctx, "del", "test").Err()
	//if errs != nil {
	//	fmt.Printf("err is %s\n", errs)
	//	return
	//}
	//fmt.Println("del success")
}
