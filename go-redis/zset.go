package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// zsetDemon
//
//	@Description: go-redis操作有序集合zset
func zsetDemon() {
	//	redis	key
	zsetKey := "language_rank"
	//	redis value
	// redis.Z{Score:90.0,Member:"Golang"},zadd 想redis中添加有序集合，redis.Z 结构体包含有序集合的成员和分数
	languages := []redis.Z{
		{
			Score:  90.0,
			Member: "Golang",
		},
		{
			Score:  98.0,
			Member: "JAVA",
		},
		{
			Score:  95.0,
			Member: "Python",
		},
		{
			Score:  92.0,
			Member: "C++",
		},
		{
			Score:  99.0,
			Member: "JavaScript",
		},
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	//	ZADD,redis命令中的zadd，添加一个或多个成员
	//	ZADD key score member [score member...]
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("ZADD error: %v\n", err)
		return
	}
	fmt.Println("ZADD success")

	//	把Golang的分数加10，redis命令中的zincrby，增加一个成员的分数，修改操作
	//	ZINCRBY key increment member
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("ZINCRBY error: %v\n", err)
		return
	}
	fmt.Printf("ZINCRBY success, Golang new score: %v\n", newScore)

	//	取分数最高的3个
	//ZRANGE language_rank 0 2 WITHSCORES  //  redis命令中的zrange，返回有序集中指定区间内的成员，其中成员的位置按分数值递增(从小到大)来排列
	//ZREVRANGE language_rank 0 2 WITHSCORES //redis命令中的zrevrange，返回有序集中指定区间内的成员，其中成员的位置按分数值递减(从大到小)来排列
	//ZREVRANGE key start stop [WITHSCORES]
	val := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range val {
		fmt.Println(z.Member, z.Score)
	}

	//	取95-100分之间的
	rangeBy := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	//ZRANGEBYSCORE key min max [WITHSCORES] //  redis命令中的zrangebyscore，返回有序集中指定分数区间内的成员，其中成员的位置按分数值递增(从小到大)来排列
	ret, err := rdb.ZRangeByScoreWithScores(ctx, zsetKey, rangeBy).Result()
	if err != nil {
		fmt.Printf("ZRANGEBYSCOREWITHSCORES error: %v\n", err)
		return
	}

	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

}

func scanKeyDemo1() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	vals, err := rdb.Keys(ctx, "user:*").Result()
	if err != nil {
		fmt.Printf("get Keys error: %v\n", err)
		return
	}
	for _, v := range vals {
		fmt.Printf("key: %v\n", v)
	}

	//ret, err := rdb.Keys(ctx, "user:*").Result()
	//if err != nil {
	//	fmt.Printf("get Keys error: %v\n", err)
	//	return
	//}
	//for _, s := range ret {
	//	fmt.Printf("key: %v\n", s)
	//}
}

// scanKeysDemo1 按前缀查找所有key示例
func scanKeysDemo2() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var cursor uint64

	for {
		var keys []string
		var err error
		// 将redis中所有以prefix:为前缀的key都扫描出来
		keys, cursor, err = rdb.Scan(ctx, cursor, "user:*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
}

func scanKeysDemo3() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	// 按前缀扫描key
	iterator := rdb.Scan(ctx, 0, "user*", 0).Iterator()
	for iterator.Next(ctx) {
		fmt.Printf("key: %s\n", iterator.Val())
	}

	err := iterator.Err()
	if err != nil {
		panic(err)
	}
}

// 批量删除指定前缀的key
func delKeysByMatch(match string, timout time.Duration) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timout)
	defer cancelFunc()

	iterator := rdb.Scan(ctx, 0, match, 0).Iterator()
	for iterator.Next(ctx) {
		err := rdb.Del(ctx, iterator.Val()).Err()
		if err != nil {
			panic(err)
		}
	}

	err := iterator.Err()
	if err != nil {
		panic(err)
	}
}
