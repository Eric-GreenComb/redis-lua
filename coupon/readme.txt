高并发秒杀项目总结笔记

https://github.com/zhongzhh8/SecKill-System
https://blog.csdn.net/weixin_41519463/article/details/103892737

hset cid13421234 left 100

redis-cli --eval ~/go/src/dev/redis-lua/coupon/coupon.lua eric-has food cid13421234

redis-cli -x script load < ~/go/src/dev/redis-lua/eval/hello.lua

return c4e251bd1b776832d640ef0a7cba462663a9ef9d

evalsha c4e251bd1b776832d640ef0a7cba462663a9ef9d 2 food meat 21

