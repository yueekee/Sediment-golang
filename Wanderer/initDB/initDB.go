package initDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *sql.DB

func init() {
	gorm.Open()

}
