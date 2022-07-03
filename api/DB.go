package api

import (
	"database_project/pkg/oracle"
)

var DB, err = oracle.NewConnection()

func Init() {
	DB.AutoMigrate(&Course{})
	DB.AutoMigrate(&Teacher{})
	DB.AutoMigrate(&Student{})
	DB.AutoMigrate(&Attendance{})
	DB.AutoMigrate(&Classroom{})
}
