package mysql

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/srcabl/services/pkg/config"
)

// ToMYSQLConfig creates a mysql configfrom the environment config
func ToMYSQLConfig(c *config.Service) *mysql.Config {
	dbConfig := mysql.NewConfig()
	dbConfig.User = c.Database.User
	dbConfig.Passwd = c.Database.Password
	dbConfig.Net = "tcp"
	dbConfig.Addr = fmt.Sprintf("%s:%s", c.Database.Address, c.Database.Port)
	dbConfig.DBName = c.Database.Name
	return dbConfig
}
