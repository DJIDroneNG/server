package dao

import (
	"gorm.io/gorm"
)

// paginate ้็จๅ้กต
func paginate(page, pageSize *int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case *pageSize > 100:
			*pageSize = 100
		case *pageSize <= 0:
			*pageSize = 10
		}
		if *page < 1 {
			*page = 1
		}
		offset := (*page - 1) * (*pageSize)
		return db.Offset(offset).Limit(*pageSize)
	}
}
