package playlist

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/zmb3/spotify"
)

// ShuffleWithRules TODO DOC
func ShuffleWithRules(client *spotify.Client, playlist spotify.SimplePlaylist, minimalSpaceBetweenSameArtistSongs int) (err error) {
	fmt.Printf("Shuffling the playlist with a minimal space of %d between same artist songs (could "+
		"be unsatisfied if the playlist is too much unbalanced)...\n", minimalSpaceBetweenSameArtistSongs)

	tracks, err := getAllTracks(client, playlist)
	if err != nil {
		return
	}

	songIdsGroupedByArtist := make(map[spotify.ID][]spotify.ID)
	for _, track := range tracks {
		// songIdsGroupedByArtist[spotify.ID(track.Artists[0].Name)] = append(songIdsGroupedByArtist[spotify.ID(track.Artists[0].Name)], track.ID) // as a simplification, we only take the first artist
		songIdsGroupedByArtist[track.Artists[0].ID] = append(songIdsGroupedByArtist[track.Artists[0].ID], track.ID) // as a simplification, we only take the first artist
	}
	sortedListOfArtistStatsForShuffle := createListOfSortedArtistStatsForShuffle(songIdsGroupedByArtist)
	fmt.Println(sortedListOfArtistStatsForShuffle)
	fmt.Println(len(sortedListOfArtistStatsForShuffle))

	if err == nil {
		fmt.Println("Shuffling done !")
	}
	return
}

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

type artistStatsForShuffle struct {
	artistID     spotify.ID
	tracksAmount int
}

// createListOfSortedArtistStatsForShuffle creates a sorted list of artistStatsForShuffle by descending tracksAmount
func createListOfSortedArtistStatsForShuffle(songIdsGroupedByArtist map[spotify.ID][]spotify.ID) (artistStatsForShuffleList []artistStatsForShuffle) {
	for key, element := range songIdsGroupedByArtist {
		artistStatsForShuffleList = append(artistStatsForShuffleList, artistStatsForShuffle{artistID: key, tracksAmount: len(element)})
	}
	sort.Slice(artistStatsForShuffleList, func(i, j int) bool {
		return artistStatsForShuffleList[i].tracksAmount > artistStatsForShuffleList[j].tracksAmount
	})
	return
}
