package googleDrive

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func GetCredentials() []byte {
	cred, err := os.ReadFile("credentials.json")

	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	return cred
}

func OauthInit(cred []byte) *oauth2.Config {
	config, err := google.ConfigFromJSON(cred, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	config.RedirectURL = "urn:ietf:wg:oauth:2.0:oob"

	return config
}

func FetchToken(config *oauth2.Config) *oauth2.Token {
	cacheTokenFile := tokenCacheFile()
	token, err := tokenFromFile(cacheTokenFile)
	if err != nil {
		token = getTokenFromWeb(config)
		saveToken(cacheTokenFile, token)
	}

	return token
}

func CreateDriveService(config *oauth2.Config, context context.Context, token *oauth2.Token) *drive.Service {
	client := config.Client(context, token)
	service, err := drive.NewService(context, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	return service
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string
	fmt.Print("Enter the authorization code: ")
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	token, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	return token
}

func saveToken(path string, token *oauth2.Token) {
	// creates directory if non existent
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		log.Fatalf("Unable to create directory for token file: %v", err)
	}

	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache OAuth token: %v", err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(token)
}

func tokenCacheFile() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get current user: %v", err)
	}

	return filepath.Join(user.HomeDir, ".credentials", "drive-go-quickstart.json")
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)

	return token, err
}
