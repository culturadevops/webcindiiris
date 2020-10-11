package model

import (
	"errors"
	"inkafarma/webcindi/libs"

	"github.com/jinzhu/gorm"
)

type UserEnvModel struct {
	gorm.Model
	User_id uint `gorm:"type:uint;not null;"`
	Env_id  uint `gorm:"type:uint;not null;"`
}

func (UserEnvModel) TableName() string {
	return "user_env"
}

func (this *UserEnvModel) Add(userid uint, envid uint) error {
	var varUser UserEnvModel
	varUser.User_id = userid
	varUser.Env_id = envid

	if !libs.DB.Where("user_id = ? AND env_id", userid, envid).First(&UserEnvModel{}).RecordNotFound() {
		return errors.New("Ya existe un item con el nombre ")
	}
	if err := libs.DB.Create(&varUser).Error; err != nil {
		return err
	}
	return nil
}
func (this *UserEnvModel) GetByUserId(userid uint) []UserEnvModel {
	var data = []UserEnvModel{}
	err := libs.DB.Where("user_id  = ? ", userid).Find(&data).Order("user_id desc")
	if err != nil {
		//log.Fatalln(err)
	}
	return data
}
