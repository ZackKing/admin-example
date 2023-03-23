package models

import (
	"admin-api/internal/services/db"
	"log"

	"gorm.io/gorm"
)

type baseMdl struct {
	dbName string
}

func (m *baseMdl) getDb() (gdb *gorm.DB) {
	gdb = db.GetGormDb(m.dbName)

	if gdb == nil {
		log.Fatal("model error db conn get: ", m.dbName)
	}
	return gdb
}
