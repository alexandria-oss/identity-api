package domain

type PaginationToken string

func NewPaginationToken(token string) PaginationToken {
	return PaginationToken(token)
}

func (t PaginationToken) GetPrimitive() string {
	return string(t)
}
