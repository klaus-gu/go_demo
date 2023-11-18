package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"time"
)

/**
导入：go get github.com/go-redis/redis/v8
*/

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "root",
		DB:       0,
	})
	// 字符串操作
	rdb.Set(ctx, "key", "value", 10*time.Second)
	result, _ := rdb.Get(ctx, "key").Result()
	fmt.Println(result)

	// hash 操作 values 需要以 k-v 的形式加入进去
	rdb.HSet(ctx, "hkey", "name", "zhangsan", "age", 11)
	name, _ := rdb.HGet(ctx, "hkey", "name").Result()
	fmt.Println(name)

	// list 操作
	rdb.LPush(ctx, "Lkey", 1, 2, 3, 4, 5)
	// 操作列表的LSet的时候必须是操作已经存在的列表，不能操作不存在的
	s, err := rdb.LSet(ctx, "Lkey", 1, "1").Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
	// LPushX ：当对应的 key 存在的时候才会进行插入操作
	lpushxr, err := rdb.LPushX(ctx, "LkeyX", 1, 2).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(lpushxr)
	lrangeR, err := rdb.LRange(ctx, "Lkey", 0, 4).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	str, err := jsoniter.MarshalToString(lrangeR)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(str)

	// set 操作
	rdb.SAdd(ctx, "Skey", 1, 2, 3, 3, 4)
	// 获取元素个数
	i, err := rdb.SCard(ctx, "Skey").Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(jsoniter.MarshalToString(i))
	sval, err := rdb.SMembers(ctx, "Skey").Result()
	fmt.Println(jsoniter.MarshalToString(sval))

	// ZSet 操作
	rdb.ZAdd(ctx, "Zkey", &redis.Z{Score: 2.5, Member: "zhangsan"})
	rdb.ZAdd(ctx, "Zkey", &redis.Z{Score: 3.51, Member: "lisi"})
	score, err := rdb.ZScore(ctx, "Zkey", "lisi").Result()
	fmt.Println(score)
}
