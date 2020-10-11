package model

import (
	"inkafarma/webcindi/libs"
	"log"
	"math"

	"github.com/jinzhu/gorm"
	config "github.com/spf13/viper"
)

type BasedatoModel struct {
	gorm.Model

	Ip        string `gorm:"type:varchar(250);not null;"`
	Env       uint   `gorm:"type:uint;not null;"`
	Ipinterno string `gorm:"type:varchar(250);not null;"`
	Port      string `gorm:"type:varchar(5);not null;"`
	Engine    string `gorm:"type:varchar(50);not null;"`
}

func (BasedatoModel) TableName() string {
	return "Basedato"
}

func (this *BasedatoModel) ListByEnvAndPage(env uint, page int) ([]BasedatoModel, int, int) {
	var data = []BasedatoModel{}
	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	offset := (page - 1) * limit

	err := libs.DB.Where("env = ? ", env).Find(&data).Offset(offset).Limit(limit).Order("id desc").Count(&totalCount)
	if err != nil {
		//log.Fatalln(err)
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return data, totalCount, totalPages
}

func (this *BasedatoModel) ListByEnv(env uint) []BastionModel {
	var data = []BastionModel{}

	err := libs.DB.Where("env = ? ", env).Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
