// Http/Requests/update_profile_request.go

package authrequests

import "com/beeda/search/pkg/validator"

type UpdateProfileRequest struct {
	// Add fields for updating the user profile
	*validator.Request
}
