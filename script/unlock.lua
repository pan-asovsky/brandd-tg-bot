-- KEYS[1] = ключ слота
-- ARGV[1] = UUID
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end