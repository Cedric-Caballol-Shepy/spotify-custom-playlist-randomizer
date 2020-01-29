// code took from https://github.com/zmb3/spotify/blob/master/examples/authenticate/authcode/authenticate.go then changed a bit to fit with my app
package authentication

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zmb3/spotify"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate)
	ch    = make(chan *spotify.Client)
	state = "abc123" // TODO (if I put this on a "real" server) : change that to a more securized version
)

// In order to run this by yourself you need to get your clientID and secretKey :
//  1. Register an application at: https://developer.spotify.com/my-applications/
//       - Use "http://localhost:8080/callback" as the redirect URI
//  2. Set the clientID  to the client ID you got in step 1.
//  3. Set the secretKey to the client secret from step 1.
func Authenticate(clientID, secretKey string) *spotify.Client {
	auth.SetAuthInfo(clientID, secretKey)
	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(fmt.Sprintf("You are logged in as: %s\n\n", user.ID))
	return client
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}
