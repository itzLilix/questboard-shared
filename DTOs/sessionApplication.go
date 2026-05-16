package dtos

import "time"

type SessionApplicationStatus string

const (
	ApplicationPending   SessionApplicationStatus = "pending"
	ApplicationAccepted  SessionApplicationStatus = "accepted"
	ApplicationRejected  SessionApplicationStatus = "rejected"
	ApplicationCancelled SessionApplicationStatus = "cancelled"
)

type SessionApplication struct {
	ID          string                   `json:"id"`
	SessionID   string                   `json:"sessionId"`
	ApplicantID string                   `json:"applicantId"`
	Status      SessionApplicationStatus `json:"status"`
	Message     *string                  `json:"message,omitempty"`
	CreatedAt   time.Time                `json:"createdAt"`
	ResolvedAt  *time.Time               `json:"resolvedAt,omitempty"`
}
