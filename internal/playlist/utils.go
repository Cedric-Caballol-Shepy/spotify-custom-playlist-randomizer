package playlist

import (
	"github.com/zmb3/spotify"
)

var maxSongsPerRequest = 100 // Spotify's API can't take more than 100 tracks by call

// replaceAllPlaylistTracks overwrites the playlist with the list of track IDs
func replaceAllPlaylistTracks(client *spotify.Client, playlist spotify.SimplePlaylist, trackIds []spotify.ID) (err error) {
	if len(trackIds) < maxSongsPerRequest {
		err = client.ReplacePlaylistTracks(playlist.ID, trackIds...)
	} else {
		// replace the 100 first tracks (removing the others if the length of the playlist is greater than 100)
		err = client.ReplacePlaylistTracks(playlist.ID, trackIds[:maxSongsPerRequest]...)
		if err != nil {
			return
		}
		// as the other tracks have been deleted by ReplacePlaylistTracks, add them
		for i := maxSongsPerRequest; i < len(trackIds); i += maxSongsPerRequest {
			end := i + maxSongsPerRequest
			if end > len(trackIds) {
				end = len(trackIds)
			}
			_, err = client.AddTracksToPlaylist(playlist.ID, trackIds[i:end]...)
			if err != nil {
				return
			}
		}
	}
	return
}

// getAllTrackIds returns a list of all the track ids from a playlist
func getAllTrackIds(client *spotify.Client, playlist spotify.SimplePlaylist) (trackIds []spotify.ID, err error) {
	trackPage, err := client.GetPlaylistTracks(playlist.ID)
	if err == nil {
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
	}
	return
}
