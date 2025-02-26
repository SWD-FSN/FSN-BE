package dto

type GetObjectsFromEnterSearchBarResponse struct {
	Users *[]UserSearchDoneResponse `json:"users"`
	Posts *[]PostResponse           `json:"posts"`
}
