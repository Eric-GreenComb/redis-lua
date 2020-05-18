redis-cli --eval ~/go/src/dev/redis-lua/eval/hello.lua food meat , 1
redis-cli --eval ~/go/src/dev/redis-lua/eval/hello.lua food meat , 21

redis-cli -x script load < ~/go/src/dev/redis-lua/eval/hello.lua

return c4e251bd1b776832d640ef0a7cba462663a9ef9d

evalsha c4e251bd1b776832d640ef0a7cba462663a9ef9d 2 food meat 21