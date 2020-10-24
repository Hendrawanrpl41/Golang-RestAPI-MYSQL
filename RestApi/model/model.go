package model

type Account struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
}

type Accounts []Account

type Card struct{
	ID int `json:"id"`
	Pin int `json:"pin"`
	AccountId int 
	Account Accounts `gorm:"foreignKey:AccountId;references:ID"`
}

type Cards []Card


