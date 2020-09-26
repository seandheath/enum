package module

import "github.com/google/uuid"

type Target struct {
	name  string
	ttype string
	id    uuid.UUID
}

func newTarget(name string, ttype string) *Target {
	t := Target{
		name:  name,
		ttype: ttype,
		id:    uuid.New(),
	}
	return &t
}
