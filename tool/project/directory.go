package project

import (
	"log"
	"os"

	"golang.org/x/text/cases"
)



func CreateProject(projectName string) {
	err:=os.Mkdir(projectName, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDirectories(projectName string, directories []string) {
	 for _, directory := range directories {
		switch directory{
		case "auth":
			subdir:=[]string{"routes","models","control"}
			for _, subdirectory := range subdir {
				err := os.MkdirAll(projectName+"/"+directory+"/"+subdirectory, 0755)
				if err != nil {
					log.Fatal(err)
				}
				switch subdirectory{
				case "routes":
					err:=os.Create(projectName+"/"+directory+"/"+subdirectory+"/authroutes.go")
					if err != nil {
						log.Fatal(err)
					}
				case "models":
					err:=os.Create(projectName+"/"+directory+"/"+subdirectory+"/authmodel.go")
					if err != nil {
						log.Fatal(err)
					}
				case "control":
					err:=os.Create(projectName+"/"+directory+"/"+subdirectory+"/authcontrol.go")
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			case "utils":
			subdir:=[]string{"middleweare","jwt","redis","db"}
			for _, subdirectory := range subdir {
				
			}
		}
		err := os.MkdirAll(projectName+"/"+directory, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

