package valueobject

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName string, lastName string) *FullName {
	fullName := &FullName{
		firstName: firstName,
		lastName:  lastName,
	}
	return fullName
}

func (f FullName) Equals(fullName *FullName) bool {
	return f.firstName == fullName.firstName && f.lastName == fullName.lastName
}
