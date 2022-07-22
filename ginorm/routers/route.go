package route

import (
	"ginorm/handlers"

	"github.com/gin-gonic/gin"
)

func Route() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.gohtml")

	frontend := r.Group("/")
	{
		frontend.GET("/", handlers.FronendService)
		frontend.GET("/register", handlers.Register)

	}

	r.Run(":8080")

}
