package main

import (
	"cloud-backup/pkg/googleDrive"
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

func main() {
	var fileName string

	rootCmd := &cobra.Command{
		Use:   "backup",
		Short: "CLI tool for cloud file backups",
		Run: func(cmd *cobra.Command, args []string) {
			cred := googleDrive.GetCredentials()
			config := googleDrive.OauthInit(cred)
			token := googleDrive.FetchToken(config)
			fmt.Println(token)
		},
	}

	rootCmd.Flags().StringVarP(&fileName, "fileName", "F", "", "File name")
	rootCmd.MarkFlagRequired("fileName")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
