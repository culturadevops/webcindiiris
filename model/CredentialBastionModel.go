package model

import (
	"errors"
	"fmt"
	"inkafarma/webcindi/libs"
	"math"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	config "github.com/spf13/viper"
)

type CredecentialBastionModel struct {
	gorm.Model
	User_id  uint   `gorm:"type:uint;not null;"`
	Account  string `gorm:"type:varchar(200);not null;"`
	Password string `gorm:"type:varchar(320);not null;"`
	Env      uint   `gorm:"type:uint;not null;"`
}
type Result struct {
	Id       uint
	IP       string
	Name     string
	Account  string
	Password string
}

func (CredecentialBastionModel) TableName() string {
	return "credecentialbastion"
}

var runes = []rune("abcdefg1234567890")

func generateRandomRune(n int) string {
	randRune := make([]rune, n)

	for i := range randRune {
		// without this, the final value will be same all the time.
		rand.Seed(time.Now().UnixNano())

		randRune[i] = runes[rand.Intn(len(runes))]
	}
	return string(randRune)
}

func (this *CredecentialBastionModel) GetById(id uint) (CredecentialBastionModel, error) {
	var data = CredecentialBastionModel{}

	if libs.DB.Where("id  = ? ", id).Find(&data).RecordNotFound() {
		return CredecentialBastionModel{}, errors.New("分类未找到")
	}
	return data, nil
}

func (this *CredecentialBastionModel) Add(userId uint, account string, password string, env uint) error {
	var varUser CredecentialBastionModel
	varUser.Account = account
	//varUser.Password = generateRandomRune(5)
	varUser.Password = password
	varUser.User_id = userId
	varUser.Env = env
	print("Creando credenciales ")
	/*if !libs.DB.Where("account = ? ", varUser.Account).First(&CredecentialBastionModel{}).RecordNotFound() {
		return errors.New("Ya existe un item con el nombre " + varUser.Account)
	}*/
	if err := libs.DB.Create(&varUser).Error; err != nil {
		return err
	}
	return nil
}
func (this *CredecentialBastionModel) List(user uint, page int) ([]Result, int, int) {
	//var data = []CredecentialBastionModel{}
	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	offset := (page - 1) * limit
	db := libs.DB
	fmt.Println(user)

	//err := db.Where("user_id = ?", user).Offset(offset).Limit(limit).Order("id desc").Find(&data).Count(&totalCount)
	var result []Result
	err := db.Table("credecentialbastion as c").Select("c.id,b.ip,e.name,c.account,c.password").Joins("join cindidevops.env e on c.env=e.id join cindidevops.bastion b on b.env=e.id ").Where("user_id = ?", user).Scan(&result).Offset(offset).Limit(limit).Order("id desc").Count(&totalCount)
	if err != nil {
		//log.Fatalln(err)
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return result, totalCount, totalPages
}
func (this *CredecentialBastionModel) PasswodUpdate(admin_id uint, password, Repassword string) error {
	if password == "" || Repassword == "" {
		return errors.New("密码不能为空")
	}
	if password != Repassword {
		return errors.New("密码不一致")
	}

	db := libs.DB
	var data CredecentialBastionModel

	if db.Where("id = ? ", admin_id).First(&data).RecordNotFound() {
		return errors.New("未查询到用户id")
	}

	if err := db.Model(&data).Update("password", password).Error; err != nil {
		return errors.New("密码修改失败")
	}

	return nil
}
