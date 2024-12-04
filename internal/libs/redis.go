package libs

import (
	"context"
	"time"
	"yug_server/global"
)

// Set 设置键值对
func RedisSet(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return global.RedisClient.Set(ctx, key, value, expiration).Err()
}

// Get 获取键的值
func RedisGet(ctx context.Context, key string) (string, error) {
	return global.RedisClient.Get(ctx, key).Result()
}

// Delete 删除键
func RedisDelete(ctx context.Context, key string) error {
	return global.RedisClient.Del(ctx, key).Err()
}

// Exists 检查键是否存在
func RedisExists(ctx context.Context, key string) (bool, error) {
	count, err := global.RedisClient.Exists(ctx, key).Result()
	return count > 0, err
}

// Expire 设置键的过期时间
func RedisExpire(ctx context.Context, key string, expiration time.Duration) error {
	return global.RedisClient.Expire(ctx, key, expiration).Err()
}

// TTL 获取键的剩余存活时间
func RedisTTL(ctx context.Context, key string) (time.Duration, error) {
	return global.RedisClient.TTL(ctx, key).Result()
}

// Incr 自增键值
func RedisIncr(ctx context.Context, key string) (int64, error) {
	return global.RedisClient.Incr(ctx, key).Result()
}

// Decr 自减键值
func RedisDecr(ctx context.Context, key string) (int64, error) {
	return global.RedisClient.Decr(ctx, key).Result()
}

// HashSet 设置哈希表的字段
func RedisHashSet(ctx context.Context, key, field string, value interface{}) error {
	return global.RedisClient.HSet(ctx, key, field, value).Err()
}

// HashGet 获取哈希表的字段
func RedisHashGet(ctx context.Context, key, field string) (string, error) {
	return global.RedisClient.HGet(ctx, key, field).Result()
}

// HashDelete 删除哈希表的字段
func RedisHashDelete(ctx context.Context, key, field string) error {
	return global.RedisClient.HDel(ctx, key, field).Err()
}

// HashGetAll 获取哈希表的所有字段和值
func RedisHashGetAll(ctx context.Context, key string) (map[string]string, error) {
	return global.RedisClient.HGetAll(ctx, key).Result()
}

// ListPush 向列表左侧插入元素
func RedisListPush(ctx context.Context, key string, value interface{}) error {
	return global.RedisClient.LPush(ctx, key, value).Err()
}

// ListPop 从列表右侧弹出元素
func RedisListPop(ctx context.Context, key string) (string, error) {
	return global.RedisClient.RPop(ctx, key).Result()
}

// SetAdd 向集合中添加元素
func RedisSetAdd(ctx context.Context, key string, members ...interface{}) error {
	return global.RedisClient.SAdd(ctx, key, members...).Err()
}

// SetMembers 获取集合中的所有成员
func RedisSetMembers(ctx context.Context, key string) ([]string, error) {
	return global.RedisClient.SMembers(ctx, key).Result()
}

// SetRemove 从集合中移除元素
func RedisSetRemove(ctx context.Context, key string, members ...interface{}) error {
	return global.RedisClient.SRem(ctx, key, members...).Err()
}
