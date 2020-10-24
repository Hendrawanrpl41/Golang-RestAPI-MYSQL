package controller

import(
	"fmt"
	"belajar/RestApi/model"
	"belajar/RestApi/config"
)

//declaration variabel model
type Account model.Account
type Accounts model.Accounts
type Card model.Card
type Cards model.Cards

//method implement interface

//store data
func (account Account) Create() interface{}{
	result := config.Db.Create(&account)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("step1")
	return result
}

//findAll
func (account Accounts) FindAll() interface{} {
	result := config.Db.Find(&account)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result)
	return result
}

//update
func (account Account) Update() interface{} {
	result := config.Db.Save(&account)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result)
	return result
}

//delete
func (account Account) Delete() interface{} {
	result := config.Db.Delete(&account, account.ID)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result)
	return result
}


//Cards
func (card Card) Create() interface{}{
	result := config.Db.Create(&card)
	if result.Error != nil {
		return result.Error
	}
	return result
}

//findAll
func (cards Cards) FindAll() interface{} {
	// var account Account
	// config.Db.Find(&account, cards[0].AccountId)
	fmt.Println()
	result := config.Db.Find(&cards)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result)
	return result
}

//update
func (card Card) Update() interface{} {
	result := config.Db.Save(&card)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result)
	return result
}

//delete
func (card Card) Delete() interface{} {
	result := config.Db.Delete(&card, card.ID)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result)
	return result
}




