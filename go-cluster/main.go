package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func createScript() *redis.Script {
	script := redis.NewScript(`
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
	`)

	return script
}

func scriptCache2Cluster(c *redis.ClusterClient) string {
	script := createScript()
	var ret string

	c.ForEachMaster(func(m *redis.Client) error {
		if result, err := script.Load(m).Result(); err != nil {
			panic("缓存脚本到节点失败")
		} else {
			ret = result
		}
		return nil
	})
	return ret
}

func main() {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			":6379",
		},
	})

	// 将脚本缓存到所有节点,执行一次拿到结果即可
	sha := scriptCache2Cluster(redisdb)

	// 执行缓存脚本
	ret := redisdb.EvalSha(sha, []string{
		"pro{yes}",
		"127.0.0.1{yes}",
	}, 10, 3, 1580479510)

	if result, err := ret.Result(); err != nil {
		fmt.Println("发生异常,返回值", err.Error())
	} else {
		fmt.Println("返回值", result)
	}

}
