-- RedisLuaCjsonDecode.lua文件
local userInfo = redis.call('get', KEYS[1])

local userJson = cjson.decode(userInfo)

return userJson.name;