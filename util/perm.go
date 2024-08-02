package util

import (
	"fmt"
)

type Perm struct {
	Owner  uint8
	Group  uint8
	Others uint8
}

const (
	NoRights  uint8 = 0
	Executing uint8 = 1 // read file meta-data
	Writing   uint8 = 2
	Reading   uint8 = 4
)

func NewPerm(owner, group, others uint8) Perm {
	return Perm{Owner: owner, Group: group, Others: others}
}

func (m Perm) String() string {
	if m.Owner > 7 {
		panic("owner mod cannot be greater than 7")
	}
	if m.Group > 7 {
		panic("group mod cannot be greater than 7")
	}
	if m.Others > 7 {
		panic("others mod cannot be greater than 7")
	}
	return fmt.Sprintf("%d%d%d", m.Owner, m.Group, m.Others)
}
