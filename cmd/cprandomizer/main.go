package main

import (
	"cprandomizer/internal/authentication"
	"cprandomizer/internal/consoleinteractions"
	"flag"
	"fmt"
	"os"
)

func main() {
	var clientID, secretKey string
	flag.StringVar(&clientID, "client-id", "", "Your Spotify client ID")
	flag.StringVar(&secretKey, "secret-key", "", "Your Spotify client secret")
	flag.Parse()
	if clientID == "" || secretKey == "" {
		fmt.Println("clientID and secretKey needed. See README for more details.")
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		client := authentication.Authenticate(clientID, secretKey)
		/*chosenPlaylist :=*/ consoleinteractions.ChoosePlaylist(client)
		//WIP

	}
}
