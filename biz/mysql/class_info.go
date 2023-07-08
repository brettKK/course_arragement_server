package mysql

import (
	"context"
	"log"
	"time"
)

// ClassInfo 班级信息
type ClassInfo struct {
	ID int `gorm:"column:id" json:"id"`
	ClassNo string `gorm:"column:class_no" json:"class_no"`
	ClassName string `gorm:"column:class_name" json:"class_name"`
	Num int `gorm:"column:num" json:"num"`
	Teacher string `gorm:"column:teacher" json:"teacher"`
	Remark string `gorm:"column:remark" json:"remark"`
	Deleted int `gorm:"column:deleted" json:"deleted"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (h *ClassInfo) TableName() string {
	return "tb_class_info"
}

// GetClassInfo 获取全部的班级信息
func GetClassInfo(ctx context.Context) ([]ClassInfo, error) {
	var rows []ClassInfo
	if err := dbCli.WithContext(ctx).Model(&ClassInfo{}).Find(&rows).Error; err != nil {
		log.Printf("GetClassInfo Error:%+v", err)
		return nil, err
	}
	return rows, nil
}