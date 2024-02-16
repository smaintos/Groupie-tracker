package Track

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

type API struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Data struct {
	Artist    API
	Relations Relation
}
type Artists1 struct {
	Artists []API
}

var ApiObject []API

func Artist(w http.ResponseWriter, r *http.Request) {
	Api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic(err)
	}
	ApiData, err := ioutil.ReadAll(Api.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(ApiData, &ApiObject)
	if err != nil {
		panic(err)
	}
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		panic(err)
	}
	filtre, _ := strconv.Atoi(r.FormValue("input-members"))
	filtre1, _ := strconv.Atoi(r.FormValue("input-date"))
	if filtre != 0 {
		var ApiObject1 []API
		for _, v := range ApiObject {
			if filtre == len(v.Members) && filtre1 <= v.CreationDate {
				ApiObject1 = append(ApiObject1, v)
			}
		}
		err = t.Execute(w, ApiObject1)
		if err != nil {
			panic(err)
		}
	} else {
		err = t.Execute(w, ApiObject)
		if err != nil {
			panic(err)
		}
	}
}
