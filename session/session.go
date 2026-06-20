// Package session re-exports the session storage interface and controls.
//
// Session storage provides secure, backend-agnostic server-side session
// management with size limits, JSON validation, and observability hooks.
package session

import (
	"github.com/plexusone/omnistorage-core/session"
)

// Re-export types.
type (
	Session          = session.Session
	Store            = session.Store
	SessionLister    = session.SessionLister
	Config           = session.Config
	ControlledStore  = session.ControlledStore
	ControlOption    = session.ControlOption
	ViolationType    = session.ViolationType
	ViolationEvent   = session.ViolationEvent
	ViolationHandler = session.ViolationHandler
	SessionError     = session.SessionError
)

// Re-export errors.
var (
	ErrNotFound             = session.ErrNotFound
	ErrExpired              = session.ErrExpired
	ErrInvalidSession       = session.ErrInvalidSession
	ErrSizeLimitExceeded    = session.ErrSizeLimitExceeded
	ErrSessionLimitExceeded = session.ErrSessionLimitExceeded
	ErrSiteMismatch         = session.ErrSiteMismatch
	ErrClosed               = session.ErrClosed
)

// Re-export error codes.
const (
	ErrCodeInvalidSession       = session.ErrCodeInvalidSession
	ErrCodeSizeLimitExceeded    = session.ErrCodeSizeLimitExceeded
	ErrCodeNotSerializable      = session.ErrCodeNotSerializable
	ErrCodeNotFound             = session.ErrCodeNotFound
	ErrCodeExpired              = session.ErrCodeExpired
	ErrCodeStoreClosed          = session.ErrCodeStoreClosed
	ErrCodeSessionLimitExceeded = session.ErrCodeSessionLimitExceeded
	ErrCodeSiteMismatch         = session.ErrCodeSiteMismatch
)

// Re-export violation types.
const (
	ViolationSizeExceeded         = session.ViolationSizeExceeded
	ViolationNotSerializable      = session.ViolationNotSerializable
	ViolationInvalidSession       = session.ViolationInvalidSession
	ViolationSessionLimitExceeded = session.ViolationSessionLimitExceeded
	ViolationSiteMismatch         = session.ViolationSiteMismatch
)

// Re-export functions.
var (
	NewSession           = session.NewSession
	GenerateSessionID    = session.GenerateSessionID
	DefaultConfig        = session.DefaultConfig
	WithControls         = session.WithControls
	WithLogger           = session.WithLogger
	WithViolationHandler = session.WithViolationHandler
	NewSessionError      = session.NewSessionError
	ErrorCode            = session.ErrorCode
	ErrorDetails         = session.ErrorDetails
)

// Re-export constants.
const SessionIDLength = session.SessionIDLength
