package auths

import(
  "log"
  "fmt"
  "net/http"
  "github.com/mrjones/oauth"
)

var AuthorizedClient http.Client

func AuthorizeApp(consumerKey,consumerSecret *string) {
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

  AuthorizedClient = *client
}
