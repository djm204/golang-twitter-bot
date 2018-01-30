package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func configure() {
	apiKey := os.Getenv("TWITTER_CONSUMER_API_KEY")
	apiSecret := os.Getenv("TWITTER_CONSUMER_API_SECRET")
	accountToken := os.Getenv("TWITTER_ACCOUNT_ACCESS_TOKEN")
	accountSecret := os.Getenv("TWITTER_ACCOUNT_ACCESS_SECRET")
	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accountToken, accountSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	demux := twitter.NewSwitchDemux()

	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}

	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}

	fmt.Println("starting stream....")

	//FILTER
	filterParams := &twitter.StreamFilterParams{
		Track:         []string{"crypto"},
		StallWarnings: twitter.Bool(true),
	}

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	go demux.HandleChan(stream.Messages)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	log.Println(<-ch)

	fmt.Println("Stopping stream...")
	stream.Stop()
}

func main() {
	//Load ENV File
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Go-Twitter Bot v0.01")
	configure()
}
