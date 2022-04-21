package main

import (
	"os"

	"github.com/dghubble/go-twitter"
	"github.com/dghubble/oauth1"
)

func main(){
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET_KEY"))
	token := oauth1.NewToken(os.Getenv("TOKEN_KEY"), os.Getenv("TOKEN_SECRET_KEY"))
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	
}