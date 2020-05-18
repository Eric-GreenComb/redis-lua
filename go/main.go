package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func createScript() string {
	script := `
	-- [[
		--     以时间窗的形式限制IP访问
		--     传入参数:
		--     业务标识
		--     IP
		--     限制时间
		--     限制时间内的访问次数
		--     请求时间戳
		-- ]]
		
		local bizID = tostring(KEYS[1])
		local ip = tostring(KEYS[2])
		
		local expireSeconds = tonumber(ARGV[1])
		local limitTimes = tonumber(ARGV[2])
		
		-- 传入额外参数,请求时间戳
		local timestamp = tonumber(ARGV[3])
		local lastTimestamp
		
		local identify = "limit" .. "_" .. bizID .. "_" .. ip
		local times = redis.call("LLEN",identify)
		
		if times < limitTimes then
			redis.call("RPUSH", identify, timestamp)
			return 1
		end
		
		lastTimestamp = redis.call("LRANGE", identify, 0, 0)
		lastTimestamp = tonumber(lastTimestamp[1])
		
		if lastTimestamp + expireSeconds >= timestamp then
			return 0
		end
		
		redis.call("LPOP", identify)
		redis.call("RPUSH", identify, timestamp)
		
		return 1
	`

	return script
}

func scriptCache2Redis(c *redis.Client) string {
	script := createScript()
	var ret string

	if result, err := c.(script).Result(); err != nil {
		panic("缓存脚本到节点失败")
	} else {
		ret = result
	}
	fmt.Println("sha", ret)
	return ret
}

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

	// 将脚本缓存到节点,执行一次拿到结果即可
	sha := scriptCache2Redis(client)

	// 执行缓存脚本
	// sha := "145ecc4c04230ee2479ebff7927dc761ff134201"
	evalSha(sha)
}

func evalSha(sha string) {
	// 执行缓存脚本
	_timestamp := time.Now().Unix()
	ret := client.EvalSha(sha, []string{
		"gopperin",
		"127.0.0.1",
	}, 60, 3, _timestamp)

	if result, err := ret.Result(); err != nil {
		fmt.Println("发生异常,返回值", err.Error())
	} else {
		fmt.Println("返回值", result)
	}
}
