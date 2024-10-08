package models

type TopTracksResponse struct {
	SpotifyResponse
	Items []Tracks `json:"items"`
}
type Artist struct {
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href string `json:"href"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
type Tracks struct {
	Album struct {
		AlbumType string   `json:"album_type"`
		Artists   []Artist `json:"artists"`
		//AvailableMarkets []string `json:"available_markets"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href                 string  `json:"href"`
		ID                   string  `json:"id"`
		Images               []Image `json:"images"`
		Name                 string  `json:"name"`
		ReleaseDate          string  `json:"release_date"`
		ReleaseDatePrecision string  `json:"release_date_precision"`
		TotalTracks          int     `json:"total_tracks"`
		Type                 string  `json:"type"`
		URI                  string  `json:"uri"`
	} `json:"album"`
	Artists []Artist `json:"artists"`
	//AvailableMarkets []string `json:"available_markets"`
	DiscNumber  int  `json:"disc_number"`
	DurationMs  int  `json:"duration_ms"`
	Explicit    bool `json:"explicit"`
	ExternalIds struct {
		Isrc string `json:"isrc"`
	} `json:"external_ids"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href        string `json:"href"`
	ID          string `json:"id"`
	IsLocal     bool   `json:"is_local"`
	Name        string `json:"name"`
	Popularity  int    `json:"popularity"`
	PreviewURL  string `json:"preview_url"`
	TrackNumber int    `json:"track_number"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
}
