package apirunner

import (
	"api/apis/api"
	"api/apis/structs"
)

func ApiRunner() ([]structs.Artists, structs.ArtistRelation, structs.ArtistLocation, structs.ArtistDates, error) {
	var artists []structs.Artists
	var relation structs.ArtistRelation
	var artistLocation structs.ArtistLocation
	var artistDates structs.ArtistDates
	var err error

	err = api.ReadApi("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return nil, structs.ArtistRelation{}, structs.ArtistLocation{}, structs.ArtistDates{}, err
	}

	err = api.ReadApi("https://groupietrackers.herokuapp.com/api/relation", &relation)
	if err != nil {
		return nil, structs.ArtistRelation{}, structs.ArtistLocation{}, structs.ArtistDates{}, err
	}

	err = api.ReadApi("https://groupietrackers.herokuapp.com/api/locations", &artistLocation)
	if err != nil {
		return nil, structs.ArtistRelation{}, structs.ArtistLocation{}, structs.ArtistDates{}, err
	}

	err = api.ReadApi("https://groupietrackers.herokuapp.com/api/dates", &artistDates)
	if err != nil {
		return nil, structs.ArtistRelation{}, structs.ArtistLocation{}, structs.ArtistDates{}, err
	}
	return artists, relation, artistLocation, artistDates, err
}
