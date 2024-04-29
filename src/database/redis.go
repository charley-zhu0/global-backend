/*
 * @Author: charley zhu
 * @Date: 2023-10-15 08:25:22
 * @LastEditTime: 2023-10-15 09:25:19
 * @LastEditors: charley zhu
 * @Description:
 */
package database

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func init() {
	if redisClient != nil {
		return
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		PoolSize: 20,
	})
}

func RdbDo(cmd string, args ...interface{}) (interface{}, error) {
	ctc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := redisClient.Do(ctc, args...).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return val, nil
}

func GetRdb() *redis.Client {
	return redisClient
}
