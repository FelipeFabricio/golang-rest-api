package entity

import (
	"github.com/felipefabricio/golang-rest-api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type Client struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (c *Client) NewClient(name, email, password string) (*Client, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Client{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (c *Client) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(password))
	return err == nil
}
