package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
)

type WebClient interface {
	authenticate(consumerKey string, consumerSecret string, accessToken string, accessSecret string)
}

type OauthInputs struct {
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
}

func AddOauthInputs(consumerKey string, consumerSecret string, accessToken string, accessSecret string) OauthInputs {
	return OauthInputs{consumerKey: consumerKey, consumerSecret: consumerSecret, accessToken: accessToken, accessSecret: accessSecret}
}

type TwitterClient struct {
	client *twitter.Client
}

func getClient() *twitter.Client {
	twitterClient := TwitterClient{}
	twitterClient.client = twitterClient.authenticate(AddOauthInputs("afUEBRPPT1AXtZtCUdM0XGgyn", "9z8ptVgxO0YZApNHOzRAOmGPy9f06tscV3WejIuFKxMaHFZm4Z", "2691889945-fkeulM3oT1OsXD4lzVfgRGR1bUx3wnHJNaNsHcM", "QTrvEvenXttAz4KSfRU6SKKtTfyZnJe8kt4lp7RA5a22Z"))
	return twitterClient.client
}

type WebAdaptee interface {
	get_status(social_channel string, post_id string)
}

type TwitterAdaptee struct {
}

func get_twitter_status(post_id string) (*twitter.Tweet, *http.Response, error) {
	post_id_int64, _ := strconv.ParseInt(post_id, 10, 64)
	return getClient().Statuses.Show(post_id_int64, nil)
}

func get_status(social_channel string, post_id string) (*twitter.Tweet, *http.Response, error) {
	if social_channel != "" {

		//need refactoring here
		if social_channel != "twitter" {
		} else {
		}
	} else {
		fmt.Println(errors.New("social channel cannot be empty"))
		//need to stop here
	}
	post_id_int64, _ := strconv.ParseInt(post_id, 10, 64)
	return getClient().Statuses.Show(post_id_int64, nil)

}

func searchTweets(query string) (*twitter.Search, *http.Response, error) {
	searchParamObject := &twitter.SearchTweetParams{
		Query: query,
	}
	return getClient().Search.Tweets(searchParamObject)
}

func searchQuery(social_channel string, query string) (*twitter.Search, *http.Response, error) {
	if social_channel != "" {

		//need refactoring here
		if social_channel != "twitter" {
		} else {
		}
	} else {
		fmt.Println(errors.New("social channel cannot be empty"))
		//need to stop here
	}
	searchParamObject := &twitter.SearchTweetParams{
		Query: query,
	}
	return getClient().Search.Tweets(searchParamObject)
}
