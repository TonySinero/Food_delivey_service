package repository

import (
	"database/sql"
	"fmt"
	"food-delivery/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

type CustomerPostgres struct {
	db *sqlx.DB
}

func NewCustomerPostgres(db *sqlx.DB) *CustomerPostgres {
	return &CustomerPostgres{db: db}
}

func (r *CustomerPostgres) CreateCustomer(input domain.Customer) error {

	createCustomerQuery := fmt.Sprintf(`INSERT INTO customers(id, full_name, address, phone_number)
			VALUES ($1, $2, $3, $4) RETURNING id`)
	var id string
	row := r.db.QueryRow(createCustomerQuery, input.ID, input.FullName, input.Address, input.PhoneNumber)

	if err := row.Scan(&id); err != nil {
		log.Error().Err(err).Msg("error occurred while creating customer")
		return err
	}

	return nil
}

func (r *CustomerPostgres) GetCustomerByID(id string) (domain.CustomerInfo, error) {

	var customer domain.CustomerInfo

	SelectCustomerQuery := fmt.Sprintf(`SELECT full_name, address, phone_number, birthday 
		FROM customers WHERE id = $1`)

	if err := r.db.Get(&customer, SelectCustomerQuery, id); err != nil {
		log.Error().Err(err).Msg("error occurred while selecting customer")
		if err == sql.ErrNoRows {
			return customer, domain.ErrCustomerNotFound
		}
	}
	return customer, nil
}

func (r *CustomerPostgres) UpdateCustomer(input domain.CustomerUpdate, id string) error {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.FullName != nil {
		setValues = append(setValues, fmt.Sprintf("full_name=$%d", argID))
		args = append(args, *input.FullName)
		argID++
	}

	if input.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argID))
		args = append(args, *input.Address)
		argID++
	}

	if input.PhoneNumber != nil {
		setValues = append(setValues, fmt.Sprintf("phone_number=$%d", argID))
		args = append(args, *input.PhoneNumber)
		argID++
	}

	if input.BirthDay != nil {
		birthday, err := time.Parse("2006-01-02", *input.BirthDay)
		if err != nil {
			log.Error().Err(err).Msg("invalid date")
			return err
		}
		setValues = append(setValues, fmt.Sprintf("birthday=$%d", argID))
		args = append(args, birthday)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	updateCustomerQuery := fmt.Sprintf("UPDATE customers SET %s WHERE id=$%d", setQuery, argID)
	args = append(args, id)

	if _, err := r.db.Exec(updateCustomerQuery, args...); err != nil {
		log.Error().Err(err).Msg("error occurred while updating customer")
		return err
	}

	return nil
}
