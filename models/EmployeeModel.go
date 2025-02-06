package models

type Employee struct {
	First_name   string
	Last_name    string
	Company_name string
	Address      string
	City         string
	Country      string
	Postal       string
	Phone        string
	Email        string
	Web          string
}
type Request struct {
	Companyname string `json:companyname`
	Firstname   string `json:firstname`
	Email       string `json:email`
}
