package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Attendance struct {
	Course_id  int       `json:"course_id"`
	Student_id int       `json:"student_id"`
	AttDate    time.Time `json:"attDate"`
	Stamp      string    `json:"stamp"`
}

func CreateAttendance(context *gin.Context) {

	course_id, _ := strconv.Atoi(context.PostForm("course"))
	student_id, _ := strconv.Atoi(context.PostForm("student"))
	date, _ := time.Parse("2006-01-02", context.PostForm("date"))
	stamp := context.PostForm("stamp")
	fmt.Println(date)
	fmt.Println(context.PostForm("date"))
	if stamp == "on" {
		stamp = "Y"
	} else {
		stamp = "N"
	}
	var attendance = Attendance{
		Course_id:  course_id,
		Student_id: student_id,
		AttDate:    date,
		Stamp:      stamp,
	}

	DB.Create(&attendance)

	location := url.URL{Path: "/admin/attendance"}
	context.Redirect(http.StatusSeeOther, location.RequestURI())
}

func GetAttendance(context *gin.Context) {
	attendances := []Attendance{}
	DB.Find(&attendances)
	context.HTML(http.StatusOK, "Attendance.gohtml", gin.H{
		"attendance": attendances,
	})
}
