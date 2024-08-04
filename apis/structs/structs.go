package structs

type Bandlocation struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type Banddates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Artists struct {
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

type ArtistRelation struct {
	Index []struct {
		ID             int         `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
type ArtistLocation struct {
	Index []Bandlocation `json:"index"`
}
type ArtistDates struct {
	Index []Banddates `json:"index"`
}
type PageData struct {
	Artists        []Artists
	ArtistLocation ArtistLocation
	ArtistDates    ArtistDates
	ArtistRelation ArtistRelation
}
