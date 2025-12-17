/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	volumeName string
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("backup called")

		dockerCmd := exec.Command("docker", "ps", "-aq", "--filter", "volume="+volumeName)
		fmt.Print(dockerCmd)
		fmt.Print("\n")

		stdout, err := dockerCmd.Output()
		if err != nil {
			return fmt.Errorf("docker ps failed: %w", err)
		}

		fmt.Print("\n")
		fmt.Print(stdout)
		fmt.Print("\n")
		lines := strings.Split(strings.TrimSpace(string(stdout)), "\n")
		fmt.Print(lines)
		countContainers := len(lines)
		if countContainers > 1 {
			return fmt.Errorf("Architectural violation! volume wird vermutlich von mehreren containern verwendet!")
		}
		fmt.Print("\n")
		fmt.Print(countContainers)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVar(&volumeName, "volume", "", "Docker volume you want to back up")
}
