package routine

import (
	"assertion"
	"errors"
	"rut"
)

type Routine struct {
	Name        string
	Assertions  []*assertion.Assertion
	Description string
}

func (r *Routine) Run(db *rut.DB) error {
	for _, a := range r.Assertions {
		err := a.Do(db)
		if err != nil {
			return errors.New(fmt.Sprintf("Routine: %s failed with: %s", cd.Name, err.Error()))
		}
	}

	return nil
}
