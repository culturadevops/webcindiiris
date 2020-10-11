package model

import (
	"errors"
	"inkafarma/webcindi/libs"
	"log"

	"github.com/jinzhu/gorm"
)

type JenkinsModel struct {
	gorm.Model

	Ip       string `gorm:"type:varchar(250);not null;"`
	Account  string `gorm:"type:varchar(50);not null;"`
	Password string `gorm:"type:varchar(100);not null;"`
	Env      uint   `gorm:"type:uint;not null;"`
}

func (JenkinsModel) TableName() string {
	return "jenkins"
}
func (this *JenkinsModel) Get(env uint) (JenkinsModel, error) {
	var data = JenkinsModel{}

	if libs.DB.Where("env  = ? ", env).Find(&data).RecordNotFound() {
		return JenkinsModel{}, errors.New("分类未找到")
	}
	return data, nil
}

func (this *JenkinsModel) ListAll() []JenkinsModel {
	var data = []JenkinsModel{}
	err := libs.DB.Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
