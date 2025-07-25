// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"admin-api/internal/dao/model"
)

func newGroupUser(db *gorm.DB, opts ...gen.DOOption) groupUser {
	_groupUser := groupUser{}

	_groupUser.groupUserDo.UseDB(db, opts...)
	_groupUser.groupUserDo.UseModel(&model.GroupUser{})

	tableName := _groupUser.groupUserDo.TableName()
	_groupUser.ALL = field.NewAsterisk(tableName)
	_groupUser.ID = field.NewInt32(tableName, "id")
	_groupUser.UID = field.NewInt32(tableName, "uid")
	_groupUser.Gid = field.NewInt32(tableName, "gid")

	_groupUser.fillFieldMap()

	return _groupUser
}

type groupUser struct {
	groupUserDo

	ALL field.Asterisk
	ID  field.Int32 // id
	UID field.Int32 // uid
	Gid field.Int32 // group.id

	fieldMap map[string]field.Expr
}

func (g groupUser) Table(newTableName string) *groupUser {
	g.groupUserDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g groupUser) As(alias string) *groupUser {
	g.groupUserDo.DO = *(g.groupUserDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *groupUser) updateTableName(table string) *groupUser {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.UID = field.NewInt32(table, "uid")
	g.Gid = field.NewInt32(table, "gid")

	g.fillFieldMap()

	return g
}

func (g *groupUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *groupUser) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 3)
	g.fieldMap["id"] = g.ID
	g.fieldMap["uid"] = g.UID
	g.fieldMap["gid"] = g.Gid
}

func (g groupUser) clone(db *gorm.DB) groupUser {
	g.groupUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g groupUser) replaceDB(db *gorm.DB) groupUser {
	g.groupUserDo.ReplaceDB(db)
	return g
}

type groupUserDo struct{ gen.DO }

type IGroupUserDo interface {
	gen.SubQuery
	Debug() IGroupUserDo
	WithContext(ctx context.Context) IGroupUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGroupUserDo
	WriteDB() IGroupUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGroupUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGroupUserDo
	Not(conds ...gen.Condition) IGroupUserDo
	Or(conds ...gen.Condition) IGroupUserDo
	Select(conds ...field.Expr) IGroupUserDo
	Where(conds ...gen.Condition) IGroupUserDo
	Order(conds ...field.Expr) IGroupUserDo
	Distinct(cols ...field.Expr) IGroupUserDo
	Omit(cols ...field.Expr) IGroupUserDo
	Join(table schema.Tabler, on ...field.Expr) IGroupUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGroupUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGroupUserDo
	Group(cols ...field.Expr) IGroupUserDo
	Having(conds ...gen.Condition) IGroupUserDo
	Limit(limit int) IGroupUserDo
	Offset(offset int) IGroupUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupUserDo
	Unscoped() IGroupUserDo
	Create(values ...*model.GroupUser) error
	CreateInBatches(values []*model.GroupUser, batchSize int) error
	Save(values ...*model.GroupUser) error
	First() (*model.GroupUser, error)
	Take() (*model.GroupUser, error)
	Last() (*model.GroupUser, error)
	Find() ([]*model.GroupUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GroupUser, err error)
	FindInBatches(result *[]*model.GroupUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GroupUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGroupUserDo
	Assign(attrs ...field.AssignExpr) IGroupUserDo
	Joins(fields ...field.RelationField) IGroupUserDo
	Preload(fields ...field.RelationField) IGroupUserDo
	FirstOrInit() (*model.GroupUser, error)
	FirstOrCreate() (*model.GroupUser, error)
	FindByPage(offset int, limit int) (result []*model.GroupUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGroupUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g groupUserDo) Debug() IGroupUserDo {
	return g.withDO(g.DO.Debug())
}

func (g groupUserDo) WithContext(ctx context.Context) IGroupUserDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g groupUserDo) ReadDB() IGroupUserDo {
	return g.Clauses(dbresolver.Read)
}

func (g groupUserDo) WriteDB() IGroupUserDo {
	return g.Clauses(dbresolver.Write)
}

func (g groupUserDo) Session(config *gorm.Session) IGroupUserDo {
	return g.withDO(g.DO.Session(config))
}

func (g groupUserDo) Clauses(conds ...clause.Expression) IGroupUserDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g groupUserDo) Returning(value interface{}, columns ...string) IGroupUserDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g groupUserDo) Not(conds ...gen.Condition) IGroupUserDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g groupUserDo) Or(conds ...gen.Condition) IGroupUserDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g groupUserDo) Select(conds ...field.Expr) IGroupUserDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g groupUserDo) Where(conds ...gen.Condition) IGroupUserDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g groupUserDo) Order(conds ...field.Expr) IGroupUserDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g groupUserDo) Distinct(cols ...field.Expr) IGroupUserDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g groupUserDo) Omit(cols ...field.Expr) IGroupUserDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g groupUserDo) Join(table schema.Tabler, on ...field.Expr) IGroupUserDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g groupUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGroupUserDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g groupUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IGroupUserDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g groupUserDo) Group(cols ...field.Expr) IGroupUserDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g groupUserDo) Having(conds ...gen.Condition) IGroupUserDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g groupUserDo) Limit(limit int) IGroupUserDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g groupUserDo) Offset(offset int) IGroupUserDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g groupUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupUserDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g groupUserDo) Unscoped() IGroupUserDo {
	return g.withDO(g.DO.Unscoped())
}

func (g groupUserDo) Create(values ...*model.GroupUser) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g groupUserDo) CreateInBatches(values []*model.GroupUser, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g groupUserDo) Save(values ...*model.GroupUser) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g groupUserDo) First() (*model.GroupUser, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupUser), nil
	}
}

func (g groupUserDo) Take() (*model.GroupUser, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupUser), nil
	}
}

func (g groupUserDo) Last() (*model.GroupUser, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupUser), nil
	}
}

func (g groupUserDo) Find() ([]*model.GroupUser, error) {
	result, err := g.DO.Find()
	return result.([]*model.GroupUser), err
}

func (g groupUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GroupUser, err error) {
	buf := make([]*model.GroupUser, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g groupUserDo) FindInBatches(result *[]*model.GroupUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g groupUserDo) Attrs(attrs ...field.AssignExpr) IGroupUserDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g groupUserDo) Assign(attrs ...field.AssignExpr) IGroupUserDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g groupUserDo) Joins(fields ...field.RelationField) IGroupUserDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g groupUserDo) Preload(fields ...field.RelationField) IGroupUserDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g groupUserDo) FirstOrInit() (*model.GroupUser, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupUser), nil
	}
}

func (g groupUserDo) FirstOrCreate() (*model.GroupUser, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupUser), nil
	}
}

func (g groupUserDo) FindByPage(offset int, limit int) (result []*model.GroupUser, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g groupUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g groupUserDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g groupUserDo) Delete(models ...*model.GroupUser) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *groupUserDo) withDO(do gen.Dao) *groupUserDo {
	g.DO = *do.(*gen.DO)
	return g
}
