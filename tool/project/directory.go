package project

import (
	"log"
	"os"

	
)



func CreateProject(projectName string) {
	err:=os.Mkdir(projectName, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDirectories(projectName string, directories []string) {
	for _, directory := range directories {
		switch directory {
		case "auth":
			subdir := []string{"routes", "models", "control"}
			for _, subdirectory := range subdir {
				err := os.MkdirAll(projectName+"/"+directory+"/"+subdirectory, 0755)
				if err != nil {
					log.Fatal(err)
				}
				switch subdirectory {
				case "routes":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/authroutes.go")
					if err != nil {
						log.Fatal(err)
					}
				case "models":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/authmodel.go")
					if err != nil {
						log.Fatal(err)
					}
				case "control":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/authcontrol.go")
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		case "utils":
			subdir := []string{"middleweare", "jwt", "redis", "db"}
			for _, subdirectory := range subdir {
				err := os.MkdirAll(projectName+"/"+directory+"/"+subdirectory, 0755)
				if err != nil {
					log.Fatal(err)
				}
				switch subdirectory {
				case "middleweare":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/auth.go")
					if err != nil {
						log.Fatal(err)
					}
				case "jwt":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/jwt.go")
					if err != nil {
						log.Fatal(err)
					}
				case "redis":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/redis.go")
					if err != nil {
						log.Fatal(err)
					}
				case "db":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/db.go")
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		case "cmd":
			subdir := []string{"main"}
			for _, subdirectory := range subdir {
				err := os.MkdirAll(projectName+"/"+directory+"/"+subdirectory, 0755)
				if err != nil {
					log.Fatal(err)
				}
				switch subdirectory {
				case "main":
					_, err := os.Create(projectName + "/" + directory + "/" + subdirectory + "/main.go")
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
		err := os.MkdirAll(projectName+"/"+directory, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}