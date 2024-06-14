package inv_errors

import "github.com/wichijan/InventoryPro/src/models"

var (
	// INV_BAD_REQUEST is used to indicate that the request was malformed
	INV_BAD_REQUEST = &models.INVError{INVErrorMessage: models.INVErrorMessage{ErrorMessage: "BAD_REQUEST"}, Status: 400}
	// INV_UNAUTHORIZED is used to indicate that the request was unauthorized
	INV_UNAUTHORIZED = &models.INVError{INVErrorMessage: models.INVErrorMessage{ErrorMessage: "UNAUTHORIZED"}, Status: 401}
	// INV_CREDENTIALS_INVALID is used to indicate that the login credentials were invalid
	INV_CREDENTIALS_INVALID = &models.INVError{INVErrorMessage: models.INVErrorMessage{ErrorMessage: "CREDENTIALS_INVALID"}, Status: 401}
	// INV_FORBIDDEN is used to indicate that the request was forbidden due to insufficient permissions
	INV_FORBIDDEN = &models.INVError{INVErrorMessage: models.INVErrorMessage{ErrorMessage: "FORBIDDEN"}, Status: 403}
	// INV_USER_NOT_FOUND is used to indicate that the requested user was not found
	INV_USER_NOT_FOUND = &models.INVError{INVErrorMessage: models.INVErrorMessage{ErrorMessage: "USER_NOT_FOUND"}, Status: 404}
	// INV_NOT_FOUND is used to indicate that the requested resource was not found
	INV_NOT_FOUND = &models.INVError{INVErrorMessage: models.INVErrorMessage{ErrorMessage: "NOT_FOUND"}, Status: 404}
	// INV_CONFLICT is used to indicate that the request could not be processed due to a conflict
	INV_CONFLICT = &models.INVError{INVErrorMessage: models.INVErrorMessage{ErrorMessage: "CONFLICT"}, Status: 409}
	// INV_UPSTREAM_ERROR is used to indicate an error in 3rd party services
	INV_UPSTREAM_ERROR = &models.INVError{INVErrorMessage: models.INVErrorMessage{ErrorMessage: "UPSTREAM_ERROR"}, Status: 500}
	// INV_INTERNAL_ERROR is used to indicate an internal, unclassified error
	INV_INTERNAL_ERROR = &models.INVError{INVErrorMessage: models.INVErrorMessage{ErrorMessage: "INTERNAL_ERROR"}, Status: 500}
)
