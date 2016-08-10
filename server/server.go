package server

import "github.com/gin-gonic/gin"

type Server struct {
	*gin.Engine
}

func NewServer(dbCon *DatabaseConnection) Server {
	router := Server{gin.Default()}

	database := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Set("databaseconnection", dbCon)
		}
	}

	router.Use(database())
	router.GET("/", indexHandler)

	return router
}
