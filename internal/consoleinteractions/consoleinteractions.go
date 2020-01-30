package consoleinteractions

import (
	"fmt"

	"github.com/zmb3/spotify"
)

func ChoosePlaylist(client *spotify.Client) (chosenPlaylist spotify.SimplePlaylist, err error) {
	playlistPage, err := client.CurrentUsersPlaylists()
	if err == nil {
		playlists := playlistPage.Playlists
		var playlist spotify.SimplePlaylist
		var i int
		fmt.Println("Choose the playlist you want to randomize...")
		for i, playlist = range playlists {
			fmt.Printf("%d - %s (%d songs)\n", i+1, playlist.Name, playlist.Tracks.Total)
		}
		fmt.Print("Type the number corresponding to the playlist : ")
		_, err = fmt.Scan(&i)
		if err == nil {
			chosenPlaylist = playlists[i-1]
			fmt.Printf("You chose %s (%d) !\n", chosenPlaylist.Name, i)
		}
	}
	return
}
