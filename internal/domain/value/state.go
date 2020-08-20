package value

import "strconv"

type State struct {
	Value bool
}

func NewStateFromString(state string) (*State, error) {
	s, err := strconv.ParseBool(state)
	if err != nil {
		return nil, err
	}

	return &State{Value: s}, nil
}

func (u *State) Activate() {
	u.Value = true
}

func (u *State) Deactivate() {
	u.Value = false
}

func (u State) ToString() string {
	return strconv.FormatBool(u.Value)
}
