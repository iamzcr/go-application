package lib

import (
	"fmt"
	"go-shop/db"
	"strconv"
	"time"
)

var beginTimeStamp int64 = 1672502400

func RedisUniqueId(keyType string) (uniqueId int64, err error) {
	//生成当前和某个时间的差值
	endTimeStamp := time.Now().Unix()
	diffSecond := endTimeStamp - beginTimeStamp

	// redis的自增操作生成id
	result, err := db.RedisClient.Incr(keyType).Result()
	if err != nil {
		fmt.Println("redis Incr:", err)
		return
	}
	//拼接时间插值和redis生成的自增id
	str := strconv.FormatInt(diffSecond, 10) + strconv.FormatInt(result, 10)

	// 将字符串转换为int64
	uniqueId, err = strconv.ParseInt(str, 10, 64)

	if err != nil {
		fmt.Println("uniqueId:", err)
		return
	}
	return
}
