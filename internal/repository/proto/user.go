package proto

import (
	"fmt"
	"food-delivery/pkg/userservice"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(ctx context.Context, user *userservice.User) (*userservice.UserState, error) {

	createUserQuery := fmt.Sprintf(`INSERT INTO users (id, name, address, phone_number, birthday) VALUES ($1, $2, $3, $4, $5)`)
	if _, err := r.db.Exec(createUserQuery, user.UUID, user.Name, user.Address, user.PhoneNumber, user.Birthday.AsTime()); err != nil {
		log.Error().Err(err).Msg("error occurred while create user")

		return &userservice.UserState{Success: false, Error: err.Error() }, err
	}

	return &userservice.UserState{Success: true}, nil
}
