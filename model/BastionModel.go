package model

import (
	"errors"
	"inkafarma/webcindi/libs"
	"log"

	"github.com/jinzhu/gorm"
)

type BastionModel struct {
	gorm.Model

	Ip  string `gorm:"type:varchar(250);not null;"`
	Env uint   `gorm:"type:uint;not null;"`
}

func (BastionModel) TableName() string {
	return "bastion"
}

func (this *BastionModel) Get(env uint) (BastionModel, error) {
	var data = BastionModel{}

	if libs.DB.Where("env  = ? ", env).Find(&data).RecordNotFound() {
		return BastionModel{}, errors.New("分类未找到")
	}
	return data, nil
}
func (this *BastionModel) ListEnv(env uint) []BastionModel {
	var data = []BastionModel{}

	err := libs.DB.Where("env = ? ", env).Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
func (this *BastionModel) ListAll() []BastionModel {
	var data = []BastionModel{}
	err := libs.DB.Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
