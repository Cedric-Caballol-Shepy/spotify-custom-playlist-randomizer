package main

import (
	"cprandomizer/internal/authentication"
	"flag"
	"fmt"
	"github.com/zmb3/spotify"
	"log"
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
		playlistPage, err := client.CurrentUsersPlaylists()

		if err != nil {
			log.Fatal(err)
		}

		var playlist spotify.SimplePlaylist
		for _, playlist = range playlistPage.Playlists {
			var trackPage *spotify.PlaylistTrackPage
			trackPage, err = client.GetPlaylistTracks(playlist.ID)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(fmt.Sprintf("%s : ", playlist.Name))
			var playlistTrack spotify.PlaylistTrack
			for _, playlistTrack = range trackPage.Tracks {
				fmt.Println(fmt.Sprintf("\t%s - %s", playlistTrack.Track.Artists, playlistTrack.Track.Name))
			}
		}
	}
}
