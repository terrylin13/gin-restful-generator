package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/terrylin13/gin-restful-generator/templates"
	"github.com/terrylin13/gin-restful-generator/utils"
)

type Resource struct {
	Name       string
	PluralName string
}

type Operation struct {
	Method   string
	Path     string
	FuncName string
}

type Route struct {
	Method      string
	Path        string
	HandlerFunc string
}

var (
	name   string
	dir    string
	module string
)

var operations = []Operation{
	{"GET", "/", "List" + "{{.Name}}"},
	{"GET", "/:" + "{{.Name}}" + "ID", "Get" + "{{.Name}}"},
	{"POST", "/", "Create" + "{{.Name}}"},
	{"PUT", "/:" + "{{.Name}}" + "ID", "Update" + "{{.Name}}"},
	{"DELETE", "/:" + "{{.Name}}" + "ID", "Delete" + "{{.Name}}"},
}

func getRoutes(res Resource) []Route {
	var routes []Route
	for _, op := range operations {
		name := utils.FirstUpper(strings.ToLower(res.Name))
		route := Route{
			Method:      op.Method,
			Path:        strings.ReplaceAll(op.Path, "{{.Name}}", name),
			HandlerFunc: strings.ReplaceAll(op.FuncName, "{{.Name}}", name),
		}
		routes = append(routes, route)
	}
	return routes
}

func main() {
	namePtr := flag.String("name", "", "name of the resource")
	dirPtr := flag.String("dir", "", "directory to create the project in")
	// modulePtr := flag.String("module", "", "the package name of this project")
	flag.Parse()
	name, dir, module = *namePtr, *dirPtr, *dirPtr
	if name == "" || dir == "" {
		fmt.Println("Missing name argument. Usage: go run main.go -name <resource-name> -dir <directory>")
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("resource:%s  dir:%s", name, dir))

	_ = os.Mkdir(dir, 0755)
	modFilePath := filepath.Join(dir, "go.mod")
	if _, err := os.Stat(modFilePath); os.IsNotExist(err) {
		modTemplate := template.Must(template.ParseFiles("./templates/go.mod.tmpl"))

		modFile, err := os.Create(modFilePath)
		if err != nil {
			fmt.Printf("Failed to create file %s: %s\n", modFilePath, err.Error())
			os.Exit(1)
		}
		defer modFile.Close()

		modModuleName := fmt.Sprintf("%s", module)

		err = modTemplate.Execute(modFile, modModuleName)
		if err != nil {
			fmt.Printf("Failed to generate go.mod file %s: %s\n", modFilePath, err.Error())
			os.Exit(1)
		}

		fmt.Printf("Generated go.mod for %s in %s\n", "handlers", dir)
	}

	res := Resource{
		Name:       strings.Title(name),
		PluralName: name + "s",
	}
	routes := getRoutes(res)

	mainFilePath := filepath.Join(dir, "main.go")
	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		fmt.Printf("Failed to generate main file %s: %s\n", mainFilePath, err.Error())
		os.Exit(1)
	}
	defer mainFile.Close()

	mainTemplate := template.Must(template.New("main").Parse(templates.MainTemplate))
	err = mainTemplate.Execute(mainFile, struct {
		ModModuleName string
		Routes        []Route
	}{module, routes})
	if err != nil {
		fmt.Printf("Failed to generate main file %s: %s\n", mainFilePath, err.Error())
		os.Exit(1)
	}

	handlerDir := filepath.Join(dir, "handlers")
	err = os.Mkdir(handlerDir, 0755)
	if err != nil {
		fmt.Printf("Failed to create handlers directory %s: %s\n", handlerDir, err.Error())
		os.Exit(1)
	}
	for _, op := range operations {
		name := utils.FirstUpper(strings.ToLower(res.Name))
		// funcName := op.FuncName + res.Name
		funcName := strings.ReplaceAll(op.FuncName, "{{.Name}}", name)
		handlerFilePath := filepath.Join(handlerDir, strings.ToLower(funcName)+".go")
		handlerFile, err := os.Create(handlerFilePath)
		if err != nil {
			fmt.Printf("Failed to create file %s: %s\n", handlerFilePath, err.Error())
			os.Exit(1)
		}
		defer handlerFile.Close()

		handlerTemplate := template.Must(template.New("handler").Parse(templates.HandlerTemplate))
		err = handlerTemplate.Execute(handlerFile, Operation{op.Method, op.Path, funcName})
		if err != nil {
			fmt.Printf("Failed to generate handle file %s: %s\n", handlerFilePath, err.Error())
			os.Exit(1)
		}
	}
	fmt.Printf("Generated %s in %s \n", res.Name, dir)
}
