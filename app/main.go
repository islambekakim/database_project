package main

import (
	"database_project/api"
	"github.com/gin-gonic/gin"
	"log"
)

func initRouter() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.gohtml")

	api.Init()

	OpenAdmin(router)

	log.Fatal(router.Run())
}

func OpenAdmin(router *gin.Engine) {
	//
	router.GET("/admin", api.Admin)

	/** Student **/
	router.GET("/admin/student", api.GetStudents)
	router.POST("/admin/student/create", api.CreateStudent)
	router.GET("/admin/student/delete/:id", api.DeleteStudent)

	/** Teacher **/
	router.GET("/admin/teacher", api.GetTeachers)
	router.POST("/admin/teacher/create", api.CreateTeacher)
	router.GET("/admin/teacher/delete/:id", api.DeleteTeacher)

	/** Attendance **/
	router.GET("/admin/attendance", api.GetAttendance)
	router.POST("/admin/attendance/create", api.CreateAttendance)

	/** Teacher **/
	router.GET("/admin/classroom", api.GetClass)
	router.POST("/admin/classroom/create", api.CreateClass)
	router.GET("/admin/classroom/delete/:id", api.DeleteClass)

	/** Teacher **/
	router.GET("/admin/course", api.GetCourse)
	router.POST("/admin/course/create", api.CreateCourse)
	router.GET("/admin/course/delete/:id", api.DeleteCourse)

}

func main() {

	initRouter()

}
