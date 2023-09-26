package entity

import (
	"errors"

	"github.com/felipefabricio/golang-rest-api/pkg/entity"
)

var (
	ErrIdInvalid     = errors.New("id is invalid")
	ErrIdRequired    = errors.New("id is required")
	ErrNameRequired  = errors.New("name is required")
	ErrPriceRequired = errors.New("price is required")
	ErrPriceInvalid  = errors.New("price is invalid")
)

type Product struct {
	ID          entity.ID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	CreatedAt   string    `json:"created_at"`
}

func (p *Product) NewProduct(name, description string, price int) (*Product, error) {
	return &Product{
		ID:          entity.NewID(),
		Name:        name,
		Description: description,
		Price:       price,
	}, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIdInvalid
	}
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.Price <= 0 {
		return ErrPriceInvalid
	}
	if p.Price == 0 {
		return ErrPriceRequired
	}
	return nil
}
