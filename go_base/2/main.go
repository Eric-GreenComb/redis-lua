package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

//noinspection GoInvalidCompositeLiteral
func main() {

	// Client.FlushAll()

	Client.Set("foo", "bar", -1)

	var luaScript = `return redis.call("INFO")`
	result, err := Client.ScriptLoad(luaScript).Result() //返回的脚本会产生一个sha1哈希值,下次用的时候可以直接使用这个值，类似于
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	foo := Client.EvalSha(result, []string{})
	fmt.Println(foo.Val())
}
