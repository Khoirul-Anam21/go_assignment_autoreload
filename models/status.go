package models


type Info struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int64 `json:"water"`
	Wind  int64 `json:"wind"` 
}
