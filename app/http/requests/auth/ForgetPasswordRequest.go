// Http/Requests/forget_password_request.go

package authrequests

import "com/beeda/search/pkg/validator"

type ForgetPasswordRequest struct {
	// Add fields for initiating forget password process
	*validator.Request
}
