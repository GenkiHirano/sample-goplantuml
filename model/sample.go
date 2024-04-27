package model

type User struct {
	ID      int
	Name    Name
	Address Address
}

type Name struct {
	First string
	Last  string
}

func (n Name) FullName() string {
	return n.First + n.Last
}

type Address struct {
	PostalCode string
	Prefecture string
	City       string
}
