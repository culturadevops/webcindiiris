package model

import (
	"errors"
	"inkafarma/webcindi/libs"
	"log"

	"github.com/jinzhu/gorm"
)

type EnvModel struct {
	gorm.Model

	Name    string `gorm:"type:varchar(10);not null;"`
	Project string `gorm:"type:varchar(20);not null;"`
	Preprd  string `gorm:"type:varchar(2);not null;"`
}

func (EnvModel) TableName() string {
	return "env"
}
func (this *EnvModel) GetByName(name string) (EnvModel, error) {
	var data = EnvModel{}

	if libs.DB.Where("name  = ? ", name).Find(&data).RecordNotFound() {
		return EnvModel{}, errors.New("分类未找到")
	}
	return data, nil
}

func (this *EnvModel) Get(name string, project string) (EnvModel, error) {
	var data = EnvModel{}

	if libs.DB.Where("name  = ? AND project = ?", name, project).Find(&data).RecordNotFound() {
		return EnvModel{}, errors.New("分类未找到")
	}
	return data, nil
}
func (this *EnvModel) ListByProject(project string) []EnvModel {
	var data = []EnvModel{}

	err := libs.DB.Where("project = ? ", project).Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
func (this *EnvModel) ListByPreprd(preprd string) []EnvModel {
	var data = []EnvModel{}

	err := libs.DB.Where("preprd = ? ", preprd).Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
func (this *EnvModel) ListAll() []EnvModel {
	var data = []EnvModel{}

	err := libs.DB.Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
