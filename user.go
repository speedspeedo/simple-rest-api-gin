package main

type User struct {
	Id int `json:"id"`

	Name string `json:"name"`

	Password string `json:"password"`

	Email string `json:"email"`
}

/**
In GO exported variable and function Name
must start with big cap
*/

func SampleUser() []User {
	var users = []User{
		{Id: 1, Name: "Randy Cummings", Password: "123456", Email: "speedodeveloper1004@gmail.com"},
		{Id: 2, Name: "Andy Cummings", Password: "123456", Email: "speedodeveloper1004@gmail.com"},
		{Id: 3, Name: "Pandy Cummings", Password: "123456", Email: "speedodeveloper1004@gmail.com"},
	}

	return users
}
