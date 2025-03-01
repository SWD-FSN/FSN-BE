package businessobject

type Message struct {
	MessageId     string `json:"message_id"`
	AuthorId      string `json:"author_id"`
	CoversationId string `json:"conversation_id"`
	Content       string `json:"content"`
	CreatedAt     string `json:"created_at"`
}

func GetMessageTable() string {
	return "message"
}
