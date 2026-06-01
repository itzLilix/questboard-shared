package dtos

import "time"

type SessionApplicationStatus string

const (
	ApplicationPending   SessionApplicationStatus = "pending"
	ApplicationAccepted  SessionApplicationStatus = "accepted"
	ApplicationRejected  SessionApplicationStatus = "rejected"
	ApplicationCancelled SessionApplicationStatus = "cancelled"
)

func (s SessionApplicationStatus) Valid() bool {
	switch s {
	case ApplicationPending, ApplicationAccepted, ApplicationRejected, ApplicationCancelled:
		return true
	}
	return false
}

func (s *SessionApplicationStatus) UnmarshalText(text []byte) error {
	return unmarshalEnum(s, text, SessionApplicationStatus.Valid, "application status")
}

type SessionApplication struct {
	ID          string                   `json:"id"`
	SessionID   string                   `json:"sessionId"`
	ApplicantID string                   `json:"applicantId"`
	Status      SessionApplicationStatus `json:"status"`
	Message     *string                  `json:"message,omitempty"`
	CreatedAt   time.Time                `json:"createdAt"`
	ResolvedAt  *time.Time               `json:"resolvedAt,omitempty"`
}
