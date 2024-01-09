package services

import "com/beeda/search/app/models"

type JWTService interface {
	GenerateToken(user *models.User) (string, error)
}
