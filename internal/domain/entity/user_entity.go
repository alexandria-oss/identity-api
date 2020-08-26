package entity

import (
	"github.com/alexandria-oss/identity-api/internal/domain/value"
	"time"
)

type User struct {
	ID                *value.UserID   `json:"sub"`
	Username          *value.Username `json:"username"`
	PreferredUsername *string         `json:"preferred_username,omitempty"`
	Email             *value.Email    `json:"email"`
	Name              string          `json:"name,omitempty"`
	MiddleName        *string         `json:"middle_name,omitempty"`
	FamilyName        *string         `json:"family_name,omitempty"`
	Locale            string          `json:"locale,omitempty"`
	Picture           *string         `json:"picture,omitempty"`
	Status            string          `json:"status"`
	CreateTime        *time.Time      `json:"create_time,omitempty"`
	UpdateTime        *time.Time      `json:"update_time,omitempty"`
	Enabled           *value.State    `json:"enabled"`
}

func (u User) IsValid() error {
	if err := u.Username.IsValid(); err != nil {
		return err
	} else if err := u.Email.IsValid(); err != nil {
		return err
	}

	return nil
}

func (u User) ToPrimitive() *UserPrimitive {
	return &UserPrimitive{
		ID:                u.ID.Value,
		Username:          u.Username.Value,
		PreferredUsername: u.PreferredUsername,
		Email:             u.Email.Value,
		Name:              u.Name,
		MiddleName:        u.MiddleName,
		FamilyName:        u.FamilyName,
		Locale:            u.Locale,
		Picture:           u.Picture,
		Status:            u.Status,
		CreateTime:        u.CreateTime,
		UpdateTime:        u.UpdateTime,
		Enabled:           u.Enabled.Value,
	}
}

type UserPrimitive struct {
	ID                string     `json:"sub"`
	Username          string     `json:"username"`
	PreferredUsername *string    `json:"preferred_username,omitempty"`
	Email             string     `json:"email"`
	Name              string     `json:"name,omitempty"`
	MiddleName        *string    `json:"middle_name,omitempty"`
	FamilyName        *string    `json:"family_name,omitempty"`
	Locale            string     `json:"locale,omitempty"`
	Picture           *string    `json:"picture,omitempty"`
	Status            string     `json:"status"`
	CreateTime        *time.Time `json:"create_time,omitempty"`
	UpdateTime        *time.Time `json:"update_time,omitempty"`
	Enabled           bool       `json:"enabled"`
}

func (u UserPrimitive) ToEntity() (*User, error) {
	id, err := value.NewUserIDFromString(u.ID)
	if err != nil {
		return nil, err
	}

	username, err := value.NewUsernameFromString(u.Username)
	if err != nil {
		return nil, err
	}

	email, err := value.NewEmailFromString(u.Email)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:                id,
		Username:          username,
		PreferredUsername: u.PreferredUsername,
		Email:             email,
		Name:              u.Name,
		MiddleName:        u.MiddleName,
		FamilyName:        u.FamilyName,
		Locale:            u.Locale,
		Picture:           u.Picture,
		Status:            u.Status,
		CreateTime:        u.CreateTime,
		UpdateTime:        u.UpdateTime,
		Enabled:           &value.State{Value: u.Enabled},
	}, nil
}
