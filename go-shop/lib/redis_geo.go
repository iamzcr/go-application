package lib

import (
	"github.com/go-redis/redis"
	"go-shop/db"
)

type RedisLocation struct {
	Key       string
	Name      string
	Longitude float64
	Latitude  float64
}

// 添加地理空间信息
func (rl *RedisLocation) GeoAdd() (err error) {
	err = db.RedisClient.GeoAdd("locations", &redis.GeoLocation{
		Name:      rl.Name,
		Longitude: rl.Longitude,
		Latitude:  rl.Latitude,
	}).Err()
	return
}

// 将指定member的坐标转为hash字符串形式
func (rl *RedisLocation) GeoHash() (hash []string, err error) {
	hash, err = db.RedisClient.GeoHash("locations", "New York").Result()
	return
}

// 返回指定member的坐标
func (rl *RedisLocation) GeoCoordinates() (coordinates []*redis.GeoPos, err error) {
	coordinates, err = db.RedisClient.GeoPos("locations", "New York").Result()
	return
}

// 计算两个点之间的距离
func (rl *RedisLocation) GeoDistance() (distance float64, err error) {
	distance, err = db.RedisClient.GeoDist("locations", "New York", "London", "km").Result()
	return
}
