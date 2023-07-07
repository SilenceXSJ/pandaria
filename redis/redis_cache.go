package internal

import (
	"context"
	"time"
)

// LoadDataWithCache 加载数据 如果缓存存在就返回 , 没有就获取 再写入缓存
func LoadDataWithCache(ctx context.Context, client *redis.Client, key string, expiration time.Duration, supplier func() (string, error)) (string, error) {
	res, err := client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			s, err := supplier()
			if err != nil {
				return "", err
			}
			client.SetEX(ctx, key, s, expiration)
			return s, nil
		} else {
			return "", err
		}
	} else {
		return res, err
	}
}
