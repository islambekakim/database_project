package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Admin(context *gin.Context) {

	context.HTML(http.StatusOK, "Admin.gohtml", gin.H{})

}
