package models

type Url struct {
	OldUrl string `json:"old_url"`
	NewUrl string `json:"new_url" validate:"regexp=^[a-zA-Z0-9-_]*$"`
}
