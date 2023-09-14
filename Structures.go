package groupi

// Structure pour les liens des apis
type Apis struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

// Structure pour les artistes
type Art struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Member       []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// Structure finale de l'artiste
type Artistes struct {
	Art      *Art
	Location []string
	Date     [][]string
}

// Structures pour les Locations
type Loca struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}

type Location struct {
	Index []Loca `json:"index"`
}

// Structures pour les Relations
type Rela struct {
	Id       int                 `json:"id"`
	Relation map[string][]string `json:"datesLocations"`
}

type Relation struct {
	Index []Rela `json:"index"`
}
