local good = tostring(KEYS[1])
local detail = tostring(KEYS[2])

local num = tonumber(ARGV[1])

if num == 1 then
    return "num:1"
end

return good .. ":" .. detail