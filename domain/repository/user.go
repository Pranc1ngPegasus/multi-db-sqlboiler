//go:generate mockgen -source $GOFILE -destination mock/$GOFILE -package=$GOPACKAGE

package repository

import (
	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/domain/model"
)

type User interface {
	Create(user *model.User) (*model.User, error)
	FindByID(id int64) (*model.User, error)
}
