package dtos

import "time"

// MessagePayload is a fully resolved chat message, as returned by
// Persister.SaveMessage/EditMessage and broadcast over the socket. Also a
// reasonable shape for a REST message-history endpoint to return, since
// it's not tied to the websocket envelope.
type MessagePayload struct {
	ID               string        `json:"id"`
	SenderID         string        `json:"senderId"`
	Body             string        `json:"body"`
	ReplyTo          *ReplySnippet `json:"replyTo,omitempty"`
	Attachments      []Attachment  `json:"attachments,omitempty"`
	MentionedUserIDs []string      `json:"mentionedUserIds,omitempty"`
	CreatedAt        time.Time     `json:"createdAt"`
	EditedAt         *time.Time    `json:"editedAt,omitempty"`
	ClientNonce      string        `json:"clientNonce,omitempty"` // only meaningful on the socket path; empty from REST
}

// ReplySnippet is a small denormalized preview of the message being
// replied to, so clients can render "replying to: ..." without a second
// round trip. Resolved by Persister.SaveMessage. Messages are soft
// deleted (deleted_at) specifically so old replies stay structurally
// valid — if the target has deleted_at set, set Deleted true and leave
// ContentPreview empty rather than surfacing content the sender removed.
//
// Heads up: this is resolved once, when the reply is sent. If the
// original message is edited afterward, replies already broadcast keep
// showing the old preview text — same tradeoff most chat apps live with.
type ReplySnippet struct {
	MessageID      string `json:"messageId"`
	SenderID       string `json:"senderId"`
	ContentPreview string `json:"contentPreview,omitempty"`
	Deleted        bool   `json:"deleted,omitempty"`
}

// Attachment describes a file already uploaded through your existing
// upload endpoint before the message referencing it was sent — this is
// the resolved shape (DB-generated ID included).
type Attachment struct {
	ID        string `json:"id"`
	FileName  string `json:"fileName"`
	URL       string `json:"url"`
	MIMEType  string `json:"mimeType"`
	SizeBytes int64  `json:"sizeBytes"`
}

// AttachmentInput is what a client sends when attaching a file — the
// metadata your upload endpoint already returned, not the file itself.
// message_attachments generates its own id on insert, so none is sent
// here.
type AttachmentInput struct {
	FileName  string `json:"fileName"`
	URL       string `json:"url"`
	MIMEType  string `json:"mimeType,omitempty"`
	SizeBytes int64  `json:"sizeBytes,omitempty"`
}

// PinPayload is the resolved pin/unpin state for a message.
type PinPayload struct {
	MessageID string    `json:"messageId"`
	PinnedBy  string    `json:"pinnedBy,omitempty"`
	PinnedAt  time.Time `json:"pinnedAt,omitempty"`
	OrderIndex *int16 `json:"order_index,omitempty"`
}

// ReadPayload is a read watermark — "user X has read up through message
// Y" — not a per-message receipt. Message IDs are UUIDv7, so a plain
// string/byte comparison of LastReadMessageID against a message's own id
// sorts the same as chronological order — no separate sequence column
// needed for a "seen by" comparison.
type ReadPayload struct {
	UserID            string    `json:"userId"`
	LastReadMessageID string    `json:"lastReadMessageId"`
	ReadAt            time.Time `json:"readAt"`
}

type DeletePayload struct {
	MessageID string    `json:"messageId"`
}

type MessagePage struct {
	Messages   []MessagePayload `json:"messages"`
	NextCursor *string          `json:"nextCursor,omitempty"`
	HasMore    bool             `json:"hasMore"`
}