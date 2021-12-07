package models

type InternSuggestion struct {
	ID         string `json:"id"`
	UserID     string `json:"userID"`
	Suggestion string `json:"suggestion"`
}

type ExternSuggestion struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Suggestion string `json:"suggestion"`
}
