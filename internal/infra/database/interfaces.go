package database

import (
	"github.com/felipefabricio/golang-rest-api/internal/entity"
)

type UserInterface interface {
	Create(client *entity.Client) error
	FindByEmail(email string) (*entity.Client, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAllProducts(page, limit int, sort string) (*[]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
