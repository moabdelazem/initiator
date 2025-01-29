package utils

import "fmt"

func createNodeProject(project *Project) error {
	fmt.Printf("Creating Node.js project: %s\n", project.Name)
	return nil
}

func createGoProject(project *Project) error {
	fmt.Printf("Creating Go project: %s\n", project.Name)
	return nil
}
