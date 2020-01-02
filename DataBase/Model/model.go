
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Tests struct {
	Name string `json:"name"`
}
