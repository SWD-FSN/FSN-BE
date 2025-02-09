package businessobject

import "time"

type Action struct { // Reaction, follow, cmt, rep cmt
	ActionId   string    `json:"action_id"`
	ActionName string    `json:"action_name"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func GetActionTable() string {
	return "action"
}
