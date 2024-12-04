package global

// redis key
const (
	UniqueKeySeed     string = "Key12368975490255" // UniqueKeySeed 随机数相关的种子
	RedisPrefix       string = "redis_"            // RedisPrefix redis存储随机前缀
	RedisUserKey      string = "redis_user_"       // RedisUserKey 用户相关的前缀
	RedisSessionKey   string = "redis_session_"    // RedisSessionKey 会话相关的前缀
	RedisDataCacheKey string = "redis_data_cache_" // RedisCacheKey 缓存相关的前缀
	RedisConfigKey    string = "redis_config_"     // RedisConfigKey 配置相关的前缀
)
