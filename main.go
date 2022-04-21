package main

import (
	"fmt"

	"github.com/dghubble/go-twitter"
	"github.com/dghubble/oauth1"
)

func main(){
	CONSUMER_KEY = "2Nmc8tm1YxkLAqU8JtyYpZqMK"
	CONSUMER_SECRET_KEY = "xQcd9lBMXElN3q4LWUQToxmN2bBbcsUNWka3D2wOl7sOm9SdJR"
	TOKEN_KEY = "1517181806726033408-9KG1dcdacVJ8a4sZwwbdemZw5hGsIW"
	TOKEN_SECRET_KEY = "YUp09dSymvNVJXlI4mCB31x3voycMJTNEzvO1Kmr7PSYi"

	config := oauth1.NewConfig(CONSUMER_KEY, CONSUMER_SECRET_KEY)
	token := oauth1.NewToken(TOKEN_KEY, TOKEN_SECRET_KEY)

	httpClient := config.Client(oauth1.NoContext,  token)
	client := twitter.NewClient(httpClient)

	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil{
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Account: @%s (%s)\n", user.ScreenName, user.Name)
}