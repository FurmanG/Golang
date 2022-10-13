package member

import "fmt"

// Member data
type Member struct {
	firstName  string
	familyName string
	age        uint
}

// Set a memeber data
func (m *Member) SetaMemeber(first string, family string, age uint) *Member {
	m.firstName = first
	m.familyName = family
	m.age = age
	return &Member{
		firstName:  m.firstName,
		familyName: m.familyName,
		age:        m.age,
	}

}

// Get a member data
func (m *Member) GetaMemeber() {
	fmt.Println(m.firstName)
	fmt.Println(m.familyName)
	fmt.Println(m.age)
}
