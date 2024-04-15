package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pbe "github.com/withlin/canal-go/protocol/entry"
	"go-canal/canal"
	"go-canal/conf"
	"go-canal/db"
	"log"
	"os"
	"time"
)

func main() {
	//初始化配置
	err := conf.InitConfig()
	if err != nil {
		fmt.Println("config init:", err)
	}
	//初始化redis链接客户端
	err = db.InitRedis()
	if err != nil {
		fmt.Println("redis init:", err)
	}
	//初始化canal链接客户端
	err = canal.InitCanal()
	if err != nil {
		fmt.Println("redis init:", err)
	}
	//订阅要监听的库，表示test库的所有表
	err = canal.CanalClient.Subscribe("test.*")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {
		message, err := canal.CanalClient.Get(100, nil, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(3000 * time.Millisecond)
			fmt.Println("===no data===")
			continue
		}
		for _, entry := range message.Entries {
			if entry.GetEntryType() == pbe.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == pbe.EntryType_TRANSACTIONEND {
				continue
			}
			rowChange := new(pbe.RowChange)
			err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
			checkError(err)
			eventType := rowChange.GetEventType()
			header := entry.GetHeader()
			fmt.Println(fmt.Sprintf("binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))
			for _, rowData := range rowChange.GetRowDatas() {
				if eventType == pbe.EventType_INSERT {
					printColumn(rowData.GetAfterColumns())
				}
			}

		}
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func printColumn(columns []*pbe.Column) {
	for _, col := range columns {
		err := db.RedisClient.Set("canal:"+col.GetName(), col.GetValue(), 100*time.Second).Err()
		if err != nil {
			fmt.Println("Failed to set key:", err)
			return
		}
	}
}
