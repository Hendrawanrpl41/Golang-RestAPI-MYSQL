package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"belajar/RestApi/model"
)

var Db, Err = gorm.Open("mysql", "root:@/golang")

func init(){
	if Err != nil{
		panic("Fail Connection Database")
	}

	//model declaration
	var accounts model.Accounts
	var cards model.Cards

	//migrate
	Db.AutoMigrate(&accounts)
	Db.AutoMigrate(&cards)
}