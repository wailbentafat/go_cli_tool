/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wailbentafat/go_cli_tool/tool/project"
)




var rootCmd = &cobra.Command{
	Use:   "project",
	Short: "kifkif jya7a",
	Run: func (cmd *cobra.Command, args []string) { 
		fmt.Println("Hello World!")
    initproject()
		
	},
}
func init(){
	rootCmd.Flags().StringP("project","p","project_name","project name")
}
func initProject() {
	projectName := rootCmd.Flag("project").Value.String()
	project.CreateProject(projectName)
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


