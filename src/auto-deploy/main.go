package main

import (
	"fmt"
	"os"

	"github.com/peterwillcn/ai-launch/libs"
)

func main() {
	container := libs.NewContainer()
	workDir := container.WorkDir
	err := os.MkdirAll(workDir+"/data", 0755)
	if err != nil {
		fmt.Printf("Error creating directory %s: %s\n", workDir, err)
	}
	libs.MakeConfig(workDir)
	templates := libs.GetTemp()
	for _, manifest := range container.ManiFests {
		libs.ApplyTemplate(workDir, manifest, templates[manifest], container)
	}
	switch container.Operation {
	case "start":
		libs.Start(workDir, container.ManiFests)
	case "stop":
		libs.Stop(workDir, container.ManiFests)
	case "backup":
		fmt.Printf("developing")
	default:
		fmt.Printf("Unknown operation: [ %s ], \nexample: [ ./launch -o start (stop, backup) ] \n", container.Operation)
	}
}
