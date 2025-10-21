package incentiveprograms

import "github.com/jayydoesdev/gokalshi"

type IncentiveProgramsQuery struct {
	Limit int
}

func (q IncentiveProgramsQuery) ToMap() map[string]string {
	return map[string]string{
		"limit": gokalshi.Sprintf(q.Limit),
	}
}
