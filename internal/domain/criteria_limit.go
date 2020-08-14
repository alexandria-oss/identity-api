package domain

import "strconv"

type Limit int

func NewLimit(l string) Limit {
	lInt, err := strconv.ParseInt(l, 10, 32)
	if err != nil {
		return 10
	} else if lInt <= 0 {
		return 10
	}

	return Limit(int(lInt))
}

func (l Limit) GetPrimitive() int {
	return int(l)
}
