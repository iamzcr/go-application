// 简单工厂模式
package main

import "fmt"

// Database 接口
type Database interface {
	Connect()
}

// MySQLDatabase MySQL数据库实现
type MySQLDatabase struct {
}

func (m *MySQLDatabase) Connect() {
	fmt.Println("连接到MySQL数据库")
	// 连接MySQL数据库的逻辑
}

// PostgreDatabase Postgre数据库实现
type PostgreDatabase struct{}

func (p *PostgreDatabase) Connect() {
	fmt.Println("连接到Postgre数据库")
	// 连接Postgre数据库的逻辑
}

// DatabaseFactory 简单工厂类
type DatabaseFactory struct{}

func (f *DatabaseFactory) CreateDatabase(dbType string) Database {
	switch dbType {
	case "MySQL":
		return &MySQLDatabase{}
	case "postgre":
		return &PostgreDatabase{}
	default:
		panic("不支持的数据库类型")
	}
}

func main() {
	// 创建简单工厂对象
	factory := &DatabaseFactory{}

	// 创建MySQL数据库实例
	mysqlDb := factory.CreateDatabase("MySQL")
	mysqlDb.Connect()

	// 创建Oracle数据库实例
	oracleDb := factory.CreateDatabase("postgre")
	oracleDb.Connect()
}
