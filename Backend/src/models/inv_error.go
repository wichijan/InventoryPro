package models

type INVErrorMessage struct {
	ErrorMessage string `json:"errorMessage"`
}

type INVError struct {
	INVErrorMessage
	Status  int    `json:"-"`
	Details string `json:"details,omitempty"`
}

func (err *INVError) WithDetails(details string) *INVError {
	newErr := new(INVError)
	*newErr = *err
	newErr.Details = details
	return newErr
}
