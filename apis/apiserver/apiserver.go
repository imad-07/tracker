package apiserver

import (
	"net/http"
	"strconv"
	"text/template"

	"api/apis/apirunner"
	"api/apis/structs"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}
	if r.Method != http.MethodGet {
		methodNotallowedHandler(w, r)
		return
	}
	artists, artistRelation, artistLocation, artistDates, e := apirunner.ApiRunner()
	if e != nil {
		serverHandler(w, r)
		return
	}
	tmpl, err := template.ParseFiles("apis/apiserver/templates/index.html")
	if err != nil {
		serverHandler(w, r)
		return
	}
	creationDateFrom := r.FormValue("creation-date-from")
	creationDateTo := r.FormValue("creation-date-to")
	albumDateFrom := r.FormValue("album-date-from")
	albumDateTo := r.FormValue("album-date-to")
	location := r.FormValue("location")
	members := r.Form["members"]

	xx := stint(members)
	df, _ := strconv.Atoi(creationDateFrom)
	dt, _ := strconv.Atoi(creationDateTo)
	af, _ := strconv.Atoi(albumDateFrom)
	at, _ := strconv.Atoi(albumDateTo)

	var ids []int
	var far []structs.Artists

	for _, x := range artistLocation.Index {
		if isin(x.Locations, location) {
			ids = append(ids, x.ID)
		}
	}
	for _, x := range artists {
		if isint(ids, x.ID) {
			if x.CreationDate >= df && x.CreationDate <= dt && adate(x.FirstAlbum) >= af && adate(x.FirstAlbum) <= at && isint(xx, len(x.Members)) {
				far = append(far, x)
			}
		}
	}
	data := structs.PageData{
		Artists:        far,
		ArtistLocation: artistLocation,
		ArtistDates:    artistDates,
		ArtistRelation: artistRelation,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		serverHandler(w, r)
		return
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("apis/apiserver/templates/404.html")
	w.WriteHeader(http.StatusNotFound)
	if err != nil {
		serverHandler(w, r)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		serverHandler(w, r)
		return
	}
}

func methodNotallowedHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("apis/apiserver/templates/Method.html")
	if err != nil {
		serverHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	err = tmpl.Execute(w, nil)
	if err != nil {
		serverHandler(w, r)
		return
	}
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("apis/apiserver/templates/Servererror.html")
	if err != nil {
		w.WriteHeader(500)
	}
	w.WriteHeader(http.StatusInternalServerError)
	err = tmpl.Execute(w, nil)
	if err != nil {
		serverHandler(w, r)
		return
	}
}

func stint(str []string) []int {
	var x []int
	for i := 0; i < len(str); i++ {
		b, _ := strconv.Atoi(str[i])
		x = append(x, b)
		b = 0
	}
	return x
}

func isin(slice []string, element string) bool {
	if element == "" {
		return true
	}
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

func isint(slice []int, element int) bool {
	if slice == nil {
		return true
	}
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

func adate(str string) int {
	x, _ := strconv.Atoi(str[len(str)-4:])
	return x
}
