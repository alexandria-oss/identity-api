package entity

import "time"

type User struct {
	Sub        string     `json:"sub"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Name       string     `json:"name"`
	MiddleName *string    `json:"middle_name,omitempty"`
	FamilyName *string    `json:"family_name,omitempty"`
	Locale     string     `json:"locale"`
	Picture    *string    `json:"picture,omitempty"`
	Status     string     `json:"status"`
	CreateTime *time.Time `json:"create_time,omitempty"`
	UpdateTime *time.Time `json:"update_time,omitempty"`
	Enabled    bool       `json:"enabled"`
}
