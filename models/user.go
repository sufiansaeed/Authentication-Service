package models

import (
	"github.com/fatih/structs"
)

type User struct {
	ID				string	`json:"id" structs:"id" db:"id"`
	FatherName		string	`json:"father_name" structs:"father_name" db:"father_name"`
	FirstName		string	`json:"first_name" structs:"first_name" db:"first_name"`
	LastName		string	`json:"last_name" structs:"last_name" db:"last_name"`
	Gender			string	`json:"gender" structs:"gender" db:"gender"`
	CNIC			string	 `json:"cnic" structs:"cnic" db:"cnic"`
	Email			string	`json:"email" structs:"email" db:"email"`
	Password		string	`json:"password" structs:"password" db:"password"`
}

// Map function returns map values.
func (u *User) Map() map[string]interface{} {
	return structs.Map(u)
}

// Names function returns field names.
func (u *User) Names() []string {
	fields := structs.Fields(u)
	names := make([]string, len(fields))
	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}
	return names
}
