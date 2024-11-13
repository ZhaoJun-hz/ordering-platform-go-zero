// configuration.go
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./rpc/sys/gen/query", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	//var dsn = "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	var dsn = "root:QWERtyui1234@tcp(10.166.66.14:3306)/ordering-platform?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open(dsn))

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("sys_api"),
		g.GenerateModel("sys_casbin_rule"),
		g.GenerateModel("sys_dept"),
		g.GenerateModel("sys_menu"),
		g.GenerateModel("sys_menu_api_rule"),
		g.GenerateModel("sys_role"),
		g.GenerateModel("sys_role_menu"),
		g.GenerateModel("sys_user"),
	)

	// Execute the generator
	g.Execute()
}
