package orm

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestModel struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	return db, mock
}

func TestNewGormOrm(t *testing.T) {
	t.Parallel()

	db, _ := setupTestDB(t)
	orm := NewGormOrm(db)
	assert.NotNil(t, orm)
}

func TestGormOrm_Create(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	model := &TestModel{Name: "test"}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `test_models`").
		WithArgs(model.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := orm.Create(model)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_First(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	model := &TestModel{}
	mock.ExpectQuery("SELECT \\* FROM `test_models`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test"))

	err := orm.First(model)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), model.ID)
	assert.Equal(t, "test", model.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_Transaction(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	mock.ExpectBegin()
	tx, err := orm.Begin()
	assert.NoError(t, err)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_WithContext(t *testing.T) {
	t.Parallel()

	db, _ := setupTestDB(t)
	orm := NewGormOrm(db)

	ctx := context.Background()
	ctxOrm := orm.WithContext(ctx)
	assert.NotNil(t, ctxOrm)
}

func TestGormOrm_Select(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	model := &TestModel{}
	mock.ExpectQuery("SELECT .+ FROM `test_models`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test"))

	err := orm.Select("id", "name").Find(model)
	assert.NoError(t, err)
	assert.Equal(t, "test", model.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_Where(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	model := &TestModel{}
	mock.ExpectQuery("SELECT \\* FROM `test_models` WHERE name = ?").
		WithArgs("test").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test"))

	err := orm.Where("name = ?", "test").Find(model)
	assert.NoError(t, err)
	assert.Equal(t, "test", model.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_Save(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	model := &TestModel{ID: 1, Name: "updated"}
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `test_models`").
		WithArgs("updated", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := orm.Save(model)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_Update(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `test_models`").
		WithArgs("new_name").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := orm.Model(&TestModel{}).Where(`1 = 1`).Update("name", "new_name")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_Delete(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `test_models`").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := orm.Delete(&TestModel{}, 1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_Joins(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	model := &TestModel{}
	mock.ExpectQuery("SELECT .* FROM `test_models` JOIN other_table").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test"))

	err := orm.Joins("JOIN other_table").First(model)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_InnerJoins(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	model := &TestModel{}
	mock.ExpectQuery("SELECT .* FROM `test_models` INNER JOIN other_table").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test"))

	err := orm.InnerJoins("INNER JOIN other_table").First(model)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_Raw(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	var result TestModel
	mock.ExpectQuery("SELECT \\* FROM test_models WHERE id = \\?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test"))

	err := orm.Raw("SELECT * FROM test_models WHERE id = ?", 1).First(&result)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_Find(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	var results []TestModel
	mock.ExpectQuery("SELECT \\* FROM `test_models`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test1").
			AddRow(2, "test2"))

	err := orm.Find(&results)
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_Take(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	model := &TestModel{}
	mock.ExpectQuery("SELECT \\* FROM `test_models`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test"))

	result := orm.Take(model)
	assert.NotNil(t, result)
}

func TestGormOrm_Scan(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	var result struct {
		Name  string
		Count int
	}

	mock.ExpectQuery("SELECT name, COUNT\\(\\*\\) as count FROM `test_models`").
		WillReturnRows(sqlmock.NewRows([]string{"name", "count"}).
			AddRow("test", 5))

	result2 := orm.Raw("SELECT name, COUNT(*) as count FROM `test_models`").Scan(&result)
	assert.NotNil(t, result2)
}

func TestGormOrm_Model(t *testing.T) {
	t.Parallel()

	db, _ := setupTestDB(t)
	orm := NewGormOrm(db)

	result := orm.Model(&TestModel{})
	assert.NotNil(t, result)
}

func TestGormOrm_Table(t *testing.T) {
	t.Parallel()

	db, _ := setupTestDB(t)
	orm := NewGormOrm(db)

	result := orm.Table("test_models")
	assert.NotNil(t, result)
}

func TestGormOrm_Debug(t *testing.T) {
	t.Parallel()

	db, _ := setupTestDB(t)
	orm := NewGormOrm(db)

	result := orm.Debug()
	assert.NotNil(t, result)
}

func TestGormOrm_Rollback(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	mock.ExpectBegin()
	tx, err := orm.Begin()
	assert.NoError(t, err)

	mock.ExpectRollback()
	err = tx.Rollback()
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormOrm_Not(t *testing.T) {
	t.Parallel()

	db, mock := setupTestDB(t)
	orm := NewGormOrm(db)

	model := &TestModel{}
	mock.ExpectQuery("SELECT \\* FROM `test_models` WHERE NOT name = ?").
		WithArgs("test").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(2, "test"))

	err := orm.Not("name = ?", "test").Find(model)
	assert.NoError(t, err)
	assert.Equal(t, "test", model.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}
