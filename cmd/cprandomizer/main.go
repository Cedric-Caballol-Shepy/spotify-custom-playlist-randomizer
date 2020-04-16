package main

import (
	"cprandomizer/internal/authentication"
	"cprandomizer/internal/consoleinteractions"
	"cprandomizer/internal/playlist"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("This is a work in progress, be careful and review the code before using it on your playlists !")
	var clientID, secretKey string
	flag.StringVar(&clientID, "client-id", "", "Your Spotify client ID")
	flag.StringVar(&secretKey, "secret-key", "", "Your Spotify client secret")
	flag.Parse()
	if clientID == "" || secretKey == "" {
		fmt.Println("clientID and secretKey needed. See README for more details.")
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		client, err := authentication.Authenticate(clientID, secretKey)
		if err != nil {
			log.Fatal(err)
		}

		chosenPlaylist, err := consoleinteractions.ChoosePlaylist(client)
		if err != nil {
			log.Fatal(err)
		}

		err = playlist.ShuffleRandom(client, chosenPlaylist)
		if err != nil {
			log.Fatal(err)
		}
	}
}
