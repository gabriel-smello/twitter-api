package main

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	CONSUMER_KEY := "2Nmc8tm1YxkLAqU8JtyYpZqMK"
	CONSUMER_SECRET_KEY := "xQcd9lBMXElN3q4LWUQToxmN2bBbcsUNWka3D2wOl7sOm9SdJR"
	TOKEN_KEY := "1517181806726033408-9KG1dcdacVJ8a4sZwwbdemZw5hGsIW"
	TOKEN_SECRET_KEY := "YUp09dSymvNVJXlI4mCB31x3voycMJTNEzvO1Kmr7PSYi"

	config := oauth1.NewConfig(CONSUMER_KEY, CONSUMER_SECRET_KEY)
	token := oauth1.NewToken(TOKEN_KEY, TOKEN_SECRET_KEY)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	/*	tweet, _, err := client.Statuses.Update("uma mensagem qualquer", nil)
		if err != nil {
			log.Fatal((err))
		}

		log.Print(tweet.Text)
	*/

	tweets, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:  "guerra",
		Count:  5,
		Locale: "BR",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range tweets.Statuses {
		log.Print("Nome do Usuario: ", val.User.Name)
		log.Print("Tweet: ", val.Text, "\n\n")
	}
}
