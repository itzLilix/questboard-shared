package dtos

import "time"

type ChatRole string

const (
	ChatMemberRole ChatRole = "member"
	ChatAdminRole  ChatRole = "admin"
	ChatOwnerRole  ChatRole = "owner"
)

func (r ChatRole) Valid() bool {
	switch r {
	case ChatMemberRole, ChatAdminRole, ChatOwnerRole:
		return true
	}
	return false
}

func (r *ChatRole) UnmarshalText(text []byte) error {
	return unmarshalEnum(r, text, ChatRole.Valid, "chat role")
}

type ChatKind string

const (
	ChatGeneralKind ChatKind = "general"
	ChatGroupKind   ChatKind = "group"
	ChatDirectKind  ChatKind = "direct"
)

func (k ChatKind) Valid() bool {
	switch k {
	case ChatGeneralKind, ChatGroupKind, ChatDirectKind:
		return true
	}
	return false
}

func (k *ChatKind) UnmarshalText(text []byte) error {
	return unmarshalEnum(k, text, ChatKind.Valid, "chat kind")
}

type ChatSummary struct {
	ID            string  `json:"id"`
	SessionID     *string `json:"sessionId"`
	CampaignID    *string `json:"campaignId"`
	Kind          ChatKind `json:"kind"`
	Title         *string `json:"title"`
	PictureURL    *string `json:"pictureUrl"`
	LastMessageAt *time.Time `json:"lastMessageAt"`
	CreatedAt     time.Time `json:"createdAt"`
	OtherUserID  *string `json:"otherUserId,omitempty"` // for direct chats only
	LastMessage   *ChatLastMessage `json:"lastMessage,omitempty"`
}

type ChatMember struct {
	UserID    string   `json:"userId"`
	Role      ChatRole `json:"role"`
	JoinedAt  time.Time `json:"joinedAt"`
	LastReadID *string  `json:"lastReadId,omitempty"`
}

type ChatLastMessage struct {
	SenderID      string  `json:"senderId"`
	Body          *string `json:"body,omitempty"`
	HasAttachment bool    `json:"hasAttachment"`
}

type ChatPermissions struct {
	Role                 ChatRole `json:"role"`
	CanSendMessages      bool     `json:"canSendMessages"`
	CanSendFiles         bool     `json:"canSendFiles"`
	CanPinMessages       bool     `json:"canPinMessages"`
	CanChangeInfo        bool     `json:"canChangeInfo"`
	CanAddMembers        bool     `json:"canAddMembers"`
	CanRemoveMembers     bool     `json:"canRemoveMembers"`
	CanDeleteMessages    bool     `json:"canDeleteMessages"` // others' messages — you can always delete your own
	CanManageRoles       bool     `json:"canManageRoles"`
	CanManagePermissions bool     `json:"canManagePermissions"`
}