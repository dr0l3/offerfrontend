package handler

import "github.com/gin-gonic/gin"

func indexHandler(c *gin.Context) {
	c.HTML(200, "templates/index.tmpl", gin.H{"name": "Gopher"})
}
