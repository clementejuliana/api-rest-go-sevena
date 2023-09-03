package models

type Local struct {
	LocalID int    `json:"local_id,omitempty"`
	Status  string `json:"status,omitempty"`
	Sala    string `json:"sala,omitempty"`
	Setor   string `json:"setor,omitempty"`
}
