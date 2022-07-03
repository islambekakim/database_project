package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
)

type Classroom struct {
	Classroom_id int `json:"classroom_id"`
	Capacity     int `json:"capacity"`
}

func CreateClass(context *gin.Context) {

	capasity, _ := strconv.Atoi(context.PostForm("capacity"))

	var classroom = Classroom{
		Classroom_id: -1,
		Capacity:     capasity,
	}

	DB.Create(&classroom)

	location := url.URL{Path: "/admin/classroom"}
	context.Redirect(http.StatusSeeOther, location.RequestURI())
}

func GetClass(context *gin.Context) {
	classrooms := []Classroom{}
	DB.Find(&classrooms)
	context.HTML(http.StatusOK, "Classroom.gohtml", gin.H{
		"classrooms": classrooms,
	})
}

func DeleteClass(context *gin.Context) {

	var class Classroom
	DB.Where("classroom_id = ?", context.Param("id")).Delete(&class)
	location := url.URL{Path: "/admin/classroom"}
	context.Redirect(http.StatusTemporaryRedirect, location.RequestURI())

}
