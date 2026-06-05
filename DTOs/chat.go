package dtos

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