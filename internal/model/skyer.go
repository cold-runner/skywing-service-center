package model

import "time"

type Applier struct {
	Num       int       `gorm:"column:stu_num"`
	Name      string    `gorm:"column:stu_name"`
	Major     string    `gorm:"column:major"`
	Class     string    `gorm:"column:class"`
	Province  string    `gorm:"column:province"`
	Introduce string    `gorm:"column:introduce"`
	Imagine   string    `gorm:"column:imagine"`
	Birth     time.Time `gorm:"birth"`
}

func (Applier) TableName() string {
	return "user"
}
