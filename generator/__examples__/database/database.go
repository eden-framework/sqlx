package database

import (
	"github.com/eden-framework/sqlx"
)

var DBTest = sqlx.NewDatabase("test")

type Gender int

const (
	GenderMale Gender = iota + 1
	GenderFemale
)

func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "male"
	case GenderFemale:
		return "female"
	}
	return ""
}
