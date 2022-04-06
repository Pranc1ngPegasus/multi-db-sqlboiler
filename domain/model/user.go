package model

import "github.com/Pranc1ngPegasus/multi-db-sqlboiler/adapter/infrastructure/model/primary"

type User struct {
	ID   int64
	Name string
}

func NewUser(
	name string,
) *User {
	return &User{
		Name: name,
	}
}

func RecordToUser(record *primary.User) *User {
	return &User{
		ID:   record.ID,
		Name: record.Name,
	}
}

func (m *User) UserToRecord() *primary.User {
	return &primary.User{
		ID:   m.ID,
		Name: m.Name,
	}
}
