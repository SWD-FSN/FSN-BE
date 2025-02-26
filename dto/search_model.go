package dto

import (
	business_object "social_network/business_object"
)

type GetObjectsFromEnterSearchBarResponse struct {
	Users *[]UserSearchDoneResponse `json:"users"`
	Posts *[]business_object.Post   `json:"posts"`
}
