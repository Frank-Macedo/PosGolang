/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Category called with name:" + category)
		// fmt.Println("Category exists:" + fmt.Sprint(exists))
		// fmt.Println("ID of category" + fmt.Sprint(id))

	},
	// PreRun: func(ccmd *cobra.Command, args []string) {
	// 	fmt.Println("Called before run")
	// },
	// PostRun: func(ccmd *cobra.Command, args []string) {
	// 	fmt.Println("Called after run")
	// },
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	return fmt.Errorf("ocorreu um erro")
	// },
	// PreRunE: func(ccmd *cobra.Command, args []string) error {
	// 	return fmt.Errorf("ocorreu um erro")
	// },
	// PostRunE: func(ccmd *cobra.Command, args []string) error {
	// 	return fmt.Errorf("ocorreu um erro")
	// },
}

var category string
var exists bool
var id int32

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "Y", "Name of the category")
	categoryCmd.PersistentFlags().BoolVarP(&exists, "exists", "e", false, "Check if category exists")
	categoryCmd.PersistentFlags().Int32VarP(&id, "Id", "i", 0, "Id of the Category")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
