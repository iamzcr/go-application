package canal

import (
	"github.com/withlin/canal-go/client"
	"go-canal/conf"
	"log"
)

var CanalClient *client.SimpleCanalConnector

func InitCanal() (err error) {
	CanalClient = client.NewSimpleCanalConnector(
		conf.AppConfig.Canal.Host,
		conf.AppConfig.Canal.Port,
		conf.AppConfig.Canal.Username,
		conf.AppConfig.Canal.Password,
		conf.AppConfig.Canal.Example,
		60000,
		60*60*1000)
	err = CanalClient.Connect()
	if err != nil {
		log.Println(err)
	}
	return
}
