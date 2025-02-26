package businessobject

import "time"

type Conversation struct {
	ConversationId   string    `json:"conversation_id"`
	ConversationName []string  `json:"conversation_name"` // Tên đoạn chat. Nếu là đoạn chat giữa 2 user thì tên sẽ có 2 hiện tùy bên user dựa trên username. Nếu là group chat sẽ có 1 tên chung
	HostId           *string   `json:"host_id"`           // Trưởng nhóm nếu đây là group chat
	Members          []string  `json:"members"`           // Nếu đoạn chat giữa 2 người sẽ mặc định 2
	IsGroup          bool      `json:"is_group"`          // Có phải là group chat hay ko
	IsDelete         *bool     `json:"is_delete"`         // Nếu là group chat thì trưởng nhóm có quyền giải tán nhóm
	CreatedAt        time.Time `json:"created_at"`
}
