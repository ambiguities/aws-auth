/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	aws_credentials_file string

	rootCmd = &cobra.Command{
		Use:   "aws-auth",
		Short: "A credential manager for the aws cli.",
		Long:  `A credential manager for the aws cli.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().StringVar(&aws_credentials_file, "aws-credentials-file", dirname+"/.aws/credentials", "aws cli credentials file (default is $HOME/.aws/credentials)")
}
