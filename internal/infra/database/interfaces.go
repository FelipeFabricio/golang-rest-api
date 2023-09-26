package database

import "github.com/felipefabricio/golang-rest-api/internal/entity"

type UserInterface interface {
	Create(client *entity.Client) error
	FindByEmail(email string) (*entity.Client, error)
}
