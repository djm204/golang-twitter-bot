package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func configure() *twitter.Client {
	apiKey := os.Getenv("TWITTER_CONSUMER_API_KEY")
	apiSecret := os.Getenv("TWITTER_CONSUMER_API_SECRET")
	accountToken := os.Getenv("TWITTER_ACCOUNT_ACCESS_TOKEN")
	accountSecret := os.Getenv("TWITTER_ACCOUNT_ACCESS_SECRET")
	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accountToken, accountSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client
}

func tweet(tweetText string, client *twitter.Client) bool {
	tweet, resp, err := client.Statuses.Update(tweetText, nil)

	fmt.Println(tweet)
	fmt.Println(resp.StatusCode)
	if err != nil {
		log.Fatal("Error tweeting")
		return false
	}

	return true
}

func main() {
	//Load ENV File
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Go-Twitter Bot v0.01")
	client := configure()
	tweet("did I get blocked?", client)

}
