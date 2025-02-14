package core

type Helper interface {
	Copy(toValue any, fromValue any) error
}
