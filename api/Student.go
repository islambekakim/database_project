package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"time"
)

type Student struct {
	//gorm.Model
	Student_id  int       `json:"student_id"`
	First_name  string    `json:"first_name"`
	Last_name   string    `json:"last_name"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	Birth_date  time.Time `json:"birth_Date"`
	Registredon time.Time `json:"registredon"`
	Lastlogin   time.Time `json:"lastlogin"`
}

func CreateStudent(context *gin.Context) {

	first := context.PostForm("FirstName")
	last := context.PostForm("LastName")
	login := context.PostForm("Login")
	password := context.PostForm("Password")
	date, _ := time.Parse("2006-01-02", context.PostForm("Birth"))
	reg := time.Now()

	var student = Student{
		Student_id:  -1,
		First_name:  first,
		Last_name:   last,
		Login:       login,
		Password:    password,
		Birth_date:  date,
		Registredon: reg,
	}

	DB.Create(&student)

	location := url.URL{Path: "/admin/student"}
	context.Redirect(http.StatusSeeOther, location.RequestURI())
}

func GetStudents(context *gin.Context) {

	students := []Student{}
	DB.Find(&students)
	context.HTML(http.StatusOK, "Student.gohtml", gin.H{
		"student": students,
	})

}

func DeleteStudent(context *gin.Context) {

	var student Student
	DB.Where("student_id = ?", context.Param("id")).Delete(&student)
	location := url.URL{Path: "/admin/student"}
	context.Redirect(http.StatusTemporaryRedirect, location.RequestURI())

}
