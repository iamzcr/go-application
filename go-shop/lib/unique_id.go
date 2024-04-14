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
	// 将时间戳转换为time.Time类型
	t := time.Unix(endTimeStamp, 0)
	// 使用Go的格式化字符串转换为日期
	currentDay := t.Format("20060102")

	//时间差值秒
	diffSecond := endTimeStamp - beginTimeStamp

	// redis的自增操作生成id,keyType + ":" + currentDay按日期定义key，可以方便统计某下单数
	incrKey := fmt.Sprintf("%s:%s", keyType, currentDay)
	incrId, err := db.RedisClient.Incr(incrKey).Result()
	if err != nil {
		fmt.Println("redis Incr:", err)
		return
	}
	//拼接时间差值和redis生成的自增id
	str := strconv.FormatInt(diffSecond, 10) + strconv.FormatInt(incrId, 10)

	// 将字符串转换为int64
	uniqueId, err = strconv.ParseInt(str, 10, 64)

	if err != nil {
		fmt.Println("uniqueId:", err)
		return
	}
	return
}
