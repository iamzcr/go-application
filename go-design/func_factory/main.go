// 工厂方法模式
package main

import "fmt"

// Database 接口
type Database interface {
	Connect()
}

// DatabaseFactory 数据库工厂接口
type DatabaseFactory interface {
	CreateDatabase() Database
}

// MySQLDatabase MySQL数据库实现
type MySQLDatabase struct{}

func (m *MySQLDatabase) Connect() {
	fmt.Println("连接到MySQL数据库")
	// 连接MySQL数据库的逻辑
}

// PostgreSQLDatabase PostgreSQL数据库实现
type PostgreSQLDatabase struct{}

func (p *PostgreSQLDatabase) Connect() {
	fmt.Println("连接到PostgreSQL数据库")
	// 连接PostgreSQL数据库的逻辑
}

// MySQLDatabaseFactory MySQL数据库工厂
type MySQLDatabaseFactory struct{}

func (f *MySQLDatabaseFactory) CreateDatabase() Database {
	return &MySQLDatabase{}
}

// PostgreSQLDatabaseFactory PostgreSQL数据库工厂
type PostgreSQLDatabaseFactory struct{}

func (f *PostgreSQLDatabaseFactory) CreateDatabase() Database {
	var database Database
	//创建一个具体的PostgreSQLDatabase
	database = new(PostgreSQLDatabase)
	return database
}

func main() {
	// 创建MySQL数据库工厂
	mysqlFactory := &MySQLDatabaseFactory{}
	mysqlDb := mysqlFactory.CreateDatabase()
	mysqlDb.Connect()

	// 创建PostgreSQL数据库的工厂
	var postgresFactory DatabaseFactory
	postgresFactory = new(PostgreSQLDatabaseFactory)

	//创建一个具体的PostgreSQL数据库
	var postgresDb Database
	postgresDb = postgresFactory.CreateDatabase()
	postgresDb.Connect()
}
