package services

import (
	"test-flow-code-gen/examples/db"
	"test-flow-code-gen/examples/models"
)

type UserService struct {
	DB *db.MockDB
}

func NewUserService() *UserService {
	return &UserService{
		DB: &db.MockDB{},
	}
}

func (s *UserService) Create(user models.User) (int, error) {
	userID := 5
	return userID, nil
}

func (s *UserService) Update(user models.User) error {
	return nil
}

func (s *UserService) Delete() error {
	return nil
}

func (s *UserService) Find() ([]models.User, error) {
	return []models.User{
		{
			ID:   1,
			Name: "John",
			Age:  23,
		},
		{
			ID:   2,
			Name: "Paul",
			Age:  26,
		},
		{
			ID:   3,
			Name: "Marry",
			Age:  29,
		},
		{
			ID:   5,
			Name: "Lora",
			Age:  39,
		},
	}, nil
}
