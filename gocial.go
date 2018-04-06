package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	flag "github.com/ogier/pflag"
)

var (
	port           string
	social_channel string
	post_id        string
	query          string
)

func print_tweet_response(resp_w http.ResponseWriter, tweet twitter.Tweet) {
	fmt.Fprint(resp_w, "{ Created At:"+string(tweet.CreatedAt)+
		", Favourite Count: "+strconv.Itoa(tweet.FavoriteCount)+
		", Status Id: "+strconv.FormatInt(tweet.ID, 10)+
		", InReplyToScreenName: "+tweet.InReplyToScreenName+
		", InReplyToStatusID : "+tweet.InReplyToStatusIDStr+
		", Language : "+tweet.Lang+
		", Retweet Count: "+strconv.Itoa(tweet.RetweetCount)+
		", Source: "+tweet.Source+
		", Text: "+tweet.Text+
		", User name: "+tweet.User.Name+
		"}")
}

func get_post_info(resp_w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {

		req.ParseForm()
		post_id := req.Form.Get("id")

		if post_id != "" {
			tweet, _, err := get_twitter_status(post_id)
			if err != nil {
				http.Error(resp_w, "Error getting details for post id: "+post_id,
					http.StatusInternalServerError)
			} else {
				print_tweet_response(resp_w, *tweet)
			}
		} else {
			http.Error(resp_w, "Invalid parameter. Must required id for post", http.StatusMethodNotAllowed)
		}
	} else {
		http.Error(resp_w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func searchTweet(resp_w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		req.ParseForm()
		query := req.Form.Get("query")

		if query != "" {
			search, _, err := searchTweets(query)
			if err != nil {
				http.Error(resp_w, "Error while searching the tweets", http.StatusInternalServerError)
			} else {
				for _, tweet := range search.Statuses {
					print_tweet_response(resp_w, tweet)
					fmt.Fprintln(resp_w, "")
				}
			}
		} else {
			http.Error(resp_w, "Please add a query term", http.StatusMethodNotAllowed)
		}
	} else {
		http.Error(resp_w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func init() {
	flag.StringVarP(&port, "port", "p", "8000", "Port No for the webservice")
	flag.StringVarP(&social_channel, "channel", "c", "twitter", "Social Network Channel Name")
	flag.StringVarP(&post_id, "post", "p", "", "Post id")
	flag.StringVarP(&query, "query", "q", "", "query term for search")
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 || (flag.NFlag() == 1 && port != "-1") { //As no flag is passed or only -p is passed, so webservice will be started
		fmt.Println("Web service started listening at post :8000")
		http.HandleFunc("/post", get_post_info)
		http.HandleFunc("/search", searchTweet)
		validatePort(port)
		http.ListenAndServe(":"+port, nil)
	} else {
		manage_cli()
	}
}

func validatePort(port string) {
	port_no, _ := strconv.ParseInt(port, 10, 16)
	if port_no >= 1024 && port_no <= 49151 {

	} else {
		panic("Invalid port no. Should be in range(1024,49151)")
	}
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error Ocurred")
		os.Exit(1)
	}
}

func manage_cli() {
	flag.Parse()

	if post_id != "" {
		tweet, _, err := get_status(social_channel, post_id)
		errorHandler(err)
		fmt.Println(tweet)
	} else if query != "" {
		search, _, err := searchQuery(social_channel, query)
		errorHandler(err)
		fmt.Println(search)

	} else {
		fmt.Println("Not enough values supplied")
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
