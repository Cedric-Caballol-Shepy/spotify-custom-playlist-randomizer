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

	trackIds, err := getAllTrackIds(client, playlist)
	if err != nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(trackIds), func(i, j int) { trackIds[i], trackIds[j] = trackIds[j], trackIds[i] })

	err = replaceAllPlaylistTracks(client, playlist, trackIds)

	if err == nil {
		fmt.Println("Shuffling done !")
	}
	return
}
