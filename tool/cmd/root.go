/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wailbentafat/go_cli_tool/tool/project"
)

var rootCmd = &cobra.Command{
	Use:   "project",
	Short: "Create a new project",
	Long:  `This command creates a new project with the specified name and initializes necessary directories.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logError(fmt.Errorf("project name is required"))
			return
		}
		projectName, err := cmd.Flags().GetString("project")
		if err != nil {
			logError(fmt.Errorf("error retrieving project name: %w", err))
			return
		}

		if projectName == "" {
			logError(fmt.Errorf("project name cannot be empty"))
			return
		}

		
		err = project.CreateProject(projectName)
		if err != nil {
			logError(fmt.Errorf("error creating project: %w", err))
			return
		}

		
		dirs := []string{"auth", "utils", "cmd"}
		err = project.CreateDirectories(projectName, dirs)
		if err != nil {
			logError(fmt.Errorf("error creating directories: %w", err))
			return
		}

		fmt.Println("Project and directories created successfully!")
	},
}

func init() {
	rootCmd.Flags().StringP("project", "p", "", "Project name (required)")
	rootCmd.MarkFlagRequired("project") // Marking the project flag as required
}

func logError(err error) {
	// Open log file
	file, e := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if e != nil {
		fmt.Printf("Failed to open log file: %v\n", e)
		return
	}
	defer file.Close()


	logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(err)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logError(err)
		os.Exit(1)
	}
}
