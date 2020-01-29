package consoleinteractions

import (
	"fmt"
	"log"

	"github.com/zmb3/spotify"
)

func ChoosePlaylist(client *spotify.Client) (chosenPlaylist spotify.SimplePlaylist) {
	playlistPage, err := client.CurrentUsersPlaylists()
	if err != nil {
		log.Fatal(err)
	}
	playlists := playlistPage.Playlists
	var playlist spotify.SimplePlaylist
	var i int
	fmt.Println("Choose the playlist you want to randomize")
	for i, playlist = range playlists {
		fmt.Println(fmt.Sprintf("%d - %s (%d songs)", i+1, playlist.Name, playlist.Tracks.Total))
	}
	fmt.Print("Type the number corresponding to the playlist : ")
	_, err = fmt.Scan(&i)
	if err != nil {
		log.Fatal(err)
	}
	chosenPlaylist = playlists[i-1]
	fmt.Println(fmt.Sprintf("You chose %s (%d) !", chosenPlaylist.Name, i))
	return
}
