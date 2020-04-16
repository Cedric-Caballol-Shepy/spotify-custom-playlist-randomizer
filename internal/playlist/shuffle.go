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

	trackIds, err := getAllTrackIds(client, trackPage)
	if err != nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(trackIds), func(i, j int) { trackIds[i], trackIds[j] = trackIds[j], trackIds[i] })

	maxPerRequest := 100 // Spotify's API can't take more than 100 tracks by call
	if len(trackIds) < maxPerRequest {
		err = client.ReplacePlaylistTracks(playlist.ID, trackIds...)
	} else {
		// replace the 100 first tracks (more optimized than removing every tracks)
		err = client.ReplacePlaylistTracks(playlist.ID, trackIds[:maxPerRequest]...)
		if err != nil {
			return
		}
		// as the other tracks have been deleted by ReplacePlaylistTracks, add them...
		for i := maxPerRequest; i < len(trackIds); i += maxPerRequest {
			end := i + maxPerRequest
			if end > len(trackIds) {
				end = len(trackIds)
			}
			_, err = client.AddTracksToPlaylist(playlist.ID, trackIds[i:end]...)
			if err != nil {
				return
			}
		}
	}

	if err == nil {
		fmt.Println("Shuffling done !")
	}
	return
}

// getAllTrackIds returns a list of all the track ids from a playlist starting from the page trackPage to the last page
func getAllTrackIds(client *spotify.Client, trackPage *spotify.PlaylistTrackPage) (trackIds []spotify.ID, err error) {
	for page := 1; ; page++ {
		for _, track := range trackPage.Tracks {
			trackIds = append(trackIds, track.Track.ID)
		}
		err = client.NextPage(trackPage)
		if err == spotify.ErrNoMorePages {
			err = nil
			break
		}
		if err != nil {
			return
		}
	}
	return
}
