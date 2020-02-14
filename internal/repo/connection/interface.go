package connection

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

// Gormw is an interface which DB implements
type Gormw interface {
	Close() error
	DB() *sql.DB
	New() Gormw
	NewScope(value interface{}) *gorm.Scope
	CommonDB() gorm.SQLCommon
	Callback() *gorm.Callback
	SetLogger(l gorm.Logger)
	LogMode(enable bool) Gormw
	SingularTable(enable bool)
	Where(query interface{}, args ...interface{}) Gormw
	Or(query interface{}, args ...interface{}) Gormw
	Not(query interface{}, args ...interface{}) Gormw
	Limit(value int) Gormw
	Offset(value int) Gormw
	Order(value string, reorder ...bool) Gormw
	Select(query interface{}, args ...interface{}) Gormw
	Omit(columns ...string) Gormw
	Group(query string) Gormw
	Having(query string, values ...interface{}) Gormw
	Joins(query string, args ...interface{}) Gormw
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) Gormw
	Unscoped() Gormw
	Attrs(attrs ...interface{}) Gormw
	Assign(attrs ...interface{}) Gormw
	First(out interface{}, where ...interface{}) Gormw
	Last(out interface{}, where ...interface{}) Gormw
	Find(out interface{}, where ...interface{}) Gormw
	Scan(dest interface{}) Gormw
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) Gormw
	Count(value interface{}) Gormw
	Related(value interface{}, foreignKeys ...string) Gormw
	FirstOrInit(out interface{}, where ...interface{}) Gormw
	FirstOrCreate(out interface{}, where ...interface{}) Gormw
	Update(attrs ...interface{}) Gormw
	Updates(values interface{}, ignoreProtectedAttrs ...bool) Gormw
	UpdateColumn(attrs ...interface{}) Gormw
	UpdateColumns(values interface{}) Gormw
	Save(value interface{}) Gormw
	Create(value interface{}) Gormw
	Delete(value interface{}, where ...interface{}) Gormw
	Raw(sql string, values ...interface{}) Gormw
	Exec(sql string, values ...interface{}) Gormw
	Model(value interface{}) Gormw
	Table(name string) Gormw
	Debug() Gormw
	Begin() Gormw
	Commit() Gormw
	Rollback() Gormw
	NewRecord(value interface{}) bool
	RecordNotFound() bool
	CreateTable(values ...interface{}) Gormw
	DropTable(values ...interface{}) Gormw
	DropTableIfExists(values ...interface{}) Gormw
	HasTable(value interface{}) bool
	AutoMigrate(values ...interface{}) Gormw
	ModifyColumn(column string, typ string) Gormw
	DropColumn(column string) Gormw
	AddIndex(indexName string, column ...string) Gormw
	AddUniqueIndex(indexName string, column ...string) Gormw
	RemoveIndex(indexName string) Gormw
	AddForeignKey(field string, dest string, onDelete string, onUpdate string) Gormw
	Association(column string) *gorm.Association
	Preload(column string, conditions ...interface{}) Gormw
	Set(name string, value interface{}) Gormw
	InstantSet(name string, value interface{}) Gormw
	Get(name string) (value interface{}, ok bool)
	SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface)
	AddError(err error) error
	GetErrors() (errors []error)

	// extra
	Error() error
	//RowsAffected() int64
}

type gormDB = gorm.DB

type gormw struct {
	*gormDB
}

// Openw is a drop-in replacement for Open()
func Openw(dialect string, args ...interface{}) (db Gormw, err error) {
	gormdb, err := gorm.Open(dialect, args...)
	return wrap(gormdb), err
}

func (it *gormw) SetLogger(l gorm.Logger) {
	it.gormDB.SetLogger(l)
}

// wrap wraps gorm.DB in an interface
func wrap(db *gorm.DB) Gormw {
	return &gormw{gormDB: db}
}

func (it *gormw) New() Gormw {
	return wrap(it.gormDB.New())
}

func (it *gormw) LogMode(enable bool) Gormw {
	return wrap(it.gormDB.LogMode(enable))
}

func (it *gormw) Where(query interface{}, args ...interface{}) Gormw {
	return wrap(it.gormDB.Where(query, args...))
}

func (it *gormw) Or(query interface{}, args ...interface{}) Gormw {
	return wrap(it.gormDB.Or(query, args...))
}

func (it *gormw) Not(query interface{}, args ...interface{}) Gormw {
	return wrap(it.gormDB.Not(query, args...))
}

func (it *gormw) Limit(value int) Gormw {
	return wrap(it.gormDB.Limit(value))
}

func (it *gormw) Offset(value int) Gormw {
	return wrap(it.gormDB.Offset(value))
}

func (it *gormw) Order(value string, reorder ...bool) Gormw {
	return wrap(it.gormDB.Order(value, reorder...))
}

func (it *gormw) Select(query interface{}, args ...interface{}) Gormw {
	return wrap(it.gormDB.Select(query, args...))
}

func (it *gormw) Omit(columns ...string) Gormw {
	return wrap(it.gormDB.Omit(columns...))
}

func (it *gormw) Group(query string) Gormw {
	return wrap(it.gormDB.Group(query))
}

func (it *gormw) Having(query string, values ...interface{}) Gormw {
	return wrap(it.gormDB.Having(query, values...))
}

func (it *gormw) Joins(query string, args ...interface{}) Gormw {
	return wrap(it.gormDB.Joins(query, args...))
}

func (it *gormw) Scopes(funcs ...func(*gorm.DB) *gorm.DB) Gormw {
	return wrap(it.gormDB.Scopes(funcs...))
}

