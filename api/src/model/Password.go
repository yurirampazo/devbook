package model

// Model of body for updating password
type Password struct {
	NewPassword string `json:"new-password,omitempty"`
	OldPassword string `json:"old-password",omitempty`
}