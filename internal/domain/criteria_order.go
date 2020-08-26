package domain

type Order string

func (t Order) GetPrimitive() string {
	return string(t)
}
