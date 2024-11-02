package project

import (
	"fmt"
	"os"
)

// CreateProject creates a new project directory
func CreateProject(projectName string) error {
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}
	return nil
}


func CreateDirectories(projectName string, directories []string) error {
	for _, directory := range directories {
		// Create the main directory for the project
		err := os.MkdirAll(projectName+"/"+directory, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", directory, err)
		}

		switch directory {
		case "auth":
			subdirs := []string{"routes", "models", "control"}
			for _, subdirectory := range subdirs {
				err := os.MkdirAll(projectName+"/"+directory+"/"+subdirectory, 0755)
				if err != nil {
					return fmt.Errorf("failed to create subdirectory %s/%s: %w", directory, subdirectory, err)
				}
				switch subdirectory {
				case "routes":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/authroutes.go")
					if err != nil {
						return fmt.Errorf("failed to create file authroutes.go: %w", err)
					}
				case "models":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/authmodel.go")
					if err != nil {
						return fmt.Errorf("failed to create file authmodel.go: %w", err)
					}
				case "control":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/authcontrol.go")
					if err != nil {
						return fmt.Errorf("failed to create file authcontrol.go: %w", err)
					}
				}
			}
		case "utils":
			subdirs := []string{"middleware", "jwt", "redis", "db"}
			for _, subdirectory := range subdirs {
				err := os.MkdirAll(projectName+"/"+directory+"/"+subdirectory, 0755)
				if err != nil {
					return fmt.Errorf("failed to create subdirectory %s/%s: %w", directory, subdirectory, err)
				}
				switch subdirectory {
				case "middleware":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/auth.go")
					if err != nil {
						return fmt.Errorf("failed to create file auth.go: %w", err)
					}
				case "jwt":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/jwt.go")
					if err != nil {
						return fmt.Errorf("failed to create file jwt.go: %w", err)
					}
				case "redis":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/redis.go")
					if err != nil {
						return fmt.Errorf("failed to create file redis.go: %w", err)
					}
				case "db":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/db.go")
					if err != nil {
						return fmt.Errorf("failed to create file db.go: %w", err)
					}
				}
			}
		case "cmd":
			subdirs := []string{"main"}
			for _, subdirectory := range subdirs {
				err := os.MkdirAll(projectName+"/"+directory+"/"+subdirectory, 0755)
				if err != nil {
					return fmt.Errorf("failed to create subdirectory %s/%s: %w", directory, subdirectory, err)
				}
				if subdirectory == "main" {
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/main.go")
					if err != nil {
						return fmt.Errorf("failed to create file main.go: %w", err)
					}
				}
			}
		}
	}
	return nil
}
