package domain

type Criteria struct {
	FilterBy FilterMap
	Token    PaginationToken
	Limit    Limit
	OrderBy  Order
}
