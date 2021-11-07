package config

import "fmt"

const (
	DBHost     = "0.0.0.0"
	DBPort     = "5432"
	DBUser     = "postgres"
	DBName     = "myblog"
	DBPassword = "postgres"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	database := fmt.Sprintf("host=%s port =%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return database
}
