package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("init redis")
}

func main() {

	// 执行缓存脚本
	sha := "c4e251bd1b776832d640ef0a7cba462663a9ef9d"
	evalSha(sha)
}

func evalSha(sha string) {
	// 执行缓存脚本
	ret := client.EvalSha(sha, []string{
		"gopperin",
		"meat",
	}, 21)

	fmt.Println(ret.Val())

	// if result, err := ret.Result(); err != nil {
	// 	fmt.Println("发生异常,返回值", err.Error())
	// } else {
	// 	fmt.Println("返回值", result)
	// }
}
