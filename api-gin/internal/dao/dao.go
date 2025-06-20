package dao

import (
	"admin-api/internal/dao/query"
	"admin-api/internal/services/db"
	"admin-api/pkg/lib/def"
	"errors"
	"reflect"
	"strings"
)

var q *query.Query

func GetQuery() *query.Query {
	if q == nil {
		inst := db.GetGDB("default")
		if inst == nil {
			panic("gorm db instance is nil")
		}
		q = query.Use(inst)
	}
	return q
}

func Row2KV(row any) (def.KV, error) {
	v := reflect.ValueOf(row)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("not struct data")
	}

	kv := def.KV{}
	t := v.Type()
	for i := range v.NumField() {
		f := t.Field(i)
		gormTag := f.Tag.Get("gorm")
		if col := getColName(gormTag); col != "" {
			kv[col] = v.Field(i).Interface()
		}
	}
	return kv, nil
}

func getColName(tag string) string {
	for s := range strings.SplitSeq(tag, ";") {
		if strings.HasPrefix(s, "column:") {
			return strings.TrimPrefix(s, "column:")
		}
	}
	return ""
}
