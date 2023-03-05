package templates

const MainTemplate = `package main

import (
	"github.com/gin-gonic/gin"
	"{{.ModModuleName}}/handlers"
)

func main(){
	r := gin.Default()
	{{range .Routes}}
	r.{{.Method}}("{{.Path}}",handlers.{{.HandlerFunc}})
	{{end}}

	err := r.Run(":8080")
	if err != nil{
		panic(err.Error())
	}
}
`
