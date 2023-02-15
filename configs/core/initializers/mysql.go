package initializers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	Connect() error
	Disconnect() error
}

type MySQL struct {
	// url string
	db *sql.DB
}

var _ DB = (*MySQL)(nil)

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db}
}

func (m *MySQL) Connect() error {
	if err := m.db.Ping(); err != nil {
		return err
	}
	// boil.SetDB(m.db)
	// boil.SetLocation(utils.TimeZone())
	return nil
}

func (m *MySQL) Disconnect() error {
	return m.db.Close()
}
