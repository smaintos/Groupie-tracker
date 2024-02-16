package Track

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type SearchBar struct {
	Artist    API
	SearchBar bool
}

var SearchBar2 SearchBar

func Search(w http.ResponseWriter, r *http.Request) {
	searchBar := r.FormValue("q")
	for i := 0; i < len(ApiObject); i++ {
		name := strings.ToLower(ApiObject[i].Name)
		searchBar = strings.ToLower(searchBar)
		album := strings.ToLower(ApiObject[i].FirstAlbum)
		creationDate := strconv.Itoa(ApiObject[i].CreationDate)
		locations := strings.ToLower(ApiObject[i].Locations)
		for z := 0; z < len(ApiObject[i].Members); z++ {
			members := strings.ToLower(ApiObject[i].Members[z])
			searchBar = strings.ToLower(searchBar)
			if name == searchBar || album == searchBar || creationDate == searchBar || members == searchBar || locations == searchBar {
				SearchBar2 = SearchBar{
					Artist:    ApiObject[i],
					SearchBar: true,
				}
				http.Redirect(w, r, "/artiste?id="+strconv.Itoa(ApiObject[i].ID), 302)
				return
			}
		}
	}
	t, _ := template.ParseFiles("./templates/erreur.html")
	t.Execute(w, nil)
}

func Error(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
