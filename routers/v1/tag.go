package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tqsq2005/gin-gorm/models"
	"github.com/tqsq2005/gin-gorm/pkg/e"
	"github.com/tqsq2005/gin-gorm/utils"
	"github.com/unknwon/com"
	"net/http"
)

//获取多个文章的标签
func GetTags(c *gin.Context)  {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(utils.GetPage(c), viper.GetInt("PAGE_SIZE"), maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

//新增文章的标签
func AddTag(c *gin.Context)  {

}

//修改文章的标签
func EditTag(c *gin.Context)  {

}

//删除文章的标签
func DeleteTag(c *gin.Context)  {

}