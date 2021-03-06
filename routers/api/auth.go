package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tqsq2005/gin-gorm/models"
	"github.com/tqsq2005/gin-gorm/pkg/e"
	"github.com/tqsq2005/gin-gorm/pkg/utils"
	"net/http"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

// @Summary 获取Token认证
// @Tags 认证
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{
		Username:username,
		Password:password,
	}
	ok, _ := valid.Valid(a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := utils.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}
