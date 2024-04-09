package tail_logs

import (
	"github.com/hpcloud/tail"
	"go-kafka/product"
	"log"
)

var tailLogsObj *tail.Tail

func InitTail(flieName string) (err error) {
	config := tail.Config{
		ReOpen: true, //重新打开
		Follow: true, //是否跟随
		//Offset：它是 tail.SeekInfo 结构体中的一个字段，表示从哪个位置开始读取日志文件
		//Whence：它是 tail.SeekInfo 结构体中的另一个字段，表示从哪个位置基准进行偏移量的计算,0,1,2
		//0（io.SeekStart）：表示相对于文件的起始位置进行偏移量的计算。
		//1（io.SeekCurrent）：表示相对于当前位置进行偏移量的计算。
		//2（io.SeekEnd）：表示相对于文件的结束位置进行偏移量的计算。
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false, //文件不存在不报错
		Poll:      true,
	}

	tailLogsObj, err = tail.TailFile(flieName, config)
	if err != nil {
		log.Fatal(err)
	}
	return
}
func GetLogs() {
	select {
	case line, _ := <-tailLogsObj.Lines:
		product.SendMsgToKafka("error_log", line.Text)
	default:
		// 没有新的日志行可读取
		// 可以在此处执行其他操作
	}
}
