package models

type Ngo struct {
	Id        int     `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	About     string  `json:"about,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}
