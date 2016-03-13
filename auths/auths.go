package auths

import(
  "log"
  "fmt"
  "os"
  "flag"
  "net/http"
  "github.com/mrjones/oauth"
  "github.com/astaxie/beego"
)

var AuthorizedClient http.Client

func AuthorizeApp(consumerKey,consumerSecret string) {
  consumer := oauth.NewConsumer(
		consumerKey,    //consumer key
	  consumerSecret, //consumer secret
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.login.yahoo.com/oauth/v2/get_request_token",
			AuthorizeTokenUrl: "https://api.login.yahoo.com/oauth/v2/request_auth",
			AccessTokenUrl:    "https://api.login.yahoo.com/oauth/v2/get_token",
		})

	//consumer.Debug(true)

	requestToken, url, err := consumer.GetRequestTokenAndUrl(beego.AppConfig.String("domain"))
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

func FindCreds() (string,string,error) {
  //Find keys in this order, top level wins:
  //1) Find in parsed arguments given
  //2) Find in environment variables
  //3) Find in .conf file

  //1) Find in parsed arguments given
  var consumerKey *string = flag.String("consumerKey", "", "Consumer key from Yahoo App")
  var consumerSecret *string = flag.String("consumerSecret", "", "Consumer secret from Yahoo app")
  flag.Parse()

  //2) Find in environment variables
  envconsumerKey := os.Getenv("cKey")
  envconsumerSecret := os.Getenv("cSecret")

  //3) Find in .conf file
  confconsumerKey := beego.AppConfig.String("cKey")
  confconsumerSecret := beego.AppConfig.String("cSecret")

  //Return any you find
  var finalKey string
  var finalSecret string
  var finalError error

  if (len(*consumerKey) != 0)        { finalKey = *consumerKey }
  if (len(*consumerSecret) != 0)     { finalSecret = *consumerSecret }
  if (len(envconsumerKey) != 0)      { finalKey = envconsumerKey }
  if (len(envconsumerSecret) != 0)   { finalSecret = envconsumerSecret }
  if (len(confconsumerKey) != 0 )    { finalKey = confconsumerKey }
  if (len(confconsumerSecret) != 0 ) { finalSecret = confconsumerSecret }

  if (len(finalKey) == 0 || len(finalSecret) == 0) { finalError = fmt.Errorf("No keys found.") }

  return finalKey,finalSecret,finalError
}
