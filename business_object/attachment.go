package businessobject

import "time"

type Attachment struct {
	AttachmentId string    `json:"attachment_id"`
	ObjectId     string    `json:"object_id"`
	Kind         string    `json:"kind"`
	Url          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func GetAttachmentTable() string {
	return "Attachment"
}
