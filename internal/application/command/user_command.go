package command

import "context"

type Enable struct {
	Ctx context.Context
	ID  string
}

type Disable struct {
	Ctx context.Context
	ID  string
}

type Remove struct {
	Ctx context.Context
	ID  string
}
