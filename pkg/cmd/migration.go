package cmd

import (
	"github.com/spf13/cobra"
	"test.com/pkg/migration"
)

func init() {
	rootCmd.AddCommand(migrationCmd)
}

var migrationCmd = &cobra.Command{
	Use:   "migrat",
	Short: "Database migration",
	Long:  "Database table creation",
	Run: func(cmd *cobra.Command, args []string) {
		migration.Migrate()
	},
}
