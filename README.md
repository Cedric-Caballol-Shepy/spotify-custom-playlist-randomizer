# cprandomizer (WIP)

cprandomizer is a commandline app to randomize your playlists with custom random functions set with rules. It's a WIP so the definition of "rules" will come later ! 😉

## Prerequisites

Before launching the app, get your Spotify developper credentials : 
* Register an application at: https://developer.spotify.com/my-applications/
   - Use "http://localhost:8080/callback" as the redirect URI
* Note your client ID 
* Note your client secret

## Launch the app

`go run main.go --client-id <your_client_ID> --secret-key <your_client_secret>`