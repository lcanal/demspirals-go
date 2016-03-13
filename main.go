package main

import (
	"fmt"
	"os"
	_ "demspirals/routers"
  "demspirals/auths"
	"github.com/astaxie/beego"
)

func main() {
	consumerKey,consumerSecret,err := auths.FindCreds()

	if (err != nil ){
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	auths.AuthorizeApp(consumerKey,consumerSecret)
	beego.Run()
}
