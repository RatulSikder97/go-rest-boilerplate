package authrequests

import "com/beeda/search/pkg/validator"

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=6"`
	*validator.Request
}
