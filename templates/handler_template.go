package templates

const HandlerTemplate = `package handlers
import (
	"github.com/gin-gonic/gin"
)

func {{.FuncName}}(c *gin.Context) {
	c.JSON(200,gin.H{
		"message":"{{.FuncName}} handler not implemented",
	})
}
`
