package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
)

type Course struct {
	Course_id    int    `json:"course_id"`
	Classroom_id int    `json:"classroom_id"`
	Course_Code  string `json:"coursecode"`
	Teacher_id   int    `json:"teacher_id"`
}

func CreateCourse(context *gin.Context) {

	classroom, _ := strconv.Atoi(context.PostForm("classroom"))
	code := context.PostForm("code")
	var teacher, _ = strconv.Atoi(context.PostForm("teacher"))

	var course = Course{
		Course_id:    -1,
		Classroom_id: classroom,
		Course_Code:  code,
		Teacher_id:   teacher,
	}
	DB.Create(&course)
	location := url.URL{Path: "/admin/course"}
	context.Redirect(http.StatusSeeOther, location.RequestURI())

}

func GetCourse(context *gin.Context) {
	courses := []Course{}
	DB.Find(&courses)
	context.HTML(http.StatusOK, "Course.gohtml", gin.H{
		"courses": courses,
	})
}

func DeleteCourse(context *gin.Context) {
	var course Course
	DB.Where("course_id = ?", context.Param("id")).Delete(&course)
	location := url.URL{Path: "/admin/course"}
	context.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
}
