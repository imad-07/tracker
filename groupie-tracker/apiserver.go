package apiserver

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"api/apis/apirunner"
	"api/apis/structs"
)

var locs map[string]bool

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotallowedHandler(w, r)
		return
	}
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
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
	locs = Getloc(artistLocation)

	xx := stint(members)
	df, err1 := strconv.Atoi(creationDateFrom)
	dt, err2 := strconv.Atoi(creationDateTo)
	af, err3 := strconv.Atoi(albumDateFrom)
	at, err4 := strconv.Atoi(albumDateTo)

	var ids []int
	var far []structs.Artists
	if _, exists := locs[location]; !exists && location != "" {
		badrequestHandler(w, r)
		return
	}
	for _, x := range artistLocation.Index {
		if isin(x.Locations, location) {
			ids = append(ids, x.ID)
		}
	}
	for _, x := range artists {
		if isint(ids, x.ID) {
			if x.CreationDate >= df && x.CreationDate <= dt && adate(x.FirstAlbum) >= af && adate(x.FirstAlbum) <= at && isint(xx, len(x.Members)) {
				if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
					badrequestHandler(w, r)
					return
				}
				far = append(far, x)
			} else if df == 0 && dt == 0 && af == 0 && at == 0 && members == nil {
				far = append(far, x)
			}
		}
	}

	data := structs.PageData{
		Artists:        far,
		ArtistLocation: artistLocation,
		ArtistDates:    artistDates,
		ArtistRelation: artistRelation,
		Locs:           locs,
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
		http.Error(w,"error",http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w,"error",http.StatusInternalServerError)
		return
	}
}

func badrequestHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("apis/apiserver/templates/Badrequest.html")
	if err != nil {
		w.WriteHeader(400)
	}
	w.WriteHeader(400)
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

func Getloc(location structs.ArtistLocation) map[string]bool {
	locs := map[string]bool{}
	for _, x := range location.Index {
		for _, y := range x.Locations {
			locs[y] = true
		}
	}

	return locs
}

func HandleLocations(w http.ResponseWriter, r *http.Request) {
	locations := strings.Split(r.URL.Path, "/")

	if _, exists := locs[locations[len(locations)-1]]; !exists && locations[len(locations)-1] != "" {
		badrequestHandler(w, r)
		return
	}
	if locations[len(locations)-1] == "" {
		badrequestHandler(w, r)
		return
	}
	link := apirunner.Geolocator(locations[len(locations)-1])
	if strings.HasPrefix(link, "no results found for city:") {
		badrequestHandler(w, r)
		return
	}
	http.Redirect(w, r, link, 303)
}
