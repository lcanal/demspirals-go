package controllers

import (
	"log"
	"io/ioutil"
	"demspirals/auths"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// To explore queries and code, use the Yahoo  YQL Console
	// https://developer.yahoo.com/yql/console/

	urlBase := "https://query.yahooapis.com/v1/yql?q="
	urlQuery := "select%20*%20from%20fantasysports.players.stats%20where%20league_key%3D'238.l.627060'%20and%20player_key%3D'238.p.6619'"
	fullURL := urlBase + urlQuery + "&format=json"

	response, err := auths.AuthorizedClient.Get(fullURL)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	c.Data["yahooResponse"] = string(contents)
	c.TplName = "index.tpl"
}
