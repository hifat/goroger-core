package orm

import (
	"context"

	core "github.com/hifat/goroger-core"
	"gorm.io/gorm"
)

type gormOrm struct {
	db *gorm.DB
	core.DBOrm
}

func NewGormOrm(db *gorm.DB) core.IOrm {
	return &gormOrm{db: db}
}

func (o *gormOrm) Model(value any) core.IOrm {
	return &gormOrm{
		db: o.db.Model(value),
	}
}

func (o *gormOrm) Table(name string, args ...any) core.IOrm {
	return &gormOrm{
		db: o.db.Table(name, args...),
	}
}

func (o *gormOrm) Debug() core.IOrm {
	return &gormOrm{
		db: o.db.Debug(),
	}
}

func (o *gormOrm) WithContext(ctx context.Context) core.IOrm {
	return &gormOrm{
		db: o.db.WithContext(ctx),
	}
}

func (o *gormOrm) Begin() (core.IOrm, error) {
	tx := o.db.Begin()

	return &gormOrm{db: tx}, tx.Error
}

func (o *gormOrm) Commit() error {
	return o.db.Commit().Error
}

func (o *gormOrm) Rollback() error {
	return o.db.Rollback().Error
}

func (o *gormOrm) Create(value any) error {
	return o.db.Create(value).Error
}

func (o *gormOrm) Save(value any) error {
	return o.db.Save(value).Error
}

func (o *gormOrm) Update(column string, value any) error {
	return o.db.Update(column, value).Error
}

func (o *gormOrm) Delete(value any, conds ...any) error {
	return o.db.Delete(value, conds...).Error
}

func (o *gormOrm) Where(query any, args ...any) core.IOrm {
	return &gormOrm{
		db: o.db.Where(query, args...),
	}
}

func (o *gormOrm) Joins(query string, args ...any) core.IOrm {
	return &gormOrm{
		db: o.db.Joins(query, args...),
	}
}

func (o *gormOrm) InnerJoins(query string, args ...any) core.IOrm {
	return &gormOrm{
		db: o.db.InnerJoins(query, args...),
	}
}

func (o *gormOrm) Raw(sql string, values ...any) core.IOrm {
	return &gormOrm{
		db: o.db.Raw(sql, values...),
	}
}

func (o *gormOrm) First(dest any, conds ...any) error {
	return o.db.First(dest, conds...).Error
}

func (o *gormOrm) Find(dest any, conds ...any) error {
	return o.db.Find(dest, conds...).Error
}

func (o *gormOrm) Take(dest any, conds ...any) core.IOrm {
	return &gormOrm{
		db: o.db.Take(dest, conds...),
	}
}

func (o *gormOrm) Scan(dest any) core.IOrm {
	return &gormOrm{
		db: o.db.Scan(dest),
	}
}
