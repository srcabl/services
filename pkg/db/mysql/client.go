package mysql

import (
	"database/sql"
	"fmt"

	mysql_db "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/srcabl/services/pkg/config"

	// migration drivers
	_ "github.com/golang-migrate/migrate/source/file"
)

// Client is the mysql client
type Client struct {
	DB *sql.DB

	dbConfig *mysql_db.Config
}

// New news up a mysql client
func New(config *config.Service) (*Client, error) {
	return &Client{
		dbConfig: ToMYSQLConfig(config),
	}, nil
}

// Connect connects to the db
func (c *Client) Connect() (func() error, error) {
	fmt.Print("CONNECTING TO DB")
	conn, err := mysql_db.NewConnector(c.dbConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db connector")
	}

	c.DB = sql.OpenDB(conn)

	//return with the close function to defer
	return c.Close, nil
}

// Close closes the connection to the db
func (c *Client) Close() error {
	err := c.DB.Close()
	if err != nil {
		return errors.Wrap(err, "failed to close the connection")
	}

	return nil
}
