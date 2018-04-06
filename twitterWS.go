package main

// OAuth1
import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func (twitterClient *TwitterClient) authenticate(oauthInputs OauthInputs) *twitter.Client {

	config := oauth1.NewConfig(oauthInputs.consumerKey, oauthInputs.consumerSecret)
	token := oauth1.NewToken(oauthInputs.accessToken, oauthInputs.accessSecret)
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	return client
}

func (twitterClient *TwitterClient) get_twitter_status(post_id string) {
	//return twitterClient
}
