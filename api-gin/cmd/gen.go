package main

import (
	"admin-api/internal/services/conf"
	"admin-api/internal/services/db"

	"gorm.io/gen"
)

func main() {
	conf.InitConf("configs", "dev")
	db.InitDb(conf.Db)

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/dao/query",
		ModelPkgPath: "./internal/dao/model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	g.UseDB(db.GetGDB("default"))

	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}
