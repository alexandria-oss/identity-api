package value

type UserID struct {
	Value string
}

func NewUserIDFromString(id string) (*UserID, error) {
	return &UserID{Value: id}, nil
}
