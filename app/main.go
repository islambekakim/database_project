package main

import (
	"database_project/api"
	"database_project/auth"
	"github.com/gin-gonic/gin"
	"log"
)

func initRouter() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.gohtml")

	auth.Init()

	OpenAdmin(router)
	OpenTeacher(router)

	log.Fatal(router.Run())
}

func OpenTeacher(router *gin.Engine) {
	teacher := router.Group("teacher/")
	teacher.GET("/:teacher_id", api.GetTeacherByID)
	teacher.GET("/:teacher_id/courses", api.GetCoursesByTeacherID)
	teacher.GET("/:teacher_id/:course_id", api.GetCourseByID)
	teacher.GET("/:teacher_id/:course_id/attendance", api.GetStudentsByCourseID)
}

func OpenAdmin(router *gin.Engine) {
	//
	admin := router.Group("admin/")
	admin.GET("", api.Admin)
	admin.GET("/logout", auth.Outing)
	admin.POST("", api.Admin)

	/** Student **/
	admin.GET("/student", api.GetStudents)
	admin.POST("/student/create", api.CreateStudent)
	admin.GET("/student/delete/:id", api.DeleteStudent)

	/** Teacher **/
	admin.GET("/teacher", api.GetTeachers)
	admin.POST("/teacher/create", api.CreateTeacher)
	admin.GET("/teacher/delete/:id", api.DeleteTeacher)

	/** Attendance **/
	admin.GET("/attendance", api.GetAttendance)
	admin.POST("/attendance/create", api.CreateAttendance)

	/** Class **/
	admin.GET("/classroom", api.GetClass)
	admin.POST("/classroom/create", api.CreateClass)
	admin.GET("/classroom/delete/:id", api.DeleteClass)

	/** Course **/
	admin.GET("/course", api.GetCourse)
	admin.POST("course/create", api.CreateCourse)
	admin.GET("/course/delete/:id", api.DeleteCourse)

}

func main() {

	initRouter()

}
