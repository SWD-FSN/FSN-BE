package businessobject

import "time"

type Message struct {
	MessageId     string    `json:"message_id"`
	AuthorId      string    `json:"author_id"`
	CoversationId string    `json:"conversation_id"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
}

func GetMessageTable() string {
	return "messages"
}
