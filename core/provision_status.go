package core

import "time"

const (
	// ProvStatusUndefined represents an empty provision state
	ProvStatusUndefined ProvisionStatus = ""

	// ProvStatusAwaiting represents a provisionable object that has been indexed, but not implemented
	ProvStatusAwaiting ProvisionStatus = "AWAITING"

	// ProvStatusInProgress represents a provisioanble object that has been indexed ans is in the process of being implemented.
	ProvStatusInProgress ProvisionStatus = "INPROGRESS"

	// ProvStatusFailed represents a provisioanble object that failed during it's implementation.
	ProvStatusFailed ProvisionStatus = "FAILED"

	// ProvStatusComplete represents a provisionable object that has successfully been implemented.
	ProvStatusComplete ProvisionStatus = "COMPLETE"

	// ProvStatusTainted represents a provisionable object that has been implemented, but is marked for re-implementation.
	ProvStatusTainted ProvisionStatus = "TAINTED"
)

// ProvisionStatus describes components of a provisioned environment and the various states they're in
type ProvisionStatus string

// Status represents the state of an individual object that could be provisioned within a laforge team environment
//easyjson:json
type Status struct {
	State     ProvisionStatus `json:"state"`
	StartedAt time.Time       `json:"started_at,omitempty"`
	EndedAt   time.Time       `json:"ended_at,omitempty"`
	Failed    bool            `json:"failed"`
	Completed bool            `json:"completed"`
	Error     string          `json:"error,omitempty"`
}

// Current returns the current status
func (s *Status) Current() ProvisionStatus {
	return s.State
}

// CanProceed is used to ensure the state allows further traversal
func (s *Status) CanProceed() bool {
	if s.State == ProvStatusComplete {
		return true
	}
	return false
}

// Hash implements the Hasher interface
func (s *Status) Hash() uint64 {
	switch s.State {
	case ProvStatusFailed:
		return uint64(666)
	case ProvStatusTainted:
		return uint64(999)
	default:
		return uint64(1)
	}
}
