module twitter-bot

go 1.19

require internal/twitter v1.0.0

replace internal/twitter => ./internal/twitter

require internal/utils v1.0.0

require (
	github.com/dghubble/oauth1 v0.7.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/oauth2 v0.2.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace internal/utils => ./internal/utils
