package generator

import (
	"log"
	"os"
	"path/filepath"
)

func BaseDirectory(projectName string) {
	err := os.MkdirAll(projectName, 0755)
	if err != nil {
		log.Fatal(err)
	}

	dList := []string{"cmd", "api", "build", "deployments", "docs", "internal", "pkg", "scripts", "test"}
	for _, name := range dList {
		path := filepath.Join(projectName, name)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}
