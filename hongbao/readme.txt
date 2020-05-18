详解利用redis + lua解决抢红包高并发的问题
https://www.jb51.net/article/98624.htm
4.抢红包时，先判断用户是否抢过红包，如果没有，则从未消费红包队列中取出一个小红包，再push到另一个已消费队列中，最后把用户ID放入去重的map中。

LPUSH hongBaoList "{\"id\":1234,\"money\":200000}"
LPUSH hongBaoList "{\"id\":1233,\"money\":100000}"
LPUSH hongBaoList "{\"id\":1232,\"money\":300000}"

redis-cli --eval ~/go/src/dev/redis-lua/hongbao/hongbao.lua hongBaoList hongBaoConsumedList hongBaoConsumedMap 234321
redis-cli --eval ~/go/src/dev/redis-lua/hongbao/hongbao.lua hongBaoList hongBaoConsumedList hongBaoConsumedMap 234322

hget hongBaoConsumedMap 234321
lrange hongBaoList 0 -1
lrange hongBaoConsumedList 0 -1
