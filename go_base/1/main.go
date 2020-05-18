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
}

//noinspection GoInvalidCompositeLiteral
func main() {

	// client.FlushAll()

	// client.Set("foo", "bar", 0)

	// var luaScript = redis.NewScript(`return redis.call("GET" , KEYS[1])`)

	// n, err := luaScript.Run(client, []string{"foo"}).Result()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(n, err)

	// var luaScript = `return redis.call("GET" , KEYS[1])`
	// sha, err := client.ScriptLoad(luaScript).Result() //返回的脚本会产生一个sha1哈希值,下次用的时候可以直接使用这个值，类似于
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(sha)

	sha := "007a9c9da3393b9d540afc7c524ca1a0a4df3bc7"
	foo := client.EvalSha(sha, []string{"foo"})
	fmt.Println(foo.Val())
}
