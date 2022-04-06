package repository

import (
	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/adapter/infrastructure"
	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/adapter/infrastructure/model/primary"
	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/domain/model"
	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/domain/repository"
	"github.com/volatiletech/sqlboiler/boil"
)

var _ repository.User = (*user)(nil)

type user struct {
	connector infrastructure.RDBConnector
}

func NewUser(
	connector infrastructure.RDBConnector,
) repository.User {
	return &user{
		connector: connector,
	}
}

func (r *user) Create(user *model.User) (*model.User, error) {
	record := user.UserToRecord()

	if err := record.Insert(
		r.connector.GetContext(),
		r.connector.GetDB1(),
		boil.Infer(),
	); err != nil {
		return nil, err
	}

	if err := record.Reload(
		r.connector.GetContext(),
		r.connector.GetDB1(),
	); err != nil {
		return nil, err
	}

	return model.RecordToUser(record), nil
}

func (r *user) FindByID(id int64) (*model.User, error) {
	record, err := primary.Users(
		primary.UserWhere.ID.EQ(id),
		primary.UserWhere.DeletedAt.IsNull(),
	).One(
		r.connector.GetContext(),
		r.connector.GetDB1(),
	)
	if err != nil {
		return nil, err
	}

	return model.RecordToUser(record), nil
}
