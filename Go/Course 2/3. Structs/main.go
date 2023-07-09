package main

import (
	"fmt"
	"time"
)

type person struct {
	firstName string
	lastName  string
	birthDate time.Time
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func newContactInfo(email string, zip int) contactInfo {
	contact := contactInfo{
		email: email,
		zip:   zip,
	}

	return contact
}

func newPerson(firstName string, lastName string, birthDate time.Time, contactInfo contactInfo) person {
	person := person{
		firstName:   firstName,
		lastName:    lastName,
		birthDate:   birthDate,
		contactInfo: contactInfo,
	}

	return person
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) age() int {
	age := (time.Since(p.birthDate).Hours() / (365 * 24))
	return int(age)
}

func main() {
	birthDate := time.Date(1990, time.September, 29, 0, 0, 0, 0, time.Local)
	contactInfo := newContactInfo("emma90@gmail.com", 8843876)
	emma := newPerson("Ema", "Nakamura", birthDate, contactInfo)
	emma.updateFirstName("Emma")
	emma.print()
}
