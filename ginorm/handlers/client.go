package handlers

import "github.com/gin-gonic/gin"

func FronendService(c *gin.Context) {
	c.HTML(200, "login.gohtml", nil)
}

func Register(c *gin.Context) {
	c.HTML(200, "register.gohtml", nil)
}
func Todo(c *gin.Context) {
	c.HTML(200, "todo.gohtml", nil)

}
