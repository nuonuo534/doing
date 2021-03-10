package models

// Tag 表
type Tag struct {
	Models
	Name      string `gorm:"<-";validate:"required,len=20"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	CreatedBy string `gorm:"<-";validate:"required,len=20`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	UpdatedBy string `gorm:"<-";validate:"required,len=20`
	State     int
}

// GetTags 获取tag
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// GetTagTotal get total
func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// ExistTagByName isExist
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// AddTag create tag
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}