func (it *gormw) Unscoped() Gormw {
	return wrap(it.gormDB.Unscoped())
}

func (it *gormw) Attrs(attrs ...interface{}) Gormw {
	return wrap(it.gormDB.Attrs(attrs...))
}

func (it *gormw) Assign(attrs ...interface{}) Gormw {
	return wrap(it.gormDB.Assign(attrs...))
}

func (it *gormw) First(out interface{}, where ...interface{}) Gormw {
	return wrap(it.gormDB.First(out, where...))
}

func (it *gormw) Last(out interface{}, where ...interface{}) Gormw {
	return wrap(it.gormDB.Last(out, where...))
}

func (it *gormw) Find(out interface{}, where ...interface{}) Gormw {
	return wrap(it.gormDB.Find(out, where...))
}

func (it *gormw) Scan(dest interface{}) Gormw {
	return wrap(it.gormDB.Scan(dest))
}

func (it *gormw) Pluck(column string, value interface{}) Gormw {
	return wrap(it.gormDB.Pluck(column, value))
}

func (it *gormw) Count(value interface{}) Gormw {
	return wrap(it.gormDB.Count(value))
}

func (it *gormw) Related(value interface{}, foreignKeys ...string) Gormw {
	return wrap(it.gormDB.Related(value, foreignKeys...))
}

func (it *gormw) FirstOrInit(out interface{}, where ...interface{}) Gormw {
	return wrap(it.gormDB.FirstOrInit(out, where...))
}

func (it *gormw) FirstOrCreate(out interface{}, where ...interface{}) Gormw {
	return wrap(it.gormDB.FirstOrCreate(out, where...))
}

func (it *gormw) Update(attrs ...interface{}) Gormw {
	return wrap(it.gormDB.Update(attrs...))
}

func (it *gormw) Updates(values interface{}, ignoreProtectedAttrs ...bool) Gormw {
	return wrap(it.gormDB.Updates(values, ignoreProtectedAttrs...))
}

func (it *gormw) UpdateColumn(attrs ...interface{}) Gormw {
	return wrap(it.gormDB.UpdateColumn(attrs...))
}

func (it *gormw) UpdateColumns(values interface{}) Gormw {
	return wrap(it.gormDB.UpdateColumns(values))
}

func (it *gormw) Save(value interface{}) Gormw {
	return wrap(it.gormDB.Save(value))
}

func (it *gormw) Create(value interface{}) Gormw {
	return wrap(it.gormDB.Create(value))
}

func (it *gormw) Delete(value interface{}, where ...interface{}) Gormw {
	return wrap(it.gormDB.Delete(value, where...))
}

func (it *gormw) Raw(sql string, values ...interface{}) Gormw {
	return wrap(it.gormDB.Raw(sql, values...))
}

func (it *gormw) Exec(sql string, values ...interface{}) Gormw {
	return wrap(it.gormDB.Exec(sql, values...))
}

func (it *gormw) Model(value interface{}) Gormw {
	return wrap(it.gormDB.Model(value))
}

func (it *gormw) Table(name string) Gormw {
	return wrap(it.gormDB.Table(name))
}

func (it *gormw) Debug() Gormw {
	return wrap(it.gormDB.Debug())
}

func (it *gormw) Begin() Gormw {
	return wrap(it.gormDB.Begin())
}

func (it *gormw) Commit() Gormw {
	return wrap(it.gormDB.Commit())
}

func (it *gormw) Rollback() Gormw {
	return wrap(it.gormDB.Rollback())
}

func (it *gormw) CreateTable(values ...interface{}) Gormw {
	return wrap(it.gormDB.CreateTable(values...))
}

func (it *gormw) DropTable(values ...interface{}) Gormw {
	return wrap(it.gormDB.DropTable(values...))
}

func (it *gormw) DropTableIfExists(values ...interface{}) Gormw {
	return wrap(it.gormDB.DropTableIfExists(values...))
}

func (it *gormw) AutoMigrate(values ...interface{}) Gormw {
	return wrap(it.gormDB.AutoMigrate(values...))
}

func (it *gormw) ModifyColumn(column string, typ string) Gormw {
	return wrap(it.gormDB.ModifyColumn(column, typ))
}

func (it *gormw) DropColumn(column string) Gormw {
	return wrap(it.gormDB.DropColumn(column))
}

func (it *gormw) AddIndex(indexName string, columns ...string) Gormw {
	return wrap(it.gormDB.AddIndex(indexName, columns...))
}

func (it *gormw) AddUniqueIndex(indexName string, columns ...string) Gormw {
	return wrap(it.gormDB.AddUniqueIndex(indexName, columns...))
}

func (it *gormw) RemoveIndex(indexName string) Gormw {
	return wrap(it.gormDB.RemoveIndex(indexName))
}

func (it *gormw) Preload(column string, conditions ...interface{}) Gormw {
	return wrap(it.gormDB.Preload(column, conditions...))
}

func (it *gormw) Set(name string, value interface{}) Gormw {
	return wrap(it.gormDB.Set(name, value))
}

func (it *gormw) InstantSet(name string, value interface{}) Gormw {
	return wrap(it.gormDB.InstantSet(name, value))
}

func (it *gormw) AddForeignKey(field string, dest string, onDelete string, onUpdate string) Gormw {
	return wrap(it.gormDB.AddForeignKey(field, dest, onDelete, onUpdate))
}

func (it *gormw) Error() error {
	return it.gormDB.Error
}
