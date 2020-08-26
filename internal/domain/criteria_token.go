package domain

type PaginationToken string

func (t PaginationToken) GetPrimitive() string {
	return string(t)
}
