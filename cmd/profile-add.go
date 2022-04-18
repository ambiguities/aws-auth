/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	access_key_id     string
	secret_access_key string
	session_token     string

	profileAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new profile.",
		Long:  `Add a new profile to aws credentials.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add called")
		},
	}
)

func init() {
	profileCmd.AddCommand(profileAddCmd)

	profileAddCmd.Flags().StringVar(&access_key_id, "access_key_id", "", "access key id (required)")
	profileAddCmd.MarkFlagRequired("access_key_id")
	profileAddCmd.Flags().StringVar(&secret_access_key, "secret_access_key", "", "secret access key (required)")
	profileAddCmd.MarkFlagRequired("secret_access_key")

	// Here you will define your flags and profileuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
