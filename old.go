package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var consumerKey *string = flag.String("consumerKey", "", "Consumer key from Yahoo App")
var consumerSecret *string = flag.String("consumerSecret", "", "Consumer secret from Yahoo app")

func main() {
	flag.Parse()

	if len(*consumerKey) == 0 || len(*consumerSecret) == 0 {
		fmt.Println("You must set the --consumerKey and --consumerSecret flags.")
		fmt.Println("---")
		//Usage()
		os.Exit(1)
	}

	consumer := oauth.NewConsumer(
		*consumerKey,    //consumer key
		*consumerSecret, //consumer secret
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.login.yahoo.com/oauth/v2/get_request_token",
			AuthorizeTokenUrl: "https://api.login.yahoo.com/oauth/v2/request_auth",
			AccessTokenUrl:    "https://api.login.yahoo.com/oauth/v2/get_token",
		})

	//consumer.Debug(true)

	requestToken, url, err := consumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("(1) Go to: " + url)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here: ")
	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := consumer.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	client, err := consumer.MakeHttpClient(accessToken)
	if err != nil {
		log.Fatal(err)
	}

	urlBase := "https://query.yahooapis.com/v1/yql?q="
	urlQuery := "select%20*%20from%20fantasysports.players.stats%20where%20league_key%3D'238.l.627060'%20and%20player_key%3D'238.p.6619'"
	fullURL := urlBase + urlQuery + "&format=json"


	fmt.Println("Using: ", fullURL)

	response, err := client.Get(fullURL)
  if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	contents, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Got ", string(contents))
	//getJson(fullURL,nil)
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("Can't retrive url '%s'.\nError: %s", url, err)
	}

	defer r.Body.Close()

	contents, _ := ioutil.ReadAll(r.Body)
	fmt.Println("URL: ", url)
	fmt.Println("Response:", string(contents))

	return json.NewDecoder(r.Body).Decode(target)
}
