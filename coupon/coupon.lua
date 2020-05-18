-- Check if User has coupon
-- KEYS[1]: hasCouponKey "{username}-has"
-- KEYS[2]: couponName   "{couponName}"
-- KEYS[3]: couponKey    "{couponName}-info"
-- 返回值有-1, -2, -3, 都代表抢购失败
-- 返回值为1代表抢购成功

-- Check if coupon exists and is cached
local couponLeft = redis.call("hget", KEYS[3], "left");
if (couponLeft == false) then
	return -2;  -- No such coupon
end

--- couponLeft是字符串类型
if (tonumber(couponLeft) == 0) then
	return -3;  --  No Coupon Left.
end

-- Check if the user has got the coupon --
local userHasCoupon = redis.call("sismember", KEYS[1], KEYS[2]);
if (userHasCoupon == 1)	then
	return -1;
end

-- User gets the coupon --
redis.call("hset", KEYS[3], "left", couponLeft - 1);
redis.call("sadd", KEYS[1], KEYS[2]);
return 1;