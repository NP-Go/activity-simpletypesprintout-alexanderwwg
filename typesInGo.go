package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	age              int
	weight           float32
	height           float32
	remainingCredits float64
	accountName      string
	accountPassword  string
	subscribedUser   bool
}

//Insert variables declarations here based on activity

func assignTypesForLuke() *person {
	p := person{firstName: "Luke"}
	p.lastName = "Skywalker"
	p.age = 20
	p.weight = 73.0
	p.height = 1.72
	p.remainingCredits = 123.55
	p.accountName = "admin"
	p.accountPassword = "password"
	p.subscribedUser = true

	return &p
}

func tellMeTypes(p *person) {
	//insert code here to print out types of variables
	text := "The following is the account information."

	fmt.Printf("%T %T %T %T %T %T %T %T %T %T", text, p.firstName,
		p.lastName, p.age, p.weight, p.height, p.remainingCredits,
		p.accountName, p.accountPassword, p.subscribedUser)
}

func main() {
	luke := assignTypesForLuke()
	tellMeTypes(luke)
}
