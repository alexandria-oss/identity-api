package domain

import "time"

type User struct {
	Sub        string     `json:"sub"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Name       string     `json:"name"`
	MiddleName *string    `json:"middle_name"`
	FamilyName *string    `json:"family_name"`
	Locale     string     `json:"locale"`
	Picture    *string    `json:"picture"`
	Status     string     `json:"status"`
	CreateTime *time.Time `json:"create_time"`
	UpdateTime *time.Time `json:"update_time"`
	Enabled    bool       `json:"enabled"`
}
