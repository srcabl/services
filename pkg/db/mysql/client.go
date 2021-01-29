package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/srcabl/services/pkg/config"
)

type Client struct {
	db sql.DB
}

func New(config *config.Environment) (*Client, error) {

}
