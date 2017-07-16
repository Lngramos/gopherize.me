// Code generated by immutableGen. DO NOT EDIT.

package other

//immutableVet:skipFile

import (
	"myitcv.io/immutable"
)

//
// MySlice is an immutable type and has the following template:
//
// 	[]string
//
type MySlice struct {
	theSlice []string
	mutable  bool
	__tmpl   _Imm_MySlice
}

var _ immutable.Immutable = new(MySlice)
var _ = new(MySlice).__tmpl

func NewMySlice(s ...string) *MySlice {
	c := make([]string, len(s))
	copy(c, s)

	return &MySlice{
		theSlice: c,
	}
}

func NewMySliceLen(l int) *MySlice {
	c := make([]string, l)

	return &MySlice{
		theSlice: c,
	}
}

func (m *MySlice) Mutable() bool {
	return m.mutable
}

func (m *MySlice) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theSlice)
}

func (m *MySlice) Get(i int) string {
	return m.theSlice[i]
}

func (m *MySlice) AsMutable() *MySlice {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *MySlice) dup() *MySlice {
	resSlice := make([]string, len(m.theSlice))

	for i := range m.theSlice {
		resSlice[i] = m.theSlice[i]
	}

	res := &MySlice{
		theSlice: resSlice,
	}

	return res
}

func (m *MySlice) AsImmutable(v *MySlice) *MySlice {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *MySlice) Range() []string {
	if m == nil {
		return nil
	}

	return m.theSlice
}

func (m *MySlice) WithMutable(f func(mi *MySlice)) *MySlice {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *MySlice) WithImmutable(f func(mi *MySlice)) *MySlice {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *MySlice) Set(i int, v string) *MySlice {
	if m.mutable {
		m.theSlice[i] = v
		return m
	}

	res := m.dup()
	res.theSlice[i] = v

	return res
}

func (m *MySlice) Append(v ...string) *MySlice {
	if m.mutable {
		m.theSlice = append(m.theSlice, v...)
		return m
	}

	res := m.dup()
	res.theSlice = append(res.theSlice, v...)

	return res
}
func (s *MySlice) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	return true
}
