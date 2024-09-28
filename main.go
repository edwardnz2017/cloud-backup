package main

import (
	"cloud-backup/pkg/googleDrive"
	"context"
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()

	var fileName string

	rootCmd := &cobra.Command{
		Use:   "backup",
		Short: "CLI tool for cloud file backups",
		Run: func(cmd *cobra.Command, args []string) {
			cred := googleDrive.GetCredentials()

			config := googleDrive.OauthInit(cred)

			token := googleDrive.FetchToken(config)

			service := googleDrive.CreateDriveService(config, ctx, token)

			// fileList, err := service.Files.List().
			// 	PageSize(10).
			// 	Fields("nextPageToken, files(id, name)").
			// 	Do()
			// if err != nil {
			// 	log.Fatalf("Unable to retrieve files: %v", err)
			// }

			// log.Println("Files:")
			// if len(fileList.Files) == 0 {
			// 	log.Println("No files found.")
			// } else {
			// 	for _, file := range fileList.Files {
			// 		log.Printf("%s (%s)\n", file.Name, file.Id)
			// 	}
			// }
		},
	}

	rootCmd.Flags().StringVarP(&fileName, "fileName", "F", "", "File name")
	rootCmd.MarkFlagRequired("fileName")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
