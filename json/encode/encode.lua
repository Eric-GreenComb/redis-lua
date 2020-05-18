-- RedisLuaCjsonEncode.lua文件
local userName = ARGV[1];

local userObject = {
            name = userName,
            age = 14,
            address = 'China'
        }

local userJson = cjson.encode(userObject);

if redis.call('set', KEYS[1], userJson) == 0
then
        return -1
else
        return userJson
end
