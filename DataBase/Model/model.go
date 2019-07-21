package model

import (
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Test struct {
  Name string `json:"name"`
}
