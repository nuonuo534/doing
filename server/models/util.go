package models

import (
	"fmt"
)

type Models struct {
	ID uint `gorm:"primaryKey";json:"id"`
}

func (m Models) ExistTableByKey(key string, value interface{}) bool {
	db.Select("id").Where(fmt.Sprintf("%s = ?", key), value).First(&m)
	if m.ID > 0 {
		return true
	}
	return false
}
