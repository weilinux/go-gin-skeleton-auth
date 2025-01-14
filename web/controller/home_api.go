package controller

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/app"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "welcome"})
}

func SwagDoc(c *gin.Context) {
	fInfo, _ := os.Stat("static/swagger.json")

	data := map[string]string{
		"EnvName":    app.EnvName,
		"AppName":    app.Name,
		"JsonFile":   "/static/swagger.json",
		"SwgUIPath":  "/static/swagger-ui",
		"AssetPath":  "/static",
		"UpdateTime": fInfo.ModTime().Format(app.DateFormat),
	}

	c.HTML(200, "swagger.tpl", data)
}

// @Tags InternalApi
// @Summary 检测API
// @Description get app health
// @Success 201 {string} json data
// @Failure 403 body is empty
// @Router /health [get]
func AppHealth(c *gin.Context) {
	data := map[string]interface{}{
		"status": "UP",
		"info":   app.GitInfo,
	}

	c.JSON(200, data)
}

func AppStatus(c *gin.Context) {
	data := map[string]interface{}{
		"status": "UP",
		"info":   app.GitInfo,
	}

	c.JSON(200, data)
}
