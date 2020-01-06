package connect

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Data *gorm.DB

func init() {
	var err error
	Data, err = gorm.Open("mysql", "<account>:<password>@tcp(<mysql server ip>:3306)/<database>")
	if err != nil {
		panic(err)
	}
}
