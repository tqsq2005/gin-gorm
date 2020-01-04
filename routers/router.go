package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	v1 "github.com/tqsq2005/gin-gorm/routers/v1"
)

func InitRouter() *gin.Engine  {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(viper.GetString("APP_MODE"))

	apiV1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		//新建标签
		apiV1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
