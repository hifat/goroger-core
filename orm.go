package core

import (
	"context"
)

type DBOrm struct {
	Error error
}

type IOrm interface {
	Model(value any) IOrm
	Table(name string, args ...any) IOrm
	Debug() IOrm

	WithContext(ctx context.Context) IOrm
	Begin() (IOrm, error)
	Commit() error
	Rollback() error

	Create(value any) error
	Save(value any) error
	Update(column string, value any) error
	Delete(query any, args ...any) error
	Where(query any, args ...any) IOrm
	Joins(query string, args ...any) IOrm
	InnerJoins(query string, args ...any) IOrm
	Raw(sql string, values ...any) IOrm

	Take(dest any, conds ...any) IOrm
	First(dest any, conds ...any) error
	Find(dest any, conds ...any) error
	Scan(dest any) IOrm
}
