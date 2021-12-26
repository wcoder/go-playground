package models

type Account struct {
	FirstName string
	LastName  string
}

func (a *Account) SetFirstName(newValue string) {
	a.FirstName = newValue
}

func (a *Account) SetLastName(newValue string) {
	a.LastName = newValue
}
