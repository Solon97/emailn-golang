package internalerrors

import "errors"

// Server Errors
var ErrInternalServer = errors.New("internal server error")

// Service Errors
var ErrRepositoryNil = errors.New("repository cannot be nil")

// Validation Errors Patterns
const (
	ErrRequiredFieldPattern = "%s is required"
	ErrMinFieldPattern      = "%s must be at least %s"
	ErrMaxFieldPattern      = "%s must be at most %s"
	ErrEmailFieldPattern    = "%s must be a valid email"
)
