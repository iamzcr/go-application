package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"pwd"`
	DB       int    `yaml:"db"`
}

type CanalConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Example  string `yaml:"example"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Redis RedisConfig `yaml:"redis"`
	Canal CanalConfig `yaml:"canal"`
}

var AppConfig = &Config{}

func InitConfig() (err error) {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("无法获取当前工作目录：%v", err)
	}
	// 打开文件
	file, err := os.Open(filepath.Join(dir, "/conf/config.yaml"))
	if err != nil {
		log.Fatalf("无法打开文件：%v", err)
	}
	defer file.Close() // 在函数结束时关闭文件

	// 读取文件内容
	yamlFile, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("无法读取文件：%v", err)
	}

	err = yaml.Unmarshal(yamlFile, &AppConfig)
	fmt.Println(AppConfig)
	if err != nil {
		log.Fatalf("无法解析YAML文件：%v", err)
	}
	return
}
