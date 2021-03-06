package def

import (
	"sort"
)

type ID interface {
	ID() string
}

type Member interface {
	ID
	Delete(component Component) error
}

type MemberSlice []Member

func (p MemberSlice) Len() int           { return len(p) }
func (p MemberSlice) Less(i, j int) bool { return p[i].ID() < p[j].ID() }
func (p MemberSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p MemberSlice) Sort()              { sort.Sort(p) }
func (l MemberSlice) Append(t Member) MemberSlice {
	ll := append(l, t)
	ll.Sort()
	return ll
}
func (l MemberSlice) Remove(index int) MemberSlice {
	var ll MemberSlice
	switch index {
	case 0:
		ll = l[1:]
	case len(l) - 1:
		ll = l[0 : len(l)-2]
	default:
		ll = append(l[:index-1], l[index+1:]...)
	}
	ll.Sort()
	return ll
}

func (l MemberSlice) Search(t Member) int {
	return l.SearchID(t.ID())
}

func (l MemberSlice) SearchID(t string) int {
	index := sort.Search(l.Len(), func(i int) bool { return l[i].ID() >= t })
	if index < l.Len() && l[index].ID() == t {
		return index
	} else {
		return -1
	}
}

type Component interface {
	ID
	Group(category string) bool
	Remove(member Member) error
	Join(member Member) error
}

type ComponentSlice []Component

func (p ComponentSlice) Len() int           { return len(p) }
func (p ComponentSlice) Less(i, j int) bool { return p[i].ID() < p[j].ID() }
func (p ComponentSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ComponentSlice) Sort()              { sort.Sort(p) }
func (l ComponentSlice) Append(t Component) ComponentSlice {
	ll := append(l, t)
	ll.Sort()
	return ll
}
func (l ComponentSlice) Remove(index int) ComponentSlice {
	var ll ComponentSlice
	switch index {
	case 0:
		ll = l[1:]
	case len(l) - 1:
		ll = l[0 : len(l)-2]
	default:
		ll = append(l[:index-1], l[index+1:]...)
	}
	ll.Sort()
	return ll
}

func (l ComponentSlice) Search(t Component) int {
	return l.SearchID(t.ID())
}

func (l ComponentSlice) SearchID(t string) int {
	index := sort.Search(l.Len(), func(i int) bool { return l[i].ID() >= t })
	if index < l.Len() && l[index].ID() == t {
		return index
	} else {
		return -1
	}
}
