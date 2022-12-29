package main

import (
	"chatapp/pkg/utils"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"os"
)

func main() {
	action := os.Args[1]
	//steps := os.Args[2]
	switch action {
	case "up":
		up()
		break
	case "down":
		down()
		break
	}
}

func dbConnect() *migrate.Migrate {
	dsn := "root@tcp(localhost:3306)/chatapp"
	db, err := sql.Open("mysql", dsn)
	utils.CheckError(err)
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance("file:///Users/briantam/Project/code/golang/chatapp/pkg/database/migration", "chatapp", driver)
	utils.CheckError(err)
	return m
}

func up() {
	m := dbConnect()
	err := m.Up()
	if err != nil {
		utils.CheckError(err)
	}
}

func down() {
	m := dbConnect()
	err := m.Down()
	if err != nil {
		utils.CheckError(err)
	}
}
