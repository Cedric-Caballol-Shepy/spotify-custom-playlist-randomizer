package playlist

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zmb3/spotify"
)

// ShuffleRandom randomly shuffles a Spotify playlist
func ShuffleRandom(client *spotify.Client, playlist spotify.SimplePlaylist) (err error) {
	fmt.Println("Shuffling randomly the playlist...")
	trackPage, err := client.GetPlaylistTracks(playlist.ID)
	if err != nil {
		return
	}

	var trackIds []spotify.ID
	for _, track := range trackPage.Tracks {
		trackIds = append(trackIds, track.Track.ID)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(trackIds), func(i, j int) { trackIds[i], trackIds[j] = trackIds[j], trackIds[i] })

	oldI, i := 0, 99
	for {
		trackIdsSlice := trackIds[oldI:i]
		err = client.ReplacePlaylistTracks(playlist.ID, trackIdsSlice...)
		if i < len(trackIds) || err != nil {
			break
		}
		oldI = i
		i += 100
	}

	if err == nil {
		fmt.Println("Shuffling done !")
	}
	return
}
