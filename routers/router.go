package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tqsq2005/gin-gorm/middleware/jwt"
	"github.com/tqsq2005/gin-gorm/routers/api"
	v1 "github.com/tqsq2005/gin-gorm/routers/v1"
)

func InitRouter() *gin.Engine  {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(viper.GetString("APP_MODE"))

	//获取token
	r.GET("/auth", api.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		//新建标签
		apiV1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiV1.GET("/article/:id", v1.GetArticle)
		//新建文章
		apiV1.POST("/article", v1.AddArticle)
		//更新指定文章
		apiV1.PUT("/article/:id", v1.EditArticle)
		//删除指定文章
		apiV1.DELETE("/article/:id", v1.DeleteArticle)
	}

	return r
}
