-- KEYS[1] = ключ слота
-- ARGV[1] = UUID
-- ARGV[2] = TTL в миллисекундах
return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2]) and 1 or 0