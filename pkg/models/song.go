package models

type Song struct {
	ID             string `json:"id"`
	Art            string `json:"art"`
	End            int    `json:"end"`
	Name           string `json:"name"`
	Start          int    `json:"start"`
	ITunesID       uint32 `json:"itunesid"`
	OriginalArtist string `json:"original_artist"`
}
