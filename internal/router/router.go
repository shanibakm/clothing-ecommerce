package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	return r
}
