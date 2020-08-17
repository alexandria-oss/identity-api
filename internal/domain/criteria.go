package domain

import "encoding/json"

type Criteria struct {
	FilterBy FilterMap       `json:"filter_by"`
	Token    PaginationToken `json:"token"`
	Limit    Limit           `json:"limit"`
	OrderBy  Order           `json:"order_by"`
}

func (c *Criteria) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c Criteria) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}
