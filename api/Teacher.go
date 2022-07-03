package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"time"
)

type Teacher struct {
	Teacher_id  int       `json:"teacher_id"`
	First_name  string    `json:"first_name"`
	Last_name   string    `json:"last_name"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	Registredon time.Time `json:"registredon"`
	Lastlogin   time.Time `json:"lastlogin"`
}

func CreateTeacher(context *gin.Context) {

	first := context.PostForm("FirstName")
	last := context.PostForm("LastName")
	login := context.PostForm("Login")
	password := context.PostForm("Password")
	reg := time.Now()

	var teacher = Teacher{
		Teacher_id:  -1,
		First_name:  first,
		Last_name:   last,
		Login:       login,
		Password:    password,
		Registredon: reg,
	}

	DB.Create(&teacher)

	location := url.URL{Path: "/admin/teacher"}
	context.Redirect(http.StatusSeeOther, location.RequestURI())
}

func GetTeachers(context *gin.Context) {
	teachers := []Teacher{}
	DB.Find(&teachers)
	context.HTML(http.StatusOK, "Teacher.gohtml", gin.H{
		"teachers": teachers,
	})
}

func DeleteTeacher(context *gin.Context) {

	var teacher Teacher
	DB.Where("teacher_id = ?", context.Param("id")).Delete(&teacher)
	location := url.URL{Path: "/admin/teacher"}
	context.Redirect(http.StatusTemporaryRedirect, location.RequestURI())

}
