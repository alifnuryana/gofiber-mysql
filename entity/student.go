package entity

type Student struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	NPM   string `json:"NPM"`
	Prodi string `json:"prodi"`
}

type Students struct {
	Students []Student `json:"students"`
}
