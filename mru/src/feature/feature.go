package feature

import (
	"errors"
	"mcase"
	"sort"
)

type Feature struct {
	Name   string
	CCount int
	Cases  map[string]*mcase.Case
}

func (f *Feature) Add(c *mcase.Case) error {
	_, ok := f.Cases[c.Name]
	if ok {
		return errors.New("Case: " + c.Name + "Alread exist for Group: " + c.Group + " Feature: " + c.Feature)
	}

	f.CCount++
	f.Cases[c.Name] = c

	return nil
}

func (f *Feature) Del(c *mcase.Case) error {
	_, ok := f.Cases[c.Name]
	if !ok {
		return errors.New("Cannot find  Case: " + c.Name + "for delete under Group: " + c.Group + " Feature: " + c.Feature)
	}

	delete(f.Cases, c.Name)
	f.CCount--
	return nil
}

func (f *Feature) Get(c *mcase.Case) (*mcase.Case, error) {
	old, ok := f.Cases[c.Name]
	if !ok {
		return nil, errors.New("Cannot find  Case: " + c.Name + "for Get under Group: " + c.Group + " Feature: " + c.Feature)
	}

	return old, nil
}

func (f *Feature) Dump() []*mcase.Case {
	cs := make([]*mcase.Case, 0, len(f.Cases))

	for _, c := range f.Cases {
		cs = append(cs, c)
	}

	//sort.Slice(cs, func(i, j int) bool { return cs[i].Name < cs[j].Name })
	sort.Stable(CaseSlice(cs))

	return cs
}

type CaseSlice []*mcase.Case

func (s CaseSlice) Len() int           { return len(s) }
func (s CaseSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s CaseSlice) Less(i, j int) bool { return s[i].Name < s[j].Name }
