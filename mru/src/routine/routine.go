package routine

import (
	"assertion"
)

type Routine struct {
	Name        string
	Assertions  []*assertion.Assertion
	Description string
}