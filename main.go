package main

import (
	"flag"
	"fmt"
	"os"
	_ "demspirals/routers"
  "demspirals/auths"
	"github.com/astaxie/beego"
)

func main() {
	var consumerKey *string = flag.String("consumerKey", "", "Consumer key from Yahoo App")
	var consumerSecret *string = flag.String("consumerSecret", "", "Consumer secret from Yahoo app")

	flag.Parse()

	if len(*consumerKey) == 0 || len(*consumerSecret) == 0 {
		fmt.Println("You must set the --consumerKey and --consumerSecret flags.")
		fmt.Println("---")
		os.Exit(1)
	}

	auths.AuthorizeApp(consumerKey,consumerSecret)
	beego.Run()
}
