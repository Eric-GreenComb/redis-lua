原理
计数器算法是指在一段窗口时间内允许通过的固定数量的请求, 比如10次/秒, 500次/30秒.

如果设置的时间粒度越细, 那么限流会更平滑.

key 可以但不限于以下的情况

ip + 接口
user_id + 接口
优点

实现简单
缺点

粒度不够细的情况下, 会出现在同一个窗口时间内出现双倍请求数
注意

尽量保持时间粒度精细
场景分析
eg. 1000/3s 的限流

极端情况1:

第1秒请求数 10

第2秒请求数 10

第3秒请求数 980

第4秒请求数 900

第5秒请求数 100

第6秒请求数 0

此时注意第3~5秒内的总请求数高达 1980

极端情况2:

第1秒请求数 1000

第2秒请求数 0

第3秒请求数 0

此时后续的第2~3秒会出现大量拒绝请求

redis-cli --eval ~/go/src/dev/redis-lua/limit/limit2/limit.lua limit_127.0.0.1 , 10 10