/*
Copyright © 2025 Bartholomaeuss
*/
package cmd

import (
	"dockvol/core/backup"
	"fmt"

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

		if err := backup.Backup(volumeName); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVar(&volumeName, "volume", "", "Docker volume you want to back up")
}
