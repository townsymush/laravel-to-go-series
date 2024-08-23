package storage

import "fmt"

func NewNotFoundError(name string) NotFoundErr {
	return NotFoundErr{
		name: name,
	}
}

type NotFoundErr struct {
	name string
}

func (e NotFoundErr) Error() string {
	return fmt.Sprintf("%s not found", e.name)
}
