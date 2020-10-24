package controller

// import(
// 	"belajar/RestApi/model"
// )

// type Accounts model.Accounts
//interface
type ActionModel interface{
	Create() interface{}
	FindAll() interface{}
	Update() interface{}
	Delete() interface{}
}