package core

import (
	core "github.com/hifat/goroger-core"
	"github.com/jinzhu/copier"
)

type helper struct{}

func New() core.Helper {
	return &helper{}
}

func (h *helper) Copy(toValue any, fromValue any) error {
	return copier.Copy(toValue, fromValue)
}
