-------简单数据-------
local tab ={}
tab["Himi"] = "himigame.com"
--数据转json
local cjson = require "cjson"
local jsonData = cjson.encode(tab)
 
print(jsonData)
-- 打印结果:  {"Himi":"himigame.com"}
 
--json转数据
local data = cjson.decode(jsonData)
 
print(data.Himi)
-- 打印结果:  himigame.com