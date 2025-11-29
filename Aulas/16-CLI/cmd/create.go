/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Frank-Macedo/16-CLI/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples`,
		RunE:  runCreate(categoryDB),
	}
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {

		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		categoryDb.Create(name, description)
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringP("name", "n", "xpto", "Category name")
	createCmd.PersistentFlags().StringP("description", "d", "yxz", "Category desciption")

	createCmd.MarkFlagsRequiredTogether("name", "description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
