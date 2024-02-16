package Track

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var LocationsObject Relation

func Detail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	Api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id + "")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	ApiData, err := ioutil.ReadAll(Api.Body)
	if err != nil {
		log.Fatal(err)
	}

	var OneApiObject API
	err = json.Unmarshal(ApiData, &OneApiObject)
	if err != nil {
		panic(err)
	}
	ApiRel, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id + "")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	ApiDataRel, err := ioutil.ReadAll(ApiRel.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(ApiDataRel, &LocationsObject)
	if err != nil {
		panic(err)
	}
	Data := Data{OneApiObject, LocationsObject}

	t, err := template.ParseFiles("./templates/artiste.html")
	t.Execute(w, Data)

}
